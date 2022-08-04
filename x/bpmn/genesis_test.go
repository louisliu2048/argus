package bpmn_test

import (
	"testing"

	keepertest "github.com/louisliu2048/argus/testutil/keeper"
	"github.com/louisliu2048/argus/testutil/nullify"
	"github.com/louisliu2048/argus/x/bpmn"
	"github.com/louisliu2048/argus/x/bpmn/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BpmnKeeper(t)
	bpmn.InitGenesis(ctx, *k, genesisState)
	got := bpmn.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
