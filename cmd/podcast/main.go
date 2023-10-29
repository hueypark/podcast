package main

import (
	"flag"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hueypark/podcast/pkg/logger"
	"github.com/hueypark/podcast/pkg/podcast"
)

var (
	// dataSourceName is the data source name for the sqlite database.
	dataSourceName string
)

func main() {
	flag.StringVar(&dataSourceName, "db", "podcast.db", "sqlite database file path")

	defer func() {
		logger.Info("podcast finished")
	}()

	pc, err := podcast.New(dataSourceName)
	if err != nil {
		logger.Error("create new podcast failed", slog.Any("error", err))
		return
	}

	err = ebiten.RunGame(pc)
	if err != nil {
		logger.Error("run game failed", slog.Any("error", err))
		return
	}
}
