package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCallEvm = "call_evm"

var _ sdk.Msg = &MsgCallEvm{}

func NewMsgCallEvm(creator string, contract string, callData string) *MsgCallEvm {
	return &MsgCallEvm{
		Creator:  creator,
		Contract: contract,
		CallData: callData,
	}
}

func (msg *MsgCallEvm) Route() string {
	return RouterKey
}

func (msg *MsgCallEvm) Type() string {
	return TypeMsgCallEvm
}

func (msg *MsgCallEvm) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCallEvm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCallEvm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
