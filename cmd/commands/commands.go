package commands

import (
	"github.com/spf13/cobra"

	"dev.vitoremanoel.tech/mobolife/modularize/cmd/commands/server"
)

func NewCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(server.Command)
}
