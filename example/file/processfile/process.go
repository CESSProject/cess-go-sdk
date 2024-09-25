/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"log"

	"github.com/CESSProject/cess-go-sdk/core/process"
)

func main() {
	var file = "process.go"
	segmentInfo, fid, err := process.FullProcessing(file, "", ".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fid: ", fid)
	fmt.Println("segment info: ", segmentInfo)
}
