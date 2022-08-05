package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/louisliu2048/argus/x/bpmn/types"
	"math/big"
)

func (k msgServer) CallEvm(goCtx context.Context, msg *types.MsgCallEvm) (*types.MsgCallEvmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// current tx info
	owner, _ := sdk.AccAddressFromBech32(msg.Creator) // has been checked in anteHandler
	fromAddr := common.BytesToAddress(owner.Bytes())
	userNonce := ctx.TxNonce()
	gasPrice := ctx.TxGasPrice()
	txGas := ctx.TxGas()

	// dynamic evm tx info
	contract, _ := sdk.AccAddressFromBech32(msg.Contract)
	contractAddr := common.BytesToAddress(contract.Bytes())
	value := big.NewInt(0) // forbidden to transfer token
	callData := common.FromHex(msg.CallData)

	fmt.Println()
	evmMsg := types2.NewMessage(fromAddr, &contractAddr, userNonce, value, txGas, gasPrice, gasPrice, gasPrice, callData, nil, false)

	tmpCtx, commit := ctx.CacheContext()
	rsp, err := k.evmKeeper.ApplyBpmnMessage(tmpCtx, evmMsg, common.BytesToHash(ctx.TxBytes()))
	if err != nil {
		return nil, sdkerrors.Wrap(err, "fail to invoke eth contract")
	}
	if rsp.Failed() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, rsp.VmError+"fail to invoke eth contract")
	}
	commit()

	return &types.MsgCallEvmResponse{}, nil
}
