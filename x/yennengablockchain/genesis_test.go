package yennengablockchain_test

import (
	"testing"

	keepertest "github.com/absonkab/yennenga_blockchain/testutil/keeper"
	"github.com/absonkab/yennenga_blockchain/testutil/nullify"
	"github.com/absonkab/yennenga_blockchain/x/yennengablockchain"
	"github.com/absonkab/yennenga_blockchain/x/yennengablockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.YennengablockchainKeeper(t)
	yennengablockchain.InitGenesis(ctx, *k, genesisState)
	got := yennengablockchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
