package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/louisliu2048/argus/x/bpmn/keeper"
	"github.com/louisliu2048/argus/x/bpmn/types"
)

func SimulateMsgInvokeFlow(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgInvokeFlow{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the InvokeFlow simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "InvokeFlow simulation not implemented"), nil, nil
	}
}
