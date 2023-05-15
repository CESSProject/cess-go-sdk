package chain

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

func (c *chainClient) Register(role string, puk []byte, income string, pledge uint64) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		err         error
		txhash      string
		pubkey      []byte
		minerinfo   MinerInfo
		acc         *types.AccountID
		call        types.Call
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, ERR_RPC_CONNECTION
	}

	var peerpuk PeerPuk
	if len(pubkey) != len(puk) {
		return txhash, fmt.Errorf("invalid pubkey: %v", pubkey)
	}
	for i := 0; i < len(pubkey); i++ {
		peerpuk[i] = types.U8(pubkey[i])
	}

	key, err := types.CreateStorageKey(c.metadata, SYSTEM, ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		return txhash, errors.Wrap(err, "[CreateStorageKey]")
	}

	switch role {
	case Role_OSS, Role_DEOSS, "deoss", "oss", "Deoss", "DeOSS":
		pk, err := c.QueryDeoss(c.keyring.PublicKey)
		if err != nil {
			if err.Error() != ERR_Empty {
				return txhash, err
			}
		} else {
			if !CompareSlice(pk, puk) {
				return c.updateAddress(key, role, peerpuk)
			}
			return "", nil
		}

		call, err = types.NewCall(c.metadata, TX_OSS_REGISTER, peerpuk)
		if err != nil {
			return txhash, errors.Wrap(err, "[NewCall]")
		}
	case Role_BUCKET, "SMINER", "bucket", "Bucket", "Sminer", "sminer":
		minerinfo, err = c.QueryStorageMiner(c.keyring.PublicKey)
		if err != nil {
			if err.Error() != ERR_Empty {
				return txhash, err
			}
		} else {
			if minerinfo.PeerPuk != peerpuk {
				return c.updateAddress(key, role, peerpuk)
			}
			acc, _ := utils.EncodePublicKeyAsCessAccount(minerinfo.BeneficiaryAcc[:])
			if acc != income {
				puk, err := utils.ParsingPublickey(income)
				if err != nil {
					return txhash, err
				}
				return c.updateIncomeAcc(key, puk)
			}
			return "", nil
		}

		pubkey, err = utils.ParsingPublickey(income)
		if err != nil {
			return txhash, errors.Wrap(err, "[DecodeToPub]")
		}
		acc, err = types.NewAccountID(pubkey)
		if err != nil {
			return txhash, errors.Wrap(err, "[NewAccountID]")
		}
		realTokens, ok := new(big.Int).SetString(strconv.FormatUint(pledge, 10)+TokenPrecision_CESS, 10)
		if !ok {
			return txhash, errors.New("[big.Int.SetString]")
		}
		call, err = types.NewCall(c.metadata, TX_SMINER_REGISTER, *acc, peerpuk, types.NewU128(*realTokens))
		if err != nil {
			return txhash, errors.Wrap(err, "[NewCall]")
		}
	default:
		return "", fmt.Errorf("invalid role name")
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
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil {
					return txhash, nil
				}
				switch role {
				case Role_OSS, Role_DEOSS, "deoss", "oss", "Deoss", "DeOSS":
					if len(events.Oss_OssRegister) > 0 {
						return txhash, nil
					}
				case Role_BUCKET, "SMINER", "bucket", "Bucket", "Sminer", "sminer":
					if len(events.Sminer_Registered) > 0 {
						return txhash, nil
					}
				default:
					return txhash, errors.New(ERR_Failed)
				}
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) UpdateAddress(role, multiaddr string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		err         error
		txhash      string
		call        types.Call
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, ERR_RPC_CONNECTION
	}

	switch role {
	case Role_OSS, Role_DEOSS, "deoss", "oss", "Deoss", "DeOSS":
		call, err = types.NewCall(c.metadata, TX_OSS_UPDATE, types.NewBytes([]byte(multiaddr)))
		if err != nil {
			return txhash, errors.Wrap(err, "[NewCall]")
		}
	case Role_BUCKET, "SMINER", "bucket", "Bucket", "Sminer", "sminer":
		call, err = types.NewCall(c.metadata, TX_SMINER_UPDATEPEERID, types.NewBytes([]byte(multiaddr)))
		if err != nil {
			return txhash, errors.Wrap(err, "[NewCall]")
		}
	default:
		return "", fmt.Errorf("Invalid role name")
	}

	key, err := types.CreateStorageKey(c.metadata, SYSTEM, ACCOUNT, c.keyring.PublicKey)
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
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil {
					return txhash, nil
				}
				switch role {
				case Role_OSS, Role_DEOSS, "deoss", "oss", "Deoss", "DeOSS":
					if len(events.Oss_OssUpdate) > 0 {
						return txhash, nil
					}
				case Role_BUCKET, "SMINER", "bucket", "Bucket", "Sminer", "sminer":
					if len(events.Sminer_UpdataIp) > 0 {
						return txhash, nil
					}
				default:
					return txhash, errors.New(ERR_Failed)
				}
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) updateAddress(key types.StorageKey, name string, pubkey PeerPuk) (string, error) {
	var (
		err         error
		txhash      string
		call        types.Call
		accountInfo types.AccountInfo
	)

	switch name {
	case Role_OSS, Role_DEOSS, "deoss", "oss", "Deoss", "DeOSS":

		call, err = types.NewCall(c.metadata, TX_OSS_UPDATE, pubkey)
		if err != nil {
			return txhash, errors.Wrap(err, "[NewCall]")
		}
	case Role_BUCKET, "SMINER", "bucket", "Bucket", "Sminer", "sminer":
		call, err = types.NewCall(c.metadata, TX_SMINER_UPDATEPEERID, pubkey)
		if err != nil {
			return txhash, errors.Wrap(err, "[NewCall]")
		}
	default:
		return "", fmt.Errorf("Invalid role name")
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
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil {
					return txhash, nil
				}
				switch name {
				case Role_OSS, Role_DEOSS, "deoss", "oss", "Deoss", "DeOSS":
					if len(events.Oss_OssUpdate) > 0 {
						return txhash, nil
					}
				case Role_BUCKET, "SMINER", "bucket", "Bucket", "Sminer", "sminer":
					if len(events.Sminer_UpdataIp) > 0 {
						return txhash, nil
					}
				default:
					return txhash, errors.New(ERR_Failed)
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

func (c *chainClient) Exit(role string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		err         error
		txhash      string
		call        types.Call
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, ERR_RPC_CONNECTION
	}

	switch role {
	case Role_OSS, Role_DEOSS, "deoss", "oss", "Deoss", "DeOSS":
		call, err = types.NewCall(c.metadata, TX_OSS_DESTORY)
		if err != nil {
			return txhash, errors.Wrap(err, "[NewCall]")
		}
	case Role_BUCKET, "SMINER", "bucket", "Bucket", "Sminer", "sminer":
		call, err = types.NewCall(c.metadata, TX_FILEBANK_MINEREXITPREP)
		if err != nil {
			return txhash, errors.Wrap(err, "[NewCall]")
		}
	default:
		return "", fmt.Errorf("Invalid role name")
	}

	key, err := types.CreateStorageKey(c.metadata, SYSTEM, ACCOUNT, c.keyring.PublicKey)
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
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil {
					return txhash, nil
				}
				switch role {
				case Role_OSS, Role_DEOSS, "deoss", "oss", "Deoss", "DeOSS":
					if len(events.Oss_OssDestroy) > 0 {
						return txhash, nil
					}
				case Role_BUCKET, "SMINER", "bucket", "Bucket", "Sminer", "sminer":
					if len(events.Sminer_MinerExit) > 0 {
						return txhash, nil
					}
				default:
					return txhash, errors.New(ERR_Failed)
				}
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, ERR_RPC_TIMEOUT
		}
	}
}
