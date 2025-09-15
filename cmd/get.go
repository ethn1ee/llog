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
	getCmd.Flags().BoolVarP(&(getOpts.Today), "today", "t", false, "get today's entries")
	getCmd.Flags().BoolVarP(&(getOpts.Yesterday), "yesterday", "y", false, "get yesterday's entries")
	getCmd.Flags().StringVar(&(getOpts.From), "from", "", "date in YYYY-MM-DD format")
	getCmd.Flags().StringVar(&(getOpts.To), "to", "", "date in YYYY-MM-DD format")

	rootCmd.AddCommand(getCmd)
}