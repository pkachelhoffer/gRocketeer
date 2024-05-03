package cphys

import "playground/core"

type PhysParams struct {
	Fixed bool
	Rect  *core.Rectangle
}

type Actor interface {
	GetParams() *PhysParams
}

//type BodyType int
//
//const (
//	BodyTypeNone   BodyType = 0
//	BodyTypeRect   BodyType = 1
//	BodyTypeCircle BodyType = 2
//)
//
//type PhysBody interface {
//	GetBodyType() BodyType
//}
