package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/absonkab/yennenga_blockchain/testutil/keeper"
	"github.com/absonkab/yennenga_blockchain/x/yennengablockchain/keeper"
	"github.com/absonkab/yennenga_blockchain/x/yennengablockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.YennengablockchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
