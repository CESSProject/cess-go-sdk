/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/process"
)

type result struct {
	Fid         string
	SegmentData []pattern.SegmentDataInfo
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing parameter")
		return
	}
	seg, rhash, err := process.ShardedEncryptionProcessing(os.Args[1], "")
	if err != nil {
		fmt.Println(err)
		return
	}
	var myResult = result{
		Fid:         rhash,
		SegmentData: seg,
	}
	buf, err := json.Marshal(myResult)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf)
}
