package chain

import (
	"time"

	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

func (c *chainClient) QueryBucketInfo(puk []byte, bucketname string) (BucketInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data BucketInfo

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	name_byte, err := codec.Encode(bucketname)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		BUCKET,
		owner,
		name_byte,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) QueryBucketList(puk []byte) ([]types.Bytes, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data []types.Bytes

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		BUCKETLIST,
		owner,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryFileMetaData
func (c *chainClient) QueryFileMetadata(roothash string) (FileMetadata, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var (
		data FileMetadata
		hash FileHash
	)

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	if len(hash) != len(roothash) {
		return data, errors.New("invalid filehash")
	}

	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(roothash[i])
	}

	b, err := codec.Encode(hash)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		FILE,
		b,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) QueryStorageOrder(roothash string) (StorageOrder, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data StorageOrder
	var hash FileHash

	if len(hash) != len(roothash) {
		return data, errors.New("invalid filehash")
	}

	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(roothash[i])
	}

	b, err := codec.Encode(hash)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		DEALMAP,
		b,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) QueryPendingReplacements(puk []byte) (types.U32, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data types.U32

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		PENDINGREPLACE,
		owner,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) SubmitIdleMetadata(idlefiles []IdleMetadata) (string, error) {
	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return txhash, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	call, err := types.NewCall(
		c.metadata,
		TX_FILEBANK_ADDIDLESPACE,
		idlefiles,
	)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewCall]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		SYSTEM,
		ACCOUNT,
		c.keyring.PublicKey,
	)
	if err != nil {
		return txhash, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return txhash, ERR_RPC_EMPTY_VALUE
	}

	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.runtimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.runtimeVersion.TransactionVersion,
	}

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		return txhash, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
	}
	defer sub.Unsubscribe()
	timeout := time.NewTimer(c.timeForBlockOut)
	defer timeout.Stop()
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				//events := EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				// h, err := c.api.RPC.State.GetStorageRaw(c.keyEvents, status.AsInBlock)
				// if err != nil {
				// 	return txhash, errors.Wrap(err, "[GetStorageRaw]")
				// }

				// types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)

				// if len(events.FileBank_DeleteFile) > 0 {
				// 	return txhash, events.FileBank_DeleteFile[0].FailedList
				// }
				return txhash, nil
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) CreateBucket(owner_pkey []byte, name string) (string, error) {
	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return txhash, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	acc, err := types.NewAccountID(owner_pkey)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewAccountID]")
	}

	call, err := types.NewCall(
		c.metadata,
		TX_FILEBANK_PUTBUCKET,
		*acc,
		types.NewBytes([]byte(name)),
	)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewCall]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		SYSTEM,
		ACCOUNT,
		c.keyring.PublicKey,
	)
	if err != nil {
		return txhash, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return txhash, ERR_RPC_EMPTY_VALUE
	}

	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.runtimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.runtimeVersion.TransactionVersion,
	}

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		return txhash, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
	}
	defer sub.Unsubscribe()
	timeout := time.NewTimer(c.timeForBlockOut)
	defer timeout.Stop()
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				events := EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				h, err := c.api.RPC.State.GetStorageRaw(c.keyEvents, status.AsInBlock)
				if err != nil {
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}

				types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)

				if len(events.FileBank_CreateBucket) > 0 {
					return txhash, nil
				}
				return txhash, errors.New(ERR_Failed)
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) DeleteBucket(owner_pkey []byte, name string) (string, error) {
	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return txhash, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	acc, err := types.NewAccountID(owner_pkey)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewAccountID]")
	}

	call, err := types.NewCall(
		c.metadata,
		TX_FILEBANK_DELBUCKET,
		*acc,
		types.NewBytes([]byte(name)),
	)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewCall]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		SYSTEM,
		ACCOUNT,
		c.keyring.PublicKey,
	)
	if err != nil {
		return txhash, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return txhash, ERR_RPC_EMPTY_VALUE
	}

	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.runtimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.runtimeVersion.TransactionVersion,
	}

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		return txhash, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
	}
	defer sub.Unsubscribe()
	timeout := time.NewTimer(c.timeForBlockOut)
	defer timeout.Stop()
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				events := EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				h, err := c.api.RPC.State.GetStorageRaw(c.keyEvents, status.AsInBlock)
				if err != nil {
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}

				types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)

				if len(events.FileBank_DeleteBucket) > 0 {
					return txhash, nil
				}
				return txhash, errors.New(ERR_Failed)
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) UploadDeclaration(filehash string, dealinfo []SegmentList, user UserBrief) (string, error) {
	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return txhash, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	var hash FileHash
	if len(filehash) != len(hash) {
		return txhash, errors.New("invalid filehash")
	}
	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(filehash[i])
	}

	call, err := types.NewCall(
		c.metadata,
		TX_FILEBANK_UPLOADDEC,
		hash,
		dealinfo,
		user,
	)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewCall]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		SYSTEM,
		ACCOUNT,
		c.keyring.PublicKey,
	)
	if err != nil {
		return txhash, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return txhash, ERR_RPC_EMPTY_VALUE
	}

	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.runtimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.runtimeVersion.TransactionVersion,
	}

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		return txhash, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
	}
	defer sub.Unsubscribe()
	timeout := time.NewTimer(c.timeForBlockOut)
	defer timeout.Stop()
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				events := EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				h, err := c.api.RPC.State.GetStorageRaw(c.keyEvents, status.AsInBlock)
				if err != nil {
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}

				types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)

				if len(events.FileBank_UploadDeclaration) > 0 {
					return txhash, nil
				}
				return txhash, errors.New(ERR_Failed)
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) DeleteFile(owner_pkey []byte, filehash string) (string, FileHash, error) {
	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return txhash, FileHash{}, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	// var hash = make([]FileHash, len(filehash))

	// for j := 0; j < len(filehash); j++ {
	// 	if len(filehash[j]) != len(hash[j]) {
	// 		return txhash, FileHash{}, errors.New("invalid filehash")
	// 	}
	// 	for i := 0; i < len(hash[j]); i++ {
	// 		hash[j][i] = types.U8(filehash[j][i])
	// 	}
	// }

	var hash FileHash
	for i := 0; i < len(filehash); i++ {
		hash[i] = types.U8(filehash[i])
	}

	acc, err := types.NewAccountID(owner_pkey)
	if err != nil {
		return txhash, FileHash{}, errors.Wrap(err, "[NewAccountID]")
	}

	call, err := types.NewCall(
		c.metadata,
		TX_FILEBANK_DELFILE,
		*acc,
		hash,
	)
	if err != nil {
		return txhash, FileHash{}, errors.Wrap(err, "[NewCall]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		SYSTEM,
		ACCOUNT,
		c.keyring.PublicKey,
	)
	if err != nil {
		return txhash, FileHash{}, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, FileHash{}, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return txhash, hash, ERR_RPC_EMPTY_VALUE
	}

	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.runtimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.runtimeVersion.TransactionVersion,
	}

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		return txhash, FileHash{}, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return txhash, FileHash{}, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
	}
	defer sub.Unsubscribe()
	timeout := time.NewTimer(c.timeForBlockOut)
	defer timeout.Stop()
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				events := EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				h, err := c.api.RPC.State.GetStorageRaw(c.keyEvents, status.AsInBlock)
				if err != nil {
					return txhash, FileHash{}, errors.Wrap(err, "[GetStorageRaw]")
				}

				types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)

				if len(events.FileBank_DeleteFile) > 0 {
					return txhash, events.FileBank_DeleteFile[0].Filehash, nil
				}
				return txhash, FileHash{}, errors.New(ERR_Failed)
			}
		case err = <-sub.Err():
			return txhash, FileHash{}, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, FileHash{}, ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) SubmitFileReport(roothash []FileHash) (string, []FileHash, error) {
	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return txhash, nil, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	call, err := types.NewCall(
		c.metadata,
		TX_FILEBANK_FILEREPORT,
		roothash,
	)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[NewCall]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		SYSTEM,
		ACCOUNT,
		c.keyring.PublicKey,
	)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return txhash, nil, ERR_RPC_EMPTY_VALUE
	}

	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.runtimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.runtimeVersion.TransactionVersion,
	}

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
	}
	defer sub.Unsubscribe()
	timeout := time.NewTimer(c.timeForBlockOut)
	defer timeout.Stop()
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				events := EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				h, err := c.api.RPC.State.GetStorageRaw(c.keyEvents, status.AsInBlock)
				if err != nil {
					return txhash, nil, errors.Wrap(err, "[GetStorageRaw]")
				}

				types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)

				if len(events.FileBank_TransferReport) > 0 {
					return txhash, events.FileBank_TransferReport[0].Failed_list, nil
				}
				return txhash, nil, errors.New(ERR_Failed)
			}
		case err = <-sub.Err():
			return txhash, nil, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, nil, ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) ReplaceIdleFiles(roothash []FileHash) (string, []FileHash, error) {
	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return txhash, nil, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	call, err := types.NewCall(
		c.metadata,
		TX_FILEBANK_REPLACEFILE,
		roothash,
	)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[NewCall]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		SYSTEM,
		ACCOUNT,
		c.keyring.PublicKey,
	)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return txhash, nil, ERR_RPC_EMPTY_VALUE
	}

	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.runtimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.runtimeVersion.TransactionVersion,
	}

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return txhash, nil, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
	}
	defer sub.Unsubscribe()
	timeout := time.NewTimer(c.timeForBlockOut)
	defer timeout.Stop()
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				events := EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				h, err := c.api.RPC.State.GetStorageRaw(c.keyEvents, status.AsInBlock)
				if err != nil {
					return txhash, nil, errors.Wrap(err, "[GetStorageRaw]")
				}

				types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)

				if len(events.FileBank_ReplaceFiller) > 0 {
					return txhash, events.FileBank_ReplaceFiller[0].Filler_list, nil
				}
				return txhash, nil, errors.New(ERR_Failed)
			}
		case err = <-sub.Err():
			return txhash, nil, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, nil, ERR_RPC_TIMEOUT
		}
	}
}
