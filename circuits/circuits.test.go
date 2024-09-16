package circuits

import (
	"fmt"
	"testing"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/ethereum/go-ethereum/common"

	"math/big"
)

// ...

func TestCircuit(t *testing.T) {
	app, err := sdk.NewBrevisApp()
	if err != nil {
		t.Fatalf("Failed to create Brevis app: %v", err)
	}

	txHash := common.HexToHash(
		"0x02869126ca667c76e819078d5326feb5d17f276ce5786de47e78334f15530e74")

	// code to get tx from ETH node omitted...
	// tx := queryTransactionFromEthNode(...)

	app.AddTransaction(sdk.TransactionData{
		Hash:                txHash,
		ChainId:             big.NewInt(0xaa36a7),
		BlockNum:            big.NewInt(6696020),
		Nonce:               5,
		GasTipCapOrGasPrice: big.NewInt(1952108673),
		GasFeeCap:           big.NewInt(2149579222),
		GasLimit:            26681,
		From:                common.HexToAddress("0x8efea8ce81c27fb620aa0c7796f6796aa0186d8f"),
		To:                  common.HexToAddress("0xa8f5dcc3035089111a9435ff25546c922a7c713a"),
		Value:               big.NewInt(0),
	})

	appCircuit := &AppCircuit{}
	appCircuitAssignment := &AppCircuit{}

	circuitInput, err := app.BuildCircuitInput(appCircuitAssignment)
	if err != nil {
		t.Fatalf("Failed to build circuit input: %v", err)
	}
	fmt.Println(circuitInput)
	fmt.Println(appCircuit)
	// test.ProverSucceeded(t, appCircuit, appCircuitAssignment, circuitInput)
}
