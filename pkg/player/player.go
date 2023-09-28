package player

import (
	"net/http"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type Player struct {
	audioPlayer *audio.Player
}

func New(audioContext *audio.Context, url string) (*Player, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	s, err := mp3.DecodeWithoutResampling(res.Body)
	if err != nil {
		return nil, err
	}

	ap, err := audioContext.NewPlayer(s)
	if err != nil {
		return nil, err
	}

	ap.Play()

	p := &Player{
		audioPlayer: ap,
	}

	return p, nil
}
