package podcast

import (
	"database/sql"
	"fmt"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hueypark/podcast/pkg/dummy"
	"github.com/hueypark/podcast/pkg/feed"
	"github.com/hueypark/podcast/pkg/persistence"
	"github.com/hueypark/podcast/pkg/player"
	"github.com/hueypark/podcast/pkg/ui/page"
)

const (
	screenWidth  = 640
	screenHeight = 480

	audioSampleRate = 48000
)

type Podcast struct {
	db *sql.DB

	feeds []feed.Feed

	audioContext *audio.Context
	player       *player.Player

	ui *ebitenui.UI
}

func New(dataSourceName string) (*Podcast, error) {
	const rssFeedURL = "https://feeds.transistor.fm/cup-o-go"

	db, err := persistence.NewSqliteDB(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("create new sqlite database failed: %w", err)
	}

	audioContext := audio.NewContext(audioSampleRate)

	//fp := gofeed.NewParser()
	//fd, err := fp.ParseURL(rssFeedURL)
	//if err != nil {
	//	return nil, err
	//}

	//pcFeed := feed.MakeFeed(fd)

	//p, err := page.NewAddFeed()
	//p, err := page.NewFeedList(dummy.Feeds)
	p, err := page.NewEpisodeList(dummy.Episodes)
	if err != nil {
		return nil, err
	}

	return &Podcast{
		//feeds:        []feed.Feed{pcFeed},
		db:           db,
		audioContext: audioContext,
		ui:           p,
	}, nil
}

func (pc *Podcast) Update() error {
	pc.ui.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		pl, err := player.New(pc.audioContext, pc.feeds[0].Episodes[0].Enclosures[0].URL)
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
