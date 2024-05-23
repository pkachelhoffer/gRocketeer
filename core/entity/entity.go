package entity

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Init()
	Update() bool
	Draw(target *ebiten.Image)
	Terminate()
}

type BaseEntity struct {
	FnTerminate []func()
}

func (b BaseEntity) Init() {

}

func (b BaseEntity) Update() bool {
	return false
}

func (b BaseEntity) Draw(target *ebiten.Image) {

}

func (b BaseEntity) AddTerminate(fn func()) {
	b.FnTerminate = append(b.FnTerminate, fn)
}

func (b BaseEntity) Terminate() {
	for _, fn := range b.FnTerminate {
		fn()
	}
}
