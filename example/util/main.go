/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"

	"github.com/CESSProject/cess-go-sdk/chain"
)

func main() {
	fmt.Println(chain.H160ToSS58("0xFd0Cc11A9ffbA29F7db7734b6dc39b1e5212Bb1c", 11330))
}
