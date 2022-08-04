package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartFlow = "start_flow"

var _ sdk.Msg = &MsgStartFlow{}

func NewMsgStartFlow(creator string, flowId string, initData string) *MsgStartFlow {
	return &MsgStartFlow{
		Creator:  creator,
		FlowId:   flowId,
		InitData: initData,
	}
}

func (msg *MsgStartFlow) Route() string {
	return RouterKey
}

func (msg *MsgStartFlow) Type() string {
	return TypeMsgStartFlow
}

func (msg *MsgStartFlow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartFlow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartFlow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
