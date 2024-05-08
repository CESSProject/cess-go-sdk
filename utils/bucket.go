/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"regexp"
	"strings"

	"github.com/CESSProject/cess-go-sdk/config"
)

var re = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

func CheckBucketName(name string) bool {
	if len(name) < config.MinBucketNameLength || len(name) > config.MaxBucketNameLength {
		return false
	}

	if !re.MatchString(name) {
		return false
	}

	if strings.Contains(name, " ") {
		return false
	}

	if strings.Count(name, ".") > 2 {
		return false
	}

	if byte(name[0]) == byte('.') ||
		byte(name[0]) == byte('-') ||
		byte(name[0]) == byte('_') ||
		byte(name[len(name)-1]) == byte('.') ||
		byte(name[len(name)-1]) == byte('-') ||
		byte(name[len(name)-1]) == byte('_') {
		return false
	}

	if IsIPv4(name) {
		return false
	}

	if IsIPv6(name) {
		return false
	}

	return true
}
