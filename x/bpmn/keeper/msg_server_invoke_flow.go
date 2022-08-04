package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/louisliu2048/argus/x/bpmn/types"
)

func (k msgServer) InvokeFlow(goCtx context.Context, msg *types.MsgInvokeFlow) (*types.MsgInvokeFlowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgInvokeFlowResponse{}, nil
}
