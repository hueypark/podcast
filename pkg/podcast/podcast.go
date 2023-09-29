package podcast

import (
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hueypark/podcast/pkg/dummy"
	"github.com/hueypark/podcast/pkg/feed"
	"github.com/hueypark/podcast/pkg/player"
	"github.com/hueypark/podcast/pkg/ui/page"
)

const (
	screenWidth  = 640
	screenHeight = 480

	audioSampleRate = 48000
)

type Podcast struct {
	feeds []feed.Feed

	audioContext *audio.Context
	player       *player.Player

	ui *ebitenui.UI
}

func New() (*Podcast, error) {
	const rssFeedURL = "https://feeds.transistor.fm/cup-o-go"

	audioContext := audio.NewContext(audioSampleRate)

	//fp := gofeed.NewParser()
	//fd, err := fp.ParseURL(rssFeedURL)
	//if err != nil {
	//	return nil, err
	//}

	//pcFeed := feed.MakeFeed(fd)

	//p, err := page.NewAddFeed()
	p, err := page.NewFeedList(dummy.Feeds)
	if err != nil {
		return nil, err
	}

	return &Podcast{
		//feeds:        []feed.Feed{pcFeed},
		audioContext: audioContext,
		ui:           p,
	}, nil
}

func (pc *Podcast) Update() error {
	pc.ui.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		pl, err := player.New(pc.audioContext, pc.feeds[0].Items[0].Enclosures[0].URL)
		if err != nil {
			return err
		}

		pc.player = pl
	}

	return nil
}

func (pc *Podcast) Draw(screen *ebiten.Image) {
	pc.ui.Draw(screen)
}

func (pc *Podcast) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
