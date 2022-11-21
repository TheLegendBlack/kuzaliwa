package keeper_test

import (
	"testing"

	testkeeper "Kuzaliwa/testutil/keeper"
	"Kuzaliwa/x/kuzaliwa/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.KuzaliwaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
