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

func CmdCallEvm() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call-evm [contract-address] [call-data]",
		Short: "Broadcast message call-evm",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContractAddr := args[0]
			argCallData := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCallEvm(
				clientCtx.GetFromAddress().String(),
				argContractAddr,
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
