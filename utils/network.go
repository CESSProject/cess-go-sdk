/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"context"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/go-ping/ping"
	"github.com/pixelbender/go-traceroute/traceroute"
)

var regstr = `\d+\.\d+\.\d+\.\d+`
var reg = regexp.MustCompile(regstr)

func FildIpv4(data []byte) (string, bool) {
	result := reg.Find(data)
	return string(result), len(result) > 0
}

func IsIntranetIpv4(ipv4 string) (bool, error) {
	ip := net.ParseIP(ipv4)
	if ip == nil || !strings.Contains(ipv4, ".") {
		return false, errors.New("invalid ipv4")
	}
	if ip.IsLoopback() {
		return true, nil
	}
	if ip.IsPrivate() {
		return true, nil
	}
	return false, nil
}

// IsIPv4 is used to determine whether ipAddr is an ipv4 address
func IsIPv4(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	return ip != nil && strings.Contains(ipAddr, ".")
}

// IsIPv6 is used to determine whether ipAddr is an ipv6 address
func IsIPv6(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	return ip != nil && strings.Contains(ipAddr, ":")
}

func IsPortInUse(port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%d", port), 3*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// CheckDomain returns an error if the domain name is not valid.
// See https://tools.ietf.org/html/rfc1034#section-3.5 and
// https://tools.ietf.org/html/rfc1123#section-2
func CheckDomain(name string) error {
	name = strings.TrimPrefix(name, "http://")
	name = strings.TrimPrefix(name, "https://")
	name = strings.TrimSuffix(name, "/")
	switch {
	case len(name) == 0:
		return nil // an empty domain name will result in a cookie without a domain restriction
	case len(name) > 50:
		return fmt.Errorf("domain name length is %d, can't exceed 50", len(name))
	}
	var l int
	for i := 0; i < len(name); i++ {
		b := name[i]
		if b == '.' {
			// check domain labels validity
			switch {
			case i == l:
				return fmt.Errorf("domain has invalid character '.' at offset %d, label can't begin with a period", i)
			case i-l > 63:
				return fmt.Errorf("domain byte length of label '%s' is %d, can't exceed 63", name[l:i], i-l)
			case name[l] == '-':
				return fmt.Errorf("domain label '%s' at offset %d begins with a hyphen", name[l:i], l)
			case name[i-1] == '-':
				return fmt.Errorf("domain label '%s' at offset %d ends with a hyphen", name[l:i], l)
			}
			l = i + 1
			continue
		}
		// test label character validity, note: tests are ordered by decreasing validity frequency
		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b == '-' || b >= 'A' && b <= 'Z') {
			// show the printable unicode character starting at byte offset i
			c, _ := utf8.DecodeRuneInString(name[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("domain has invalid rune at offset %d", i)
			}
			return fmt.Errorf("domain has invalid character '%c' at offset %d", c, i)
		}
	}
	// check top level domain validity
	switch {
	case l == len(name):
		return fmt.Errorf("domain has missing top level domain, domain can't end with a period")
	case len(name)-l > 63:
		return fmt.Errorf("domain's top level domain '%s' has byte length %d, can't exceed 63", name[l:], len(name)-l)
	case name[l] == '-':
		return fmt.Errorf("domain's top level domain '%s' at offset %d begin with a hyphen", name[l:], l)
	case name[len(name)-1] == '-':
		return fmt.Errorf("domain's top level domain '%s' at offset %d ends with a hyphen", name[l:], l)
	case name[l] >= '0' && name[l] <= '9':
		return fmt.Errorf("domain's top level domain '%s' at offset %d begins with a digit", name[l:], l)
	}
	return nil
}

// PingNode be used to ping target node by ICMP protocol, It returns the TTL from target node.
// Requires root privileges to use.
func PingNode(addr string, timeout time.Duration) (time.Duration, error) {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return 0, err
	}
	pinger.Count = 4
	pinger.Timeout = timeout
	err = pinger.Run()
	if err != nil {
		return 0, err
	}
	stats := pinger.Statistics()
	return stats.AvgRtt, nil
}

// TraceRoute be used to trace route to target node by ICMP protocol, It returns the distance and TTL from target node.
// Requires root privileges to use
func TraceRoute(ip string, timeout time.Duration) (int, time.Duration, error) {
	t := &traceroute.Tracer{
		Config: traceroute.Config{
			Delay:    50 * time.Millisecond,
			Timeout:  timeout,
			MaxHops:  20,
			Count:    3,
			Networks: []string{"ip4:icmp", "ip4:ip"},
		},
	}
	defer t.Close()
	var (
		avgDist int
		avgRtt  time.Duration
		count   int
	)
	err := t.Trace(context.Background(), net.ParseIP(ip), func(reply *traceroute.Reply) {
		count++
		avgDist += reply.Hops
		avgRtt += reply.RTT
	})

	if err != nil {
		return 0, 0, err
	}
	if count > 0 {
		return avgDist / count, avgRtt / time.Duration(count), nil
	}
	return 0, 0, errors.New("network unreachable")
}
