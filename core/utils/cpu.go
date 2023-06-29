/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetDirFreeSpace(dir string) (uint64, error) {
	sageStat, err := disk.Usage(dir)
	return sageStat.Free, err
}

func GetSysMemAvailable() (uint64, error) {
	var result uint64
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, errors.Wrapf(err, "[mem.VirtualMemory]")
	}
	result = memInfo.Available
	swapInfo, err := mem.SwapMemory()
	if err != nil {
		return result, nil
	}
	return result + swapInfo.Free, nil
}
