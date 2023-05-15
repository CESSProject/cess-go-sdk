/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package rule

import "time"

// BlockInterval is the time interval for generating blocks, in seconds
const BlockInterval = time.Second * time.Duration(6)

// CESSTokenSymbol is the symbol of the CESS token
const CESSTokenSymbol = "TCESS"

const MaxSubmitedIdleFileMeta = 30
