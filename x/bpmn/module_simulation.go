package bpmn

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/louisliu2048/argus/testutil/sample"
	bpmnsimulation "github.com/louisliu2048/argus/x/bpmn/simulation"
	"github.com/louisliu2048/argus/x/bpmn/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = bpmnsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgDeployFlow = "op_weight_msg_deploy_flow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeployFlow int = 100

	opWeightMsgStartFlow = "op_weight_msg_start_flow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStartFlow int = 100

	opWeightMsgInvokeFlow = "op_weight_msg_invoke_flow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInvokeFlow int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bpmnGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bpmnGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgDeployFlow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeployFlow, &weightMsgDeployFlow, nil,
		func(_ *rand.Rand) {
			weightMsgDeployFlow = defaultWeightMsgDeployFlow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeployFlow,
		bpmnsimulation.SimulateMsgDeployFlow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStartFlow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStartFlow, &weightMsgStartFlow, nil,
		func(_ *rand.Rand) {
			weightMsgStartFlow = defaultWeightMsgStartFlow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartFlow,
		bpmnsimulation.SimulateMsgStartFlow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInvokeFlow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInvokeFlow, &weightMsgInvokeFlow, nil,
		func(_ *rand.Rand) {
			weightMsgInvokeFlow = defaultWeightMsgInvokeFlow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInvokeFlow,
		bpmnsimulation.SimulateMsgInvokeFlow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
