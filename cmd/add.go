/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"github.com/ethn1ee/llog/internal/handler"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a log entry",
	Long:  `Add a log entry to the database.`,
	Args:  cobra.MinimumNArgs(1),
	RunE:  withLog(add),
}

func add(cmd *cobra.Command, args []string) error {
	handler, err := handler.New(cmd)
	if err != nil {
		return err
	}
	return handler.AddEntry(cmd, args)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
