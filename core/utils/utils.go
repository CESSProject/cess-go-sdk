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

// TODO: write the function docs
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
