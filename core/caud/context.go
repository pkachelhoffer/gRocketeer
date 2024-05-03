package caud

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"time"
)

var (
	aCtx *audio.Context
)

const sampleRate = 44800

func init() {
	aCtx = audio.NewContext(sampleRate)
	for !aCtx.IsReady() {
		time.Sleep(time.Millisecond * 10)
	}
}
