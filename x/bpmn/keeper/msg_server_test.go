package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/louisliu2048/argus/testutil/keeper"
	"github.com/louisliu2048/argus/x/bpmn/keeper"
	"github.com/louisliu2048/argus/x/bpmn/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BpmnKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
