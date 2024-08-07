/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package config

const (
	SIZE_1KiB = 1024
	SIZE_1MiB = 1024 * SIZE_1KiB
	SIZE_1GiB = 1024 * SIZE_1MiB
	SIZE_1TiB = 1024 * SIZE_1GiB

	SegmentSize  = 32 * SIZE_1MiB
	FragmentSize = 8 * SIZE_1MiB
	DataShards   = 4
	ParShards    = 8
)

const (
	MinBucketNameLength = 3
	MaxBucketNameLength = 63
	MaxDomainNameLength = 100
)

const (
	// default name
	CharacterName_Default = "cess-sdk-go"

	// offcial gateway address
	PublicGatewayAddr = "https://deoss-pub-gateway.cess.network/"
	// offcial gateway account
	PublicGatewayAccount = "cXhwBytXqrZLr1qM5NHJhCzEMckSTzNKw17ci2aHft6ETSQm9"
)
