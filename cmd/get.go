/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"github.com/ethn1ee/llog/internal/handler"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get log entries",
	Long:  `Get log entries from the database. You can get a specific entry by its ID or get all entries.`,
}

var getAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all log entries",
	Long:  `Get all log entries from the database.`,
	Args:  cobra.NoArgs,
	RunE:  withLog(getAllFn),
}

func getAllFn(cmd *cobra.Command, args []string) error {
	handler, err := handler.New(cmd)
	if err != nil {
		return err
	}
	return handler.GetAllEntry(cmd, args)
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getAllCmd)
}
