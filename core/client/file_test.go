/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessingData(t *testing.T) {
	var file = "./file_test.go"
	var cli = new(Cli)
	segments, roothash, err := cli.ProcessingData(file)
	assert.NoError(t, err)
	fmt.Println(roothash)
	fmt.Println(segments)
}
