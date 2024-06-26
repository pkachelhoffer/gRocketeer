package core

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

func LoadImageBytes(bs []byte) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func LoadShaderBytes(bs []byte) *ebiten.Shader {
	s, err := ebiten.NewShader(bs)
	if err != nil {
		log.Fatal(err)
	}

	return s
}
