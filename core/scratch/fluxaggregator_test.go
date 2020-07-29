package scratch

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/smartcontractkit/chainlink/core/scratch/gethwrappers/eacaggregatorproxy"
	"github.com/smartcontractkit/chainlink/core/scratch/gethwrappers/fluxaggregator"
	"github.com/stretchr/testify/require"
)

// newIdentity returns a go-ethereum abstraction of an ethereum account for
// interacting with contract golang wrappers
func newIdentity(t *testing.T) *bind.TransactOpts {
	key, err := crypto.GenerateKey()
	require.NoError(t, err, "failed to generate ethereum identity")
	return bind.NewKeyedTransactor(key)
}

func TestProxy(t *testing.T) {
	owner := newIdentity(t)
	oneEth := big.NewInt(1000000000000000000)
	genesisData := core.GenesisAlloc{owner.From: {Balance: oneEth}}
	gasLimit := eth.DefaultConfig.Miner.GasCeil
	backend := backends.NewSimulatedBackend(genesisData, gasLimit)
	aggregatorAddress, _, _, err := fluxaggregator.DeployFluxAggregator(owner, backend, common.Address{}, big.NewInt(0), 0, common.Address{}, big.NewInt(0), big.NewInt(0), 10, "")
	require.NoError(t, err)
	proxyAddress, _, _, err := eacaggregatorproxy.DeployAggregatorProxy(owner, backend, aggregatorAddress)
	require.NoError(t, err)
	abi, err := abi.JSON(strings.NewReader(fluxaggregator.FluxAggregatorABI))
	require.NoError(t, err)
	rawData, err := abi.Pack("latestRoundData")
	require.NoError(t, err)
	latestRoundDataRaw := ethereum.CallMsg{To: &proxyAddress, Data: rawData}
	latestRoundDataEstimate, err := backend.EstimateGas(context.TODO(), latestRoundDataRaw)
	require.NoError(t, err)
	fmt.Println("EstimateGas", latestRoundDataEstimate)
	latestRoundDataRaw.To = &aggregatorAddress
	estimate, err := backend.EstimateGas(context.TODO(), latestRoundDataRaw)
	require.NoError(t, err)
	fmt.Println("EstimateGas", estimate)
	latestAnswerRaw, err := abi.Pack("latestAnswer")
	require.NoError(t, err)
	latestAnswerCall := ethereum.CallMsg{To: &proxyAddress, Data: latestAnswerRaw}
	latestAnswerEstimate, err := backend.EstimateGas(context.TODO(), latestAnswerCall)
	require.NoError(t, err)
	fmt.Println("latestRoundData - latestAnswer", latestRoundDataEstimate-latestAnswerEstimate)
}
