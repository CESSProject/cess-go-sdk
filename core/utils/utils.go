/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	// std
	"fmt"
	"strings"
	"unicode/utf8"

	// 3rd party libs
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func CompareSlice(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// NumsToByteStr utility function takes an array of unsigned integers and output the corresponding byte string representing it.
// For example: `[18, 15]` to `120F`.
//
// For `opts` in second parameter it is a map expecting:
//   - `space` bool: whether to add a space between each byte, default to `false`
//   - `prefix` bool: whether to add `0x` as the prefix, default to `false`
//   - `uppercase` bool: whether to display hexadecimal in upper case, default to `true`
//
// If there is an integer larger than 255 in the array, an error is returned.
func NumsToByteStr[T constraints.Unsigned](nums []T, opts map[string]bool) (string, error) {
	// default value for opts
	var _opts = map[string]bool{}
	val, ok := opts["space"]
	_opts["space"] = lo.Ternary(ok, val, false)
	val, ok = opts["prefix"]
	_opts["prefix"] = lo.Ternary(ok, val, false)
	val, ok = opts["uppercase"]
	_opts["uppercase"] = lo.Ternary(ok, val, true)

	byteStr := []string{}
	for _, num := range nums {
		if num > 255 {
			return "", fmt.Errorf("contains number larger than 255: %v", num)
		}
		byteStr = append(byteStr, fmt.Sprintf(lo.Ternary((_opts)["uppercase"], "%02X", "%02x"), num))
	}

	prefix := lo.Ternary((_opts)["prefix"], "0x", "")
	return prefix + strings.Join(byteStr, lo.Ternary((_opts)["space"], " ", "")), nil
}

func NumsToByteStrDefault[T constraints.Unsigned](nums []T) (string, error) {
	return NumsToByteStr(nums, map[string]bool{})
}

// CheckDomain returns an error if the domain name is not valid.
// See https://tools.ietf.org/html/rfc1034#section-3.5 and
// https://tools.ietf.org/html/rfc1123#section-2.
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
