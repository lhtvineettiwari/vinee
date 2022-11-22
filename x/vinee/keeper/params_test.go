package keeper_test

import (
	"testing"

	testkeeper "github.com/lhtvineettiwari/vinee/testutil/keeper"
	"github.com/lhtvineettiwari/vinee/x/vinee/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VineeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
