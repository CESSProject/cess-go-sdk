/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"reflect"
)

func IsInterfaceNIL(i interface{}) bool {
	ret := i == nil
	if !ret {
		defer func() {
			recover()
		}()
		va := reflect.ValueOf(i)
		if va.Kind() == reflect.Ptr {
			return va.IsNil()
		}
		return false
	}
	return ret
}

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
