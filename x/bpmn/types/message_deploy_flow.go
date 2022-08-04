package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeployFlow = "deploy_flow"

var _ sdk.Msg = &MsgDeployFlow{}

func NewMsgDeployFlow(creator string, data string) *MsgDeployFlow {
	return &MsgDeployFlow{
		Creator: creator,
		Data:    data,
	}
}

func (msg *MsgDeployFlow) Route() string {
	return RouterKey
}

func (msg *MsgDeployFlow) Type() string {
	return TypeMsgDeployFlow
}

func (msg *MsgDeployFlow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeployFlow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeployFlow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
