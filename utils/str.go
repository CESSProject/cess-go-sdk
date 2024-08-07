/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ----------------------- Random key -----------------------
const (
	letterIdBits = 6
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
	baseStr      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]{}+-*/_=."
)

var regexp_array = regexp.MustCompile(`\[(.*?)\]`)

// Generate random password
func GetRandomcode(length uint8) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, length)
	l := len(baseStr)
	for i := uint8(0); i < length; i++ {
		bytes[i] = baseStr[r.Intn(l)]
	}
	return string(bytes)
}

func RandStr(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(baseStr) {
			sb.WriteByte(baseStr[idx])
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return sb.String()
}

func RandSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	if rv.Type().Kind() != reflect.Slice {
		return
	}

	length := rv.Len()
	if length < 2 {
		return
	}

	swap := reflect.Swapper(slice)
	for i := length - 1; i >= 0; i-- {
		j := rand.New(rand.NewSource(time.Now().Unix())).Intn(length)
		swap(i, j)
	}
	return
}

func ExtractArray(str string) []byte {
	match := regexp_array.FindString(str)
	match = strings.TrimPrefix(match, "[")
	match = strings.TrimSuffix(match, "]")
	array := strings.Split(match, " ")
	var s = make([]byte, len(array))
	for i := 0; i < len(array); i++ {
		in, err := strconv.Atoi(array[i])
		if err != nil {
			panic(err)
		}
		s[i] = byte(in)
	}
	return s
}
