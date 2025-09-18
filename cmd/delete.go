
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
	Use:     "delete",
	Short:   "Delete log entries",
	Long:    `Delete log entries. You can delete with id or interactively`,
	PreRunE: handler.ValidateOptions(cfg, deleteOpts),
	RunE:    handler.Delete(cfg, db, deleteOpts),
}

func init() {
	handler.ApplyFlags(deleteCmd, deleteOpts)
	rootCmd.AddCommand(deleteCmd)
}
