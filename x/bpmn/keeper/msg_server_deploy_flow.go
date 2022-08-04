package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/louisliu2048/argus/x/bpmn/types"
)

func (k msgServer) DeployFlow(goCtx context.Context, msg *types.MsgDeployFlow) (*types.MsgDeployFlowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeployFlowResponse{}, nil
}
