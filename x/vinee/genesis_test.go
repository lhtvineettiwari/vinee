package vinee_test

import (
	"testing"

	keepertest "github.com/lhtvineettiwari/vinee/testutil/keeper"
	"github.com/lhtvineettiwari/vinee/testutil/nullify"
	"github.com/lhtvineettiwari/vinee/x/vinee"
	"github.com/lhtvineettiwari/vinee/x/vinee/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VineeKeeper(t)
	vinee.InitGenesis(ctx, *k, genesisState)
	got := vinee.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
