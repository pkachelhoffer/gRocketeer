package entity

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Init()
	Update() bool
	Draw(target *ebiten.Image)
	Terminate()
}
