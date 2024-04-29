package chain

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

func (c *ChainClient) SendEvmCall(source types.H160, target types.H160, input types.Bytes, value types.U256, gasLimit types.U64, maxFeePerGas types.U256, accessList []pattern.AccessInfo) (string, error) {
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
		return txhash, fmt.Errorf("chainSDK.UploadDeclaration(): GetChainState(): %v", pattern.ERR_RPC_CONNECTION)
	}

	var nonce types.Option[types.U256]
	nonce.SetNone()

	var maxPriorityFeePerGas types.Option[types.U256]
	maxPriorityFeePerGas.SetNone()

	call, err := types.NewCall(c.metadata, pattern.TX_EVM_CALL, source, target, input, value, gasLimit, maxFeePerGas, maxPriorityFeePerGas, nonce, accessList)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), pattern.EVM, err)
		c.SetChainState(false)
		return txhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.EVM, err)
		c.SetChainState(false)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.EVM, err)
		c.SetChainState(false)
		return txhash, err
	}

	if !ok {
		keyStr, _ := utils.NumsToByteStr(key, map[string]bool{})
		return txhash, fmt.Errorf(
			"chain rpc.state.GetStorageLatest[%v]: %v",
			keyStr,
			pattern.ERR_RPC_EMPTY_VALUE,
		)
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

	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), pattern.EVM, err)
		c.SetChainState(false)
		return txhash, err
	}

	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), pattern.ERR_RPC_PRIORITYTOOLOW) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.EVM, err)
				c.SetChainState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.EVM, err)
			c.SetChainState(false)
			return txhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				txhash = status.AsInBlock.Hex()
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}
