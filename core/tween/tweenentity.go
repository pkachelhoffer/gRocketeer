package tween

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EaseEntity struct {
	ease *Ease
	val  *float64
}

func NewEaseLooper(ease *Ease, val *float64) *EaseEntity {
	return &EaseEntity{
		ease: ease,
		val:  val,
	}
}

func (e *EaseEntity) Init() {

}

func (e *EaseEntity) Update() bool {
	e.ease.Update()
	*e.val = e.ease.Value()
	if e.ease.IsDone() {
		return true
	}

	return false
}

func (e *EaseEntity) Draw(_ *ebiten.Image) {

}
