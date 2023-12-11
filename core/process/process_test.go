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

func TestProcessingData(t *testing.T) {
	var processFile = "./process.go"
	segmentData, roothash, err := ProcessingData(processFile)
	assert.NoError(t, err)
	fmt.Println(roothash)
	for _, segment := range segmentData {
		for _, fragment := range segment.FragmentHash {
			os.Remove(fragment)
		}
	}
}

func TestShardedEncryptionProcessing(t *testing.T) {
	var processFile = "./process.go"
	segmentData, roothash, err := ShardedEncryptionProcessing(processFile, "123456")
	assert.NoError(t, err)
	fmt.Println(roothash)
	for _, segment := range segmentData {
		for _, fragment := range segment.FragmentHash {
			os.Remove(fragment)
		}
	}
}
