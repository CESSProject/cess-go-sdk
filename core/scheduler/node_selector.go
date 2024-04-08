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
	DEFAULT_MAX_TTL      = time.Millisecond * 300
	DEFAULT_TIMEOUT      = 5 * time.Second
	MAX_FAILED_CONN      = 3
)

type Selector interface {
	NewPeersIterator(minNum int) (Iterator, error)
	Feedback(id string, isWork bool)
	FlushPeerNodes(pingTimeout time.Duration, peers ...peer.AddrInfo)
}

type Iterator interface {
	GetPeer() (peer.AddrInfo, bool)
}

type NodeList struct {
	AllowedPeers    []string `json:"allowed_peers"` //
	DisallowedPeers []string `json:"disallowed_peers"`
}

type SelectorConfig struct {
	Strategy     string `name:"Strategy" toml:"Strategy" yaml:"Strategy"` //"fixed","priority"
	NodeFilePath string `name:"NodeFile" toml:"NodeFile" yaml:"NodeFile"`
	MaxNodeNum   int    `name:"MaxNodeNum" toml:"MaxNodeNum" yaml:"MaxNodeNum"`
	MaxTTL       int64  `name:"MaxTTL" toml:"MaxTTL" yaml:"MaxTTL"`
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
	)
}

func NewNodeSelector(strategy, nodeFilePath string, maxNodeNum int, maxTTL int64) (Selector, error) {
	selector := new(NodeSelector)
	if maxNodeNum <= 0 || maxNodeNum > MAX_ALLOWED_NODES {
		maxNodeNum = MAX_ALLOWED_NODES
	}
	if maxTTL <= 0 || maxTTL > int64(DEFAULT_MAX_TTL) {
		maxTTL = int64(DEFAULT_MAX_TTL)
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
		MaxNodeNum:   maxNodeNum,
		MaxTTL:       maxTTL,
		Strategy:     strategy,
		NodeFilePath: nodeFilePath,
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
	log.Println("node queue:", c.queue)
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
		info.TTL = GetConnectTTL(peer.Addrs, pingTimeout)
		if info.TTL > 0 && info.TTL < time.Duration(s.config.MaxTTL) {
			info.Available = true
		}
		if value, ok := s.listPeers.Load(key); ok {
			v := value.(NodeInfo)
			info.NePoints = v.NePoints
			s.listPeers.Swap(key, info)
			continue
		}

		if value, ok := s.activePeers.Load(key); ok {
			v := value.(NodeInfo)
			info.NePoints = v.NePoints
			s.activePeers.Swap(key, v)
			continue
		}
		s.activePeers.Store(key, info)
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
	return nil, nil
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
	var minTTL time.Duration
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
			ttl = minTTL
		}
	}
	return minTTL
}
