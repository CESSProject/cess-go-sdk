package chain

import (
	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

func (c *chainClient) QueryChallengeSnapshot() (ChallengeSnapShot, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data ChallengeSnapShot

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		NETSNAPSHOT,
		CHALLENGESNAPSHOT,
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
