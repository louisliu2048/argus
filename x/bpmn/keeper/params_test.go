package keeper_test

import (
	"testing"

	testkeeper "github.com/louisliu2048/argus/testutil/keeper"
	"github.com/louisliu2048/argus/x/bpmn/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BpmnKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
