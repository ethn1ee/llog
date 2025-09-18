package view

import (
	"fmt"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/ethn1ee/llog/internal/model"
	"github.com/fatih/color"
)

func PrintEntries(cfg *config.Config, entries []model.Entry) {
	for _, e := range entries {
		PrintEntry(cfg, e)
	}
}

func PrintEntry(cfg *config.Config, entry model.Entry) {
	time := color.HiBlackString(entry.CreatedAt.Format(cfg.TimeLayout))
	id := color.HiCyanString(fmt.Sprintf("[%*d]", cfg.Internal.MaxIdDigits, entry.ID))

	fmt.Printf("%s %s %s\n", id, time, entry.Body)
}
