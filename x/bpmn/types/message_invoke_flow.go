package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInvokeFlow = "invoke_flow"

var _ sdk.Msg = &MsgInvokeFlow{}

func NewMsgInvokeFlow(creator string, instanceId string, callData string) *MsgInvokeFlow {
	return &MsgInvokeFlow{
		Creator:    creator,
		InstanceId: instanceId,
		CallData:   callData,
	}
}

func (msg *MsgInvokeFlow) Route() string {
	return RouterKey
}

func (msg *MsgInvokeFlow) Type() string {
	return TypeMsgInvokeFlow
}

func (msg *MsgInvokeFlow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInvokeFlow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInvokeFlow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
