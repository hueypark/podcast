package resource

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hueypark/podcast/pkg/logger"
	"golang.org/x/exp/slog"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
)

var (
	Font font.Face

	ButtonImage *widget.ButtonImage
)

func init() {
	f, err := loadFont(14)
	if err != nil {
		logger.Error("loafFond failed", slog.Any("error", err))
		return
	}

	Font = f

	ButtonImage = loadButtonImage()
}

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(gomonobold.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}

func loadButtonImage() *widget.ButtonImage {
	idle := image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}
}
