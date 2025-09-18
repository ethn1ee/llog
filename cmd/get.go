/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"github.com/ethn1ee/llog/internal/handler"
	"github.com/spf13/cobra"
)

var getOpts = &handler.GetOpts{}

var getCmd = &cobra.Command{
	Use:   "get [ID]",
	Short: "Get log entries",
	Long: `Get log entries.

You can retrieve a single entry by providing its ID as an argument.
Alternatively, you can retrieve a list of entries by using flags to filter by date range or limit the count.

Note: Providing an ID is mutually exclusive with using any flags.`,
	Args:         cobra.MaximumNArgs(1),
	PreRunE:      handler.ValidateOptions(cfg, getOpts),
	RunE:         handler.Get(cfg, db, getOpts),
	SilenceUsage: true,
}

func init() {
	handler.ApplyFlags(getCmd, getOpts)
	rootCmd.AddCommand(getCmd)
}
