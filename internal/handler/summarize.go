package handler

import (
	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/ethn1ee/llog/internal/logger"
	"github.com/spf13/cobra"
)

func Summarize(cfg *config.Config, db *_db.DB, opts *SummarizeOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		logger.LogCmdStart(cmd)
		defer logger.LogCmdComplete(cmd)

		return nil
	}
}

type SummarizeOpts struct {
	Time timeOpts
}

func (o *SummarizeOpts) applyFlags(cmd *cobra.Command) {
	o.Time.applyFlags(cmd)
}

func (o *SummarizeOpts) validate(cfg *config.Config, args []string, flags []string) error {
	return o.Time.validate(cfg, args, flags)
}
