package cphys

//
//import (
//	"fmt"
//	"playground/core"
//	"playground/core/config"
//)
//
//type WorldActor struct {
//	Actor Actor
//	Vel   core.Vec
//}
//
//type World struct {
//	actors map[string]*WorldActor
//
//	Gravity core.Vec
//}
//
//func NewWorld() *World {
//	return &World{
//		actors: make(map[string]*WorldActor),
//	}
//}
//
//func (w *World) AddActor(key string, a Actor) {
//	if config.Debug {
//		fmt.Println("w add", key)
//	}
//
//	w.actors[key] = &WorldActor{Actor: a}
//}
//
//func (w *World) RemoveActor(key string) {
//	_, ok := w.actors[key]
//	if ok {
//		if config.Debug {
//			fmt.Println("w remove", key)
//		}
//		delete(w.actors, key)
//	} else {
//		if config.Debug {
//			fmt.Println("w remove not found", key)
//		}
//	}
//}
//
//func (w *World) Update() {
//	for k1, a1 := range w.actors {
//		if a1.Actor.GetParams().Fixed {
//			continue
//		}
//
//		// Apply gravity to velocity
//		a1.Vel.X += core.PerSecond(w.Gravity.X)
//		a1.Vel.Y += core.PerSecond(w.Gravity.Y)
//
//		if a1.Vel.IsZero() {
//			continue
//		}
//
//		// Apply velocity to position
//		par := a1.Actor.GetParams()
//		par.Rect.Pos.X += a1.Vel.X
//		par.Rect.Pos.Y += a1.Vel.Y
//
//		// Check collisions
//		for k2, a2 := range w.actors {
//			if k1 == k2 {
//				continue
//			}
//
//			inter := a1.Actor.GetParams().Rect.Intersect(*a2.Actor.GetParams().Rect)
//			if !inter.Bounds.IsZero() && !inter.Bounds.IsNegative() {
//				if config.Debug {
//					fmt.Println("intersect", k1, k2, inter.Bounds)
//				}
//
//				if inter.Bounds.Y > 0 {
//					a1.Vel.Y = 0
//					if a1.Actor.GetParams().Rect.Pos.Y < inter.Pos.Y {
//						a1.Actor.GetParams().Rect.Pos.Y -= inter.Bounds.Y
//					} else {
//						a1.Actor.GetParams().Rect.Pos.Y += inter.Bounds.Y
//					}
//				} else if inter.Bounds.X > 0 {
//					a1.Vel.X = 0
//					if a1.Actor.GetParams().Rect.Pos.X < inter.Pos.X {
//						a1.Actor.GetParams().Rect.Pos.X -= inter.Bounds.X
//					} else {
//						a1.Actor.GetParams().Rect.Pos.X += inter.Bounds.X
//					}
//				}
//			}
//		}
//	}
//}
