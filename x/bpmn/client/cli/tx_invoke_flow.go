package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/louisliu2048/argus/x/bpmn/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInvokeFlow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke-flow [instance-id] [call-data]",
		Short: "Broadcast message invoke-flow",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argInstanceId := args[0]
			argCallData := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInvokeFlow(
				clientCtx.GetFromAddress().String(),
				argInstanceId,
				argCallData,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
