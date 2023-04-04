package utils

import (
	"regexp"
	"strings"

	"github.com/CESSProject/sdk-go/core/rule"
)

func CheckBucketName(name string) bool {
	if len(name) < rule.MinBucketNameLength || len(name) > rule.MaxBucketNameLength {
		return false
	}

	re, err := regexp.Compile(`^[a-z0-9.-]{3,63}$`)
	if err != nil {
		return false
	}

	if !re.MatchString(name) {
		return false
	}

	if strings.Contains(name, "..") {
		return false
	}

	if byte(name[0]) == byte('.') ||
		byte(name[0]) == byte('-') ||
		byte(name[len(name)-1]) == byte('.') ||
		byte(name[len(name)-1]) == byte('-') {
		return false
	}

	return !IsIPv4(name)
}
