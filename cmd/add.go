/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"github.com/ethn1ee/llog/internal/handler"
	"github.com/spf13/cobra"
)

var addOpts = &handler.AddOpts{}

var addCmd = &cobra.Command{
	Use:     "add [body]",
	Short:   "Add a log entry",
	Long:    `Add a log entry.`,
	Args:    cobra.ExactArgs(1),
	PreRunE: handler.ValidateOptions(cfg, addOpts),
	RunE:    handler.Add(cfg, db, addOpts),
}

func init() {
	rootCmd.AddCommand(addCmd)
}