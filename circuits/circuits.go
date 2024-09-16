package circuits

import (
	"github.com/brevis-network/brevis-sdk/sdk"
)

type AppCircuit struct{}

func (c *AppCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	// Our app is only ever going to use one transaction data at a time so
	// we can simply limit the max number of data for transaction to 1 and
	// 0 for all others
	return 0, 0, 1
}

func (c *AppCircuit) Define(api *sdk.CircuitAPI, in sdk.DataInput) error {
	// Making a datastream is not necessarily useful here since we only
	// want to access the first element. But we will learn about this
	// API later in another section.
	txs := sdk.NewDataStream(api, in.Transactions)

	// We only have one element
	tx := sdk.GetUnderlying(txs, 0)

	// This is our main check logic. This satisfies our spec of "there
	// exists a tx with nonce = 0"
	api.Uint248.AssertIsEqual(tx.Nonce, sdk.ConstUint248(0))

	// Output variables can be later accessed in our app contract
	api.OutputAddress(tx.From)
	api.OutputUint(64, tx.BlockNum)

	return nil
}

var _ sdk.AppCircuit = &AppCircuit{}
