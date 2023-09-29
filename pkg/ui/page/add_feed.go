package page

import (
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
)

func NewAddFeed() (*ebitenui.UI, error) {
	face, err := loadFont(14)
	if err != nil {
		return nil, err
	}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.Black)),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
			),
		),
	)

	feedURLInput := widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
					Stretch:  true,
				},
			),
		),
		widget.TextInputOpts.MobileInputMode("text"),
		widget.TextInputOpts.Image(
			&widget.TextInputImage{
				Idle: image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
			},
		),
		widget.TextInputOpts.Face(face),
		widget.TextInputOpts.Color(
			&widget.TextInputColor{
				Idle:          color.White,
				Caret:         color.White,
				Disabled:      color.NRGBA{R: 200, G: 200, B: 200, A: 255},
				DisabledCaret: color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			},
		),
		widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),
		widget.TextInputOpts.CaretOpts(widget.CaretOpts.Size(face, 2)),
		widget.TextInputOpts.Placeholder("RSS feed URL"),
		widget.TextInputOpts.ClearOnSubmit(true),
		widget.TextInputOpts.IgnoreEmptySubmit(true),
		widget.TextInputOpts.SubmitHandler(
			func(args *widget.TextInputChangedEventArgs) {
				fmt.Println("Text Submitted: ", args.InputText)
			},
		),
	)
	rootContainer.AddChild(feedURLInput)

	button := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  false,
			}),
		),
		widget.ButtonOpts.Image(loadButtonImage()),
		widget.ButtonOpts.Text(
			"Add feed",
			face,
			&widget.ButtonTextColor{
				Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
			},
		),
		widget.ButtonOpts.TextPadding(
			widget.Insets{
				Left:   30,
				Right:  30,
				Top:    5,
				Bottom: 5,
			},
		),
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				feedURLInput.Submit()
			},
		),
	)
	rootContainer.AddChild(button)

	return &ebitenui.UI{
		Container: rootContainer,
	}, nil
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
