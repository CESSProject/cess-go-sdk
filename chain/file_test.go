/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessingData(t *testing.T) {
	var c = &chainClient{}
	var processFile = "./file_test.go"
	segmentData, roothash, err := c.ProcessingData(processFile)
	assert.NoError(t, err)
	fmt.Println(roothash)
	for _, segment := range segmentData {
		for _, fragment := range segment.FragmentHash {
			os.Remove(fragment)
		}
	}
}
