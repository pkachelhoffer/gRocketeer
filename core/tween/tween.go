package tween

import "math"

type EaseFunc func(i int, len int, start float64, end float64) float64

func Linear(i int, len int, start float64, end float64) float64 {
	perc := float64(i) / float64(len)
	return start + (perc * (end - start))
}

func QuadraticOut(i int, len int, start float64, end float64) float64 {
	perc := float64(i) / float64(len)
	change := end - start
	return start + change*math.Pow(perc, 2)
}

func QuadraticIn(i int, len int, start float64, end float64) float64 {
	perc := float64(i) / float64(len)
	change := end - start
	return start - change*perc*(perc-2)
}

type Ease struct {
	len   int
	i     int
	start float64
	end   float64
	val   float64

	f EaseFunc
}

func NewEase(len int, start float64, end float64, f EaseFunc) *Ease {
	return &Ease{
		len:   len,
		start: start,
		end:   end,
		f:     f,
		val:   start,
	}
}

func (e *Ease) Update() {
	if e.IsDone() {
		return
	}

	e.i++
	e.val = e.f(e.i, e.len, e.start, e.end)
}

func (e *Ease) Value() float64 {
	return e.val
}

func (e *Ease) IsDone() bool {
	return e.i >= e.len
}

func (e *Ease) Reset() {
	e.i = 0
}

func (e *Ease) Perc() float64 {
	return float64(e.i) / float64(e.len)
}

func (e *Ease) SetDone() {
	e.i = e.len
	e.val = e.f(e.i, e.len, e.start, e.end)
}
