package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ethn1ee/llog/internal/config"
)

type Logger struct {
	file *os.File
}

func Init(cfg *config.Config, lg *Logger) (error) {
	file, err := os.OpenFile(cfg.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return  fmt.Errorf("failed to open log file; %w", err)
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})

	slog.SetDefault(slog.New(handler))

	lg.file = file

	return nil
}

func (l *Logger) Close() error {
	return l.file.Close()
}
