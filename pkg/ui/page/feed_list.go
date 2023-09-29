package page

import (
	"fmt"
	"image/color"
	"log/slog"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hueypark/podcast/pkg/feed"
	"github.com/hueypark/podcast/pkg/logger"
	"github.com/hueypark/podcast/pkg/ui/resource"
)

func NewFeedList(feeds []feed.Feed) (*ebitenui.UI, error) {
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

	for _, fd := range feeds {
		fc, err := newFeedContainer(fd)
		if err != nil {
			return nil, fmt.Errorf("newFeedContainer %s failed, error: %w", fd.Title, err)
		}

		rootContainer.AddChild(fc)
	}

	return &ebitenui.UI{
		Container: rootContainer,
	}, nil
}

func newFeedContainer(fd feed.Feed) (*widget.Container, error) {
	feedContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.White)),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
				widget.RowLayoutOpts.Spacing(10),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
					Stretch:  true,
				},
			),
			widget.WidgetOpts.MouseButtonReleasedHandler(
				func(args *widget.WidgetMouseButtonReleasedEventArgs) {
					if !args.Inside {
						return
					}

					logger.Info("Feed clicked", slog.String("title", fd.Title))
				},
			),
		),
	)

	title := widget.NewText(
		widget.TextOpts.Text(
			fd.Title,
			resource.Font,
			color.Black,
		),
	)
	feedContainer.AddChild(title)

	return feedContainer, nil
}
