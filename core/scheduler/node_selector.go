/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0

    This package aims to provide CESS developers with an integrated, customizable user data scheduling module
	(that is, distributing user data to nearby nodes with good communication conditions).
	This package contains a node selection module and a data scheduling module built on top of it,
	which can be used by developers respectively.
*/

package scheduler

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

const (
	FIXED_STRATEGY       = "fixed"
	PRIORITY_STRATEGY    = "priority"
	DEFAULT_STRATEGY     = "priority"
	MAX_ALLOWED_NODES    = 120
	MAX_DISALLOWED_NODES = 1024
	DEFAULT_MAX_TTL      = time.Millisecond * 500
	DEFAULT_TIMEOUT      = 5 * time.Second
	DEFAULT_FLUSH_TIME   = time.Hour * 4
	MAX_FAILED_CONN      = 3
)

type Selector interface {
	NewPeersIterator(minNum int) (Iterator, error)
	Feedback(id string, isWork bool)
	FlushPeerNodes(pingTimeout time.Duration, peers ...peer.AddrInfo)
	//FlushlistedPeerNodes(pingTimeout time.Duration, discoverer Discoverer)
}

type Iterator interface {
	GetPeer() (peer.AddrInfo, bool)
}

type Discoverer interface {
	FindPeer(ctx context.Context, id peer.ID) (pi peer.AddrInfo, err error)
}

type NodeList struct {
	AllowedPeers    []string `json:"allowed_peers"` //
	DisallowedPeers []string `json:"disallowed_peers"`
}

type SelectorConfig struct {
	Strategy      string `name:"Strategy" toml:"Strategy" yaml:"Strategy"` //case in "fixed","priority", default: "priority"
	NodeFilePath  string `name:"NodeFile" toml:"NodeFile" yaml:"NodeFile"`
	MaxNodeNum    int    `name:"MaxNodeNum" toml:"MaxNodeNum" yaml:"MaxNodeNum"`
	MaxTTL        int64  `name:"MaxTTL" toml:"MaxTTL" yaml:"MaxTTL"`                      // unit: millisecond, default: 300 ms
	FlushInterval int64  `name:"FlushInterval" toml:"FlushInterval" yaml:"FlushInterval"` // unit: hours, default: 4h
}

type NodeInfo struct {
	NePoints  int
	Available bool
	AddrInfo  peer.AddrInfo
	TTL       time.Duration
	FlushTime time.Time
}

type NodeChan struct {
	queue []NodeInfo
	count int
	index int
}

type NodeSelector struct {
	listPeers   *sync.Map
	blackList   *bloom.BloomFilter
	activePeers *sync.Map
	config      SelectorConfig
}

func NewNodeSelectorWithConfig(config SelectorConfig) (Selector, error) {
	return NewNodeSelector(
		config.Strategy,
		config.NodeFilePath,
		config.MaxNodeNum,
		config.MaxTTL,
		config.FlushInterval,
	)
}

// NewNodeSelector create a new Selector for selecting a best peer list to storage file.
// strategy can be "fixed" or "priority", which means only using or priority using the specified node list respectively.
// nodeFilePath is the json file path of the node list you specify.
// maxNodeNum is used to specify the maximum available stable node.
// The units of parameters maxTTL and flushInterval are milliseconds and hours respectively.
func NewNodeSelector(strategy, nodeFilePath string, maxNodeNum int, maxTTL, flushInterval int64) (Selector, error) {
	selector := new(NodeSelector)
	if maxNodeNum <= 0 || maxNodeNum > MAX_ALLOWED_NODES {
		maxNodeNum = MAX_ALLOWED_NODES
	}
	maxTTL *= int64(time.Millisecond)
	if maxTTL <= 0 || maxTTL > int64(DEFAULT_MAX_TTL) {
		maxTTL = int64(DEFAULT_MAX_TTL)
	}
	flushInterval *= int64(time.Hour)
	if flushInterval <= 0 || flushInterval > int64(DEFAULT_FLUSH_TIME) {
		flushInterval = int64(DEFAULT_FLUSH_TIME)
	}

	var nodeList NodeList
	if _, err := os.Stat(nodeFilePath); err == nil {
		bytes, err := os.ReadFile(nodeFilePath)
		if err != nil {
			return nil, errors.Wrap(err, "create node selector error")
		}
		err = json.Unmarshal(bytes, &nodeList)
		if err != nil {
			return nil, errors.Wrap(err, "create node selector error")
		}
	} else {
		f, err := os.Create(nodeFilePath)
		if err == nil {
			jbytes, err := json.Marshal(nodeList)
			if err == nil {
				f.Write(jbytes)
			}
			f.Close()
		}
	}

	if len(nodeList.AllowedPeers) > MAX_ALLOWED_NODES {
		nodeList.AllowedPeers =
			nodeList.AllowedPeers[:MAX_ALLOWED_NODES]
	}
	if len(nodeList.DisallowedPeers) > MAX_DISALLOWED_NODES {
		nodeList.DisallowedPeers =
			nodeList.DisallowedPeers[:MAX_DISALLOWED_NODES]
	}
	selector.blackList = bloom.NewWithEstimates(100000, 0.01)
	selector.listPeers = &sync.Map{}

	for _, peer := range nodeList.DisallowedPeers {
		selector.blackList.AddString(peer)
	}

	for _, p := range nodeList.AllowedPeers {
		selector.listPeers.Store(p, NodeInfo{})
	}

	switch strategy {
	case PRIORITY_STRATEGY, FIXED_STRATEGY:
	default:
		strategy = DEFAULT_STRATEGY
	}
	selector.config = SelectorConfig{
		MaxNodeNum:    maxNodeNum,
		MaxTTL:        maxTTL,
		Strategy:      strategy,
		NodeFilePath:  nodeFilePath,
		FlushInterval: flushInterval,
	}
	selector.activePeers = &sync.Map{}
	return selector, nil
}

func (c *NodeChan) GetPeer() (peer.AddrInfo, bool) {
	if c.index >= c.count {
		return peer.AddrInfo{}, false
	}
	c.index++
	return c.queue[c.index-1].AddrInfo, true
}

func (c *NodeChan) insertNode(info NodeInfo, maxNum int) {
	var i int
	for i = c.count - 1; i >= c.index; i-- {
		ttl := c.queue[i].TTL - info.TTL
		if ttl > time.Microsecond*5 ||
			(ttl >= 0 && info.NePoints < c.queue[i].NePoints) {
			continue
		}
		break
	}
	if c.count == maxNum {
		if i >= c.count-1 {
			return
		}
	} else {
		c.count++
		c.queue = append(c.queue, NodeInfo{})
	}
	copy(c.queue[i+1+1:], c.queue[i+1:c.count-1])
	c.queue[i+1] = info
}

// Deprecated: This refresh method is invalid in the new SDK version
func (s *NodeSelector) FlushlistedPeerNodes(pingTimeout time.Duration, discoverer Discoverer) {
	s.listPeers.Range(func(key, value any) bool {
		k := key.(string)
		v := value.(NodeInfo)
		if v.Available && time.Since(v.FlushTime) < time.Hour {
			return true
		}
		addr, err := discoverer.FindPeer(context.Background(), peer.ID(k))
		if err != nil {
			log.Println("flush list peer error", err)
			return true
		}
		log.Println("flush list peer", addr)
		v.AddrInfo = addr
		v.FlushTime = time.Now()
		v.TTL = GetConnectTTL(addr.Addrs, pingTimeout)
		if v.TTL > 0 && v.TTL < time.Duration(s.config.MaxTTL) {
			v.Available = true
		}
		s.listPeers.Swap(k, v)
		return true
	})
}

func (s *NodeSelector) FlushPeerNodes(pingTimeout time.Duration, peers ...peer.AddrInfo) {
	for _, peer := range peers {
		key := peer.ID.String()

		if s.blackList.TestString(key) {
			continue
		}
		info := NodeInfo{
			AddrInfo:  peer,
			FlushTime: time.Now(),
		}
		point := s.activePeers
		if value, ok := s.listPeers.Load(key); ok {
			v := value.(NodeInfo)
			if v.Available &&
				time.Since(v.FlushTime) < time.Duration(s.config.FlushInterval) {
				continue
			}
			info.NePoints = v.NePoints
			point = s.listPeers
		} else if value, ok := s.activePeers.Load(key); ok {
			v := value.(NodeInfo)
			if v.Available &&
				time.Since(v.FlushTime) < time.Duration(s.config.FlushInterval) {
				continue
			}
			info.NePoints = v.NePoints
		}
		info.TTL = GetConnectTTL(peer.Addrs, pingTimeout)
		if info.TTL > 0 && info.TTL < time.Duration(s.config.MaxTTL) {
			info.Available = true
		}
		log.Println("save", info)
		point.Store(key, info)
	}
}

func (s *NodeSelector) NewPeersIterator(minNum int) (Iterator, error) {
	if minNum > s.config.MaxNodeNum {
		return nil, errors.Wrap(errors.New("not enough nodes"), "create peers iterator error")
	}
	nodeCh := &NodeChan{
		queue: make([]NodeInfo, 0, s.config.MaxNodeNum),
	}
	handle := func(key, value any) bool {
		v := value.(NodeInfo)
		if !v.Available {
			return true
		}
		log.Println("inert node to queue", v.AddrInfo.ID.String())
		nodeCh.insertNode(v, s.config.MaxNodeNum)
		return true
	}
	s.listPeers.Range(handle)

	if s.config.Strategy != FIXED_STRATEGY {
		s.activePeers.Range(handle)
	}
	if nodeCh.count < minNum {
		return nil, errors.Wrap(errors.New("not enough nodes"), "create peers iterator error")
	}
	log.Println("node queue", nodeCh.queue)
	return nodeCh, nil
}

func (s *NodeSelector) Feedback(id string, isWrok bool) {
	if isWrok {
		s.reflashPeer(id)
	} else {
		s.removePeer(id)
	}
}

func (s *NodeSelector) reflashPeer(id string) {
	if s.blackList.TestString(id) {
		return
	}
	if v, ok := s.listPeers.Load(id); ok {
		info := v.(NodeInfo)
		info.NePoints = 0
		s.activePeers.Store(id, info)
		return
	}
	if v, ok := s.activePeers.Load(id); ok {
		info := v.(NodeInfo)
		info.NePoints = 0
		s.activePeers.Store(id, info)
	}
}

func (s *NodeSelector) removePeer(id string) {
	if s.blackList.TestString(id) {
		return
	}
	if v, ok := s.listPeers.Load(id); ok {
		info := v.(NodeInfo)
		info.NePoints++
		s.activePeers.Store(id, info)
		return
	}
	if v, ok := s.activePeers.Load(id); ok {
		info := v.(NodeInfo)
		info.NePoints++
		if info.NePoints < MAX_FAILED_CONN {
			s.activePeers.Store(id, info)
		}
		s.activePeers.Delete(id)
		s.blackList.AddString(id)
	}
}

func GetConnectTTL(addrs []multiaddr.Multiaddr, timeout time.Duration) time.Duration {
	var minTTL time.Duration = timeout
	for _, addr := range addrs {
		ip := strings.Split(addr.String(), "/")[2]
		if ok, err := utils.IsIntranetIpv4(ip); err != nil || ok {
			continue
		}
		ttl, err := utils.PingNode(ip, timeout)
		if err != nil {
			continue
		}
		if ttl < minTTL {
			minTTL = ttl
		}
	}
	if minTTL == timeout {
		minTTL = 0
	}
	return minTTL
}
