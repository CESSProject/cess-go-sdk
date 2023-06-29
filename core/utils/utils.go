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
