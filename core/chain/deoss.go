package chain

import (
	"log"

	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

// QueryDeoss
func (c *chainClient) QueryDeoss(pubkey []byte) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Panicln(utils.RecoverError(err))
		}
	}()
	var data types.Bytes

	if !c.GetChainState() {
		return "", ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, OSS, OSS, pubkey)
	if err != nil {
		return "", errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return "", errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return "", ERR_RPC_EMPTY_VALUE
	}
	return string(data), nil
}

func (c *chainClient) QuaryAuthorizedAcc(puk []byte) (types.AccountID, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.AccountID

	if !c.GetChainState() {
		return data, ERR_RPC_CONNECTION
	}

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	b, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, OSS, AUTHORITYLIST, b)
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
