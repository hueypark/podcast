package main

import (
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hueypark/podcast/pkg/logger"
	"github.com/hueypark/podcast/pkg/podcast"
)

func main() {
	defer func() {
		logger.Info("podcast finished")
	}()

	pc, err := podcast.New()
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
