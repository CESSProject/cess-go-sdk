/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package process

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullProcessing(t *testing.T) {
	var processFile = "./process.go"

	// not encryption
	segmentData, fid, err := FullProcessing(processFile, "", ".")
	assert.NoError(t, err)
	fmt.Println("not encryption: ", fid)
	for _, segment := range segmentData {
		for _, fragment := range segment.FragmentHash {
			os.Remove(fragment)
		}
	}

	// encryption
	segmentData, fid, err = FullProcessing(processFile, "123456", ".")
	assert.NoError(t, err)
	fmt.Println("encryption: ", fid)
	for _, segment := range segmentData {
		for _, fragment := range segment.FragmentHash {
			os.Remove(fragment)
		}
	}
}
