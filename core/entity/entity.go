package entity

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Init()
	Update() bool
	Draw(target *ebiten.Image)
	Terminate()
}

type BaseEntity struct {
}

func (b BaseEntity) Init() {

}

func (b BaseEntity) Update() bool {
	return false
}

func (b BaseEntity) Draw(target *ebiten.Image) {

}

func (b BaseEntity) Terminate() {

}
