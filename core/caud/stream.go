package caud

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"log"
)

func CreateStream(bts []byte) *vorbis.Stream {
	s, err := vorbis.DecodeWithoutResampling(bytes.NewReader(bts))
	if err != nil {
		log.Fatal(err)
	}

	return s
}
