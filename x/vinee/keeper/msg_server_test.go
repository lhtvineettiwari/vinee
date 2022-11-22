package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/lhtvineettiwari/vinee/testutil/keeper"
	"github.com/lhtvineettiwari/vinee/x/vinee/keeper"
	"github.com/lhtvineettiwari/vinee/x/vinee/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VineeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
