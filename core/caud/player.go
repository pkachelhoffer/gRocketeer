package caud

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"io"
	"log"
)

type Player struct {
	pl *audio.Player
}

func NewPlayer(s io.ReadSeeker) *Player {
	np := &Player{}
	pl, err := aCtx.NewPlayer(s.(io.Reader))
	if err != nil {
		log.Fatal(err)
	}

	np.pl = pl

	return np
}

func (p *Player) Play() {
	err := p.pl.Rewind()
	if err != nil {
		log.Fatal(err)
	}
	p.pl.Play()
}

func (p *Player) SetVolume(v float64) {
	p.pl.SetVolume(v)
}

func (p *Player) Pause() {
	p.pl.Pause()
}
