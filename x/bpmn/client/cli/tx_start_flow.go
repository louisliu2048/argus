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

func CmdStartFlow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-flow [flow-id] [init-data]",
		Short: "Broadcast message start-flow",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFlowId := args[0]
			argInitData := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStartFlow(
				clientCtx.GetFromAddress().String(),
				argFlowId,
				argInitData,
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
