package keeper_test

import (
	"testing"

	testkeeper "github.com/absonkab/yennenga_blockchain/testutil/keeper"
	"github.com/absonkab/yennenga_blockchain/x/yennengablockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.YennengablockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
