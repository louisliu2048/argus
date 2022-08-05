package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeployFlow{}, "bpmn/DeployFlow", nil)
	cdc.RegisterConcrete(&MsgStartFlow{}, "bpmn/StartFlow", nil)
	cdc.RegisterConcrete(&MsgInvokeFlow{}, "bpmn/InvokeFlow", nil)
	cdc.RegisterConcrete(&MsgCallEvm{}, "bpmn/CallEvm", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeployFlow{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStartFlow{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInvokeFlow{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCallEvm{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
