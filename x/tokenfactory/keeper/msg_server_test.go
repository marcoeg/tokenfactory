package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/marcoeg/tokenfactory/testutil/keeper"
	"github.com/marcoeg/tokenfactory/x/tokenfactory/keeper"
	"github.com/marcoeg/tokenfactory/x/tokenfactory/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TokenfactoryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
