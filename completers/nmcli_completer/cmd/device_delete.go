package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var device_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the software devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_deleteCmd).Standalone()

	deviceCmd.AddCommand(device_deleteCmd)

	carapace.Gen(device_deleteCmd).PositionalAnyCompletion(action.ActionDevices(""))
}
