package kuzaliwa_test

import (
	"testing"

	keepertest "Kuzaliwa/testutil/keeper"
	"Kuzaliwa/testutil/nullify"
	"Kuzaliwa/x/kuzaliwa"
	"Kuzaliwa/x/kuzaliwa/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KuzaliwaKeeper(t)
	kuzaliwa.InitGenesis(ctx, *k, genesisState)
	got := kuzaliwa.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
