package view

import (
	"os"
	"strconv"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/ethn1ee/llog/internal/model"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

func PrintEntries(cfg *config.Config, entries []model.Entry) {
	data := make([][]string, len(entries))
	for i, e := range entries {
		data[i] = []string{
			color.HiCyanString(strconv.FormatUint(e.ID, 10)),
			color.HiBlackString(e.CreatedAt.Format(cfg.TimeLayout)),
			e.Body,
		}
	}

	symbols := tw.NewSymbolCustom("minimal").
		WithRow("").
		WithColumn("")

	table := tablewriter.NewTable(os.Stdout, tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{Symbols: symbols})))
	_ = table.Bulk(data)
	_ = table.Render()
}
