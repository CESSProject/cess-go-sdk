package sdk

import (
	"cess-go-sdk/config"
	"cess-go-sdk/internal/chain"
	"cess-go-sdk/tools"
	"encoding/json"
	"github.com/pkg/errors"
)

type PurchaseSDK struct {
	config.CessConf
}
type TradeOperate interface {
	ObtainFromFaucet(string) error
	Expansion(int, int, int) error
}

type faucet struct {
	Ans    answer `json:"Result"`
	Status string `json:"Status"`
}
type answer struct {
	Err       string `json:"Err"`
	AsInBlock bool   `json:"AsInBlock"`
}

/*
ObtainFromFaucet means to obtain tCESS for transaction spending through the faucet
walletaddress:wallet address
*/
func (ts PurchaseSDK) ObtainFromFaucet(walletaddress string) error {
	pubkey, err := tools.DecodeToPub(walletaddress, tools.ChainCessTestPrefix)
	if err != nil {
		errors.Wrap(err, "[Error]The wallet address you entered is incorrect, please re-enter")
		return err
	}
	var ob = struct {
		Address string `json:"Address"`
	}{
		tools.PubBytesTo0XString(pubkey),
	}
	var res faucet
	resp, err := tools.Post(ts.ChainData.FaucetAddress, ob)
	if err != nil {
		return errors.Wrap(err, "[Error]System error")
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return errors.Wrap(err, "unmarshal error")
	}
	if res.Ans.Err != "" {
		err = errors.New(res.Ans.Err)
		return errors.Wrap(err, "[Error]Obtain from faucet fail")
	}

	if res.Ans.AsInBlock {
		return nil
	} else {
		return errors.New("[Fail]Obtain from faucet fail,Please wait 24 hours to get it again")
	}
}

/*
Expansion means the purchase of storage capacity for the current customer
quantity:The amount of space to be purchased (1/G)
duration:Market for space that needs to be purchased (1/month)
expected:The expected number of prices when buying is required to prevent price fluctuations when buying. When it is 0, it means that any price can be accepted
*/
func (ts PurchaseSDK) Expansion(quantity, duration, expected int) error {
	chain.Chain_Init(ts.ChainData.CessRpcAddr)
	if quantity == 0 || duration == 0 {
		return errors.New("Please enter the correct purchase number")
	}
	var ci chain.CessInfo
	ci.RpcAddr = ts.ChainData.CessRpcAddr
	ci.IdentifyAccountPhrase = ts.ChainData.IdAccountPhraseOrSeed
	ci.TransactionName = chain.BuySpaceTransactionName

	//Buying space on-chain, failure could mean running out of money
	err := ci.BuySpaceOnChain(quantity, duration, expected/1024)
	if err != nil {
		return errors.Wrap(err, "[Error] Failed to buy space, please check if you have enough money")
	}
	return nil
}
