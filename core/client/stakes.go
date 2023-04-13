package client

import (
	"fmt"
	"math/big"

	"github.com/CESSProject/sdk-go/core/chain"
)

func (c *Cli) IncreaseStakes(token string) (string, error) {
	tokens, ok := new(big.Int).SetString(token+chain.TokenPrecision_CESS, 10)
	if !ok {
		return "", fmt.Errorf("Invalid tokens: %s", token)
	}
	return c.Chain.IncreaseStakes(tokens)
}
