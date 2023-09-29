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

func NewEpisodeList(episodes []feed.Episode) (*ebitenui.UI, error) {
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

	for _, episode := range episodes {
		fc, err := newEpisodeContainer(episode)
		if err != nil {
			return nil, fmt.Errorf("newFeedContainer %s failed, error: %w", episode.Title, err)
		}

		rootContainer.AddChild(fc)
	}

	return &ebitenui.UI{
		Container: rootContainer,
	}, nil
}

func newEpisodeContainer(episode feed.Episode) (*widget.Container, error) {
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

					logger.Info("Episode clicked", slog.String("title", episode.Title))
				},
			),
		),
	)

	title := widget.NewText(
		widget.TextOpts.Text(
			episode.Title,
			resource.Font,
			color.Black,
		),
	)
	feedContainer.AddChild(title)

	return feedContainer, nil
}
