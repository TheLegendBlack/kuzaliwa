package keeper_test

import (
	"context"
	"testing"

	keepertest "Kuzaliwa/testutil/keeper"
	"Kuzaliwa/x/kuzaliwa/keeper"
	"Kuzaliwa/x/kuzaliwa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.KuzaliwaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
