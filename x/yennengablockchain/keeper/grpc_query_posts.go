package keeper

import (
	"context"

	"github.com/absonkab/yennenga_blockchain/x/yennengablockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Posts(goCtx context.Context, req *types.QueryPostsRequest) (*types.QueryPostsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryPostsResponse{Title: "Hello Texacoin team!", Body: "This is Yennenga Blockchain built with Ignite CLI"}, nil
}
