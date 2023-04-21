package chain

import (
	"fmt"

	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

func (c *chainClient) QuerySpacePricePerGib() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data types.U128

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return "", ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		STORAGEHANDLER,
		UNITPRICE,
	)
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

	return fmt.Sprintf("%v", data), nil
}

func (c *chainClient) QueryUserSpaceInfo(puk []byte) (UserSpaceInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data UserSpaceInfo

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
		STORAGEHANDLER,
		USERSPACEINFO,
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
