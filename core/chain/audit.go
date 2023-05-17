package chain

import (
	"log"
	"time"

	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

func (c *chainClient) QueryAssignedProof() ([][]ProofAssignmentInfo, error) {
	var list [][]ProofAssignmentInfo
	key := createPrefixedKey(AUDIT, UNVERIFYPROOF)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		return list, errors.Wrap(err, "[GetKeysLatest]")
	}
	set, err := c.api.RPC.State.QueryStorageAtLatest(keys)

	if err != nil {
		return list, errors.Wrap(err, "[QueryStorageAtLatest]")
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			var data []ProofAssignmentInfo
			if err := codec.Decode(change.StorageData, &data); err != nil {
				log.Println(err)
				continue
			}
			list = append(list, data)
		}
	}
	return list, nil
}

func (c *chainClient) QueryTeeAssignedProof(puk []byte) ([]ProofAssignmentInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Panicln(utils.RecoverError(err))
		}
	}()
	var data []ProofAssignmentInfo

	if !c.GetChainState() {
		return data, ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, AUDIT, UNVERIFYPROOF, puk)
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

func (c *chainClient) ReportProof(idlesigma, servicesigma string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_AUDIT_SUBMITPROOF, types.NewBytes([]byte(idlesigma)), types.NewBytes([]byte(servicesigma)))
	if err != nil {
		return txhash, errors.Wrap(err, "[NewCall]")
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
				if err != nil || len(events.Audit_SubmitProof) > 0 {
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
