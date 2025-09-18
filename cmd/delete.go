/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"github.com/ethn1ee/llog/internal/handler"
	"github.com/spf13/cobra"
)

var deleteOpts = &handler.DeleteOpts{}

var deleteCmd = &cobra.Command{
	Use:          "delete [id]",
	Short:        "Delete log entries",
	Long:         `Delete log entries. You can delete with entry id or interactively select them.`,
	Args:         cobra.ExactArgs(1),
	PreRunE:      handler.ValidateOptions(cfg, deleteOpts),
	RunE:         handler.Delete(cfg, db, deleteOpts),
	SilenceUsage: true,
}

func init() {
	handler.ApplyFlags(deleteCmd, deleteOpts)
	rootCmd.AddCommand(deleteCmd)
}
