package view

import (
	"fmt"
	"strconv"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/ethn1ee/llog/internal/model"
	"github.com/fatih/color"
)

func PrintEntries(cfg *config.Config, entries []model.Entry) {
	maxIdDigits := 0

	for _, e := range entries {
		idDigits := len(strconv.FormatUint(e.ID, 10))
		if idDigits > maxIdDigits {
			maxIdDigits = idDigits
		}
	}

	for _, e := range entries {
		time := color.HiBlackString(e.CreatedAt.Format(cfg.TimeLayout))
		id := color.HiCyanString(fmt.Sprintf("[%*d]", maxIdDigits, e.ID))

		fmt.Printf("%s %s %s\n", id, time, e.Body)
	}
}

func PrintEntry(cfg *config.Config, entry model.Entry) {
	time := color.HiBlackString(entry.CreatedAt.Format(cfg.TimeLayout))
	id := color.HiCyanString(fmt.Sprintf("[%d]", entry.ID))

	fmt.Printf("%s %s %s\n", id, time, entry.Body)
}
