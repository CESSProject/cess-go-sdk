/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"log"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/process"
)

type result struct {
	Fid         string
	SegmentData []pattern.SegmentDataInfo
}

func main() {
	var file = "process.go"
	seg, rhash, err := process.ShardedEncryptionProcessing(file, "")
	if err != nil {
		log.Println(err)
		return
	}
	var myResult = result{
		Fid:         rhash,
		SegmentData: seg,
	}
	fmt.Println(myResult)
}
