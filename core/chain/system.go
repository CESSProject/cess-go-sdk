package chain

import (
	"log"

	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

// QueryNodeSynchronizationSt
func (c *chainClient) QueryNodeSynchronizationSt() (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetChainState() {
		return false, ERR_RPC_CONNECTION
	}
	h, err := c.api.RPC.System.Health()
	if err != nil {
		return false, err
	}
	return h.IsSyncing, nil
}

// QueryBlockHeight
func (c *chainClient) QueryBlockHeight(hash string) (uint32, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if hash != "" {
		var h types.Hash
		err := codec.DecodeFromHex(hash, &h)
		if err != nil {
			return 0, err
		}
		block, err := c.api.RPC.Chain.GetBlock(h)
		if err != nil {
			return 0, errors.Wrap(err, "[GetBlock]")
		}
		return uint32(block.Block.Header.Number), nil
	}

	block, err := c.api.RPC.Chain.GetBlockLatest()
	if err != nil {
		return 0, errors.Wrap(err, "[GetBlockLatest]")
	}
	return uint32(block.Block.Header.Number), nil
}

// QueryAccountInfo
func (c *chainClient) QueryAccountInfo(puk []byte) (types.AccountInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.AccountInfo

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

	key, err := types.CreateStorageKey(c.metadata, SYSTEM, ACCOUNT, b)
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

func (c *chainClient) SysProperties() (SysProperties, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data SysProperties
	if !c.GetChainState() {
		return data, ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, RPC_SYS_Properties)
	return data, err
}

func (c *chainClient) SyncState() (SysSyncState, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data SysSyncState
	if !c.GetChainState() {
		return data, ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, RPC_SYS_SyncState)
	return data, err
}

func (c *chainClient) SysVersion() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Text
	if !c.GetChainState() {
		return "", ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, RPC_SYS_Version)
	return string(data), err
}

func (c *chainClient) NetListening() (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Bool
	if !c.GetChainState() {
		return false, ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, RPC_NET_Listening)
	return bool(data), err
}
