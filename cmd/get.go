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
	Use:     "get",
	Short:   "Get log entries",
	Long:    `Get log entries. You can specify date range with flags.`,
	Args:    cobra.NoArgs,
	PreRunE: handler.ValidateOptions(cfg, getOpts),
	RunE:    handler.Get(cfg, db, getOpts),
}

func init() {
	handler.ApplyFlags(getCmd, getOpts)
	rootCmd.AddCommand(getCmd)
}
