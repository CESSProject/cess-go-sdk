/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"errors"
	"fmt"
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

func GenerateFile(mb uint32, path string) error {
	if mb == 0 {
		return fmt.Errorf("file size is zero")
	}
	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("%s already exists", path)
	}
	baseDir := filepath.Dir(path)
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
	line := mb * 1024 * 1024 / 4096
	for i := uint32(0); i < line; i++ {
		f.WriteString(RandStr(4095))
		f.WriteString("\n")
	}
	err = f.Sync()
	if err != nil {
		return err
	}

	return os.Rename(tempFilepath, path)
}
