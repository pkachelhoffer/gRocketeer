package cphys

import (
	"playground/core"
	"testing"
)

func TestCollideCircle(t *testing.T) {
	tcs := []struct {
		name  string
		circ1 Circle
		circ2 Circle
		exp   bool
	}{
		{
			name:  "just touching",
			circ1: NewCircle(core.NewVec(0, 0), 0.5),
			circ2: NewCircle(core.NewVec(1, 0), 0.5),
			exp:   true,
		},
		{
			name:  "just outside",
			circ1: NewCircle(core.NewVec(0, 0), 0.5),
			circ2: NewCircle(core.NewVec(1.001, 0), 0.5),
			exp:   false,
		},
		{
			name:  "outside 2",
			circ1: NewCircle(core.NewVec(0, 0), 2.5),
			circ2: NewCircle(core.NewVec(4, 3.01), 2.5),
			exp:   false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			collide := tc.circ1.CollideCirc(tc.circ2)
			if collide != tc.exp {
				t.Fatalf("exp collide %t", tc.exp)
			}
		})
	}
}
