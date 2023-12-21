/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func WriteBufToFile(buf []byte, file string) error {
	baseDir := filepath.Dir(file)
	tempFilepath := filepath.Join(baseDir, uuid.New().String())
	f, err := os.Create(tempFilepath)
	if err != nil {
		return err
	}
	defer func() {
		if f != nil {
			f.Close()
		}
		os.Remove(tempFilepath)
	}()

	num, err := f.Write(buf)
	if err != nil || num != len(buf) {
		return errors.New("write file error")
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	return os.Rename(tempFilepath, file)
}
