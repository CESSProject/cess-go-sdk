/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/CESSProject/cess-go-sdk/utils"
	keyring "github.com/CESSProject/go-keyring"
	"github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"
)

func (c *chainClient) StoreFile(url, file, bucket string) (string, error) {
	fstat, err := os.Stat(file)
	if err != nil {

		return "", err
	}

	if fstat.IsDir() {
		return "", errors.New("not a file")
	}

	if fstat.Size() == 0 {
		return "", errors.New("empty file")
	}

	if !utils.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
	}

	kr, _ := keyring.FromURI(c.GetURI(), keyring.NetSubstrate{})

	// sign message
	message := utils.GetRandomcode(16)
	sig, _ := kr.Sign(kr.SigningContext([]byte(message)))

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	//
	formFile, err := writer.CreateFormFile("file", fstat.Name())
	if err != nil {
		return "", err
	}

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(formFile, f)
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("BucketName", bucket)
	req.Header.Set("Account", c.GetSignatureAcc())
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		if len(respbody) > 0 {
			return "", errors.New(string(respbody))
		}
		return "", errors.New(fmt.Sprintf("upload failed, code: %d", resp.StatusCode))
	}

	return strings.TrimPrefix(strings.TrimSuffix(string(respbody), "\""), "\""), nil
}

func (c *chainClient) StoreObject(url string, reader io.Reader, bucket string) (string, error) {
	if !utils.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
	}

	kr, _ := keyring.FromURI(c.GetURI(), keyring.NetSubstrate{})

	// sign message
	message := utils.GetRandomcode(16)
	sig, _ := kr.Sign(kr.SigningContext([]byte(message)))

	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return "", err
	}

	req.Header.Set("BucketName", bucket)
	req.Header.Set("Account", c.GetSignatureAcc())
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		if len(respbody) > 0 {
			return "", errors.New(string(respbody))
		}
		return "", errors.New(fmt.Sprintf("upload failed, code: %d", resp.StatusCode))
	}

	return strings.TrimPrefix(strings.TrimSuffix(string(respbody), "\""), "\""), nil
}

func (c *chainClient) RetrieveFile(url, fid, savepath string) error {
	fstat, err := os.Stat(savepath)
	if err == nil {
		if fstat.IsDir() {
			savepath = filepath.Join(savepath, fid)
		}
		if fstat.Size() > 0 {
			return nil
		}
	}

	if url == "" {
		return errors.New("empty url")
	}

	if url[len(url)-1] != byte(47) {
		url += "/"
	}

	f, err := os.Create(savepath)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodGet, url+fid, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Operation", "download")
	req.Header.Set("Account", c.GetSignatureAcc())

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed")
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c *chainClient) RetrieveObject(url, fid string) (io.ReadCloser, error) {
	if url == "" {
		return nil, errors.New("empty url")
	}

	if url[len(url)-1] != byte(47) {
		url += "/"
	}

	req, err := http.NewRequest(http.MethodGet, url+fid, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Operation", "download")

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Retrieve failed, code: %d", resp.StatusCode))
	}

	return resp.Body, nil
}

// func (c *chainClient) StorageData(roothash string, segment []pattern.SegmentDataInfo, minerTaskList []pattern.MinerTaskList) error {
// 	if !c.enabledP2P {
// 		return errors.New("P2P network not enabled")
// 	}

// 	var err error

// 	// query all assigned miner multiaddr
// 	peerids, err := c.QueryAssignedMinerPeerId(minerTaskList)
// 	if err != nil {
// 		return errors.Wrapf(err, "[QueryAssignedMinerPeerId]")
// 	}

// 	basedir := filepath.Dir(segment[0].FragmentHash[0])
// 	for i := 0; i < len(peerids); i++ {
// 		for j := 0; j < len(minerTaskList[i].Hash); j++ {
// 			err = c.WriteFileAction(peerids[j], roothash, filepath.Join(basedir, string(minerTaskList[i].Hash[j][:])))
// 			if err != nil {
// 				return errors.Wrapf(err, "[WriteFileAction]")
// 			}
// 		}
// 	}

// 	return nil
// }

// func (c *chainClient) FindPeers() map[string]peer.AddrInfo {
// 	var peerMap = make(map[string]peer.AddrInfo, 100)
// 	timeOut := time.NewTicker(time.Second * 10)
// 	defer timeOut.Stop()
// 	c.RouteTableFindPeers(0)
// 	for {
// 		select {
// 		case peer, ok := <-c.GetDiscoveredPeers():
// 			if !ok {
// 				return peerMap
// 			}
// 			if len(peer.Responses) == 0 {
// 				break
// 			}
// 			for _, v := range peer.Responses {
// 				peerMap[v.ID.Pretty()] = *v
// 			}
// 		case <-timeOut.C:
// 			return peerMap
// 		}
// 	}
// }

// func (c *chainClient) QueryAssignedMinerPeerId(minerTaskList []pattern.MinerTaskList) ([]peer.ID, error) {
// 	var peerids = make([]peer.ID, len(minerTaskList))
// 	for i := 0; i < len(minerTaskList); i++ {
// 		minerInfo, err := c.QueryStorageMiner(minerTaskList[i].Account[:])
// 		if err != nil {
// 			return peerids, err
// 		}
// 		peerids[i], _ = peer.Decode(string(minerInfo.PeerId[:]))
// 	}
// 	return peerids, nil
// }
