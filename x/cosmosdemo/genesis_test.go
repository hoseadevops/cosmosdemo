package cosmosdemo_test

import (
	"testing"

	keepertest "github.com/hoseadevops/cosmosdemo/testutil/keeper"
	"github.com/hoseadevops/cosmosdemo/testutil/nullify"
	"github.com/hoseadevops/cosmosdemo/x/cosmosdemo"
	"github.com/hoseadevops/cosmosdemo/x/cosmosdemo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosdemoKeeper(t)
	cosmosdemo.InitGenesis(ctx, *k, genesisState)
	got := cosmosdemo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
