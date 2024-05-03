package core

import "testing"

func TestRectCollide(t *testing.T) {
	tcs := []struct {
		name  string
		rect1 Rectangle
		rect2 Rectangle
		exp   bool
	}{
		{
			name:  "1",
			rect1: NewRect(NewVec(0, 0), NewVec(5, 5)),
			rect2: NewRect(NewVec(-2, -2), NewVec(5, 5)),
			exp:   true,
		},
		{
			name:  "2",
			rect1: NewRect(NewVec(0, 0), NewVec(5, 5)),
			rect2: NewRect(NewVec(2, -2), NewVec(5, 5)),
			exp:   true,
		},
		{
			name:  "3",
			rect1: NewRect(NewVec(0, 0), NewVec(5, 5)),
			rect2: NewRect(NewVec(2, 2), NewVec(5, 5)),
			exp:   true,
		},
		{
			name:  "4",
			rect1: NewRect(NewVec(0, 0), NewVec(5, 5)),
			rect2: NewRect(NewVec(-2, 2), NewVec(5, 5)),
			exp:   true,
		},
		{
			name:  "sides touching",
			rect1: NewRect(NewVec(0, 0), NewVec(5, 5)),
			rect2: NewRect(NewVec(5, 0), NewVec(5, 5)),
			exp:   true,
		},
		{
			name:  "just over",
			rect1: NewRect(NewVec(0, 0), NewVec(5, 5)),
			rect2: NewRect(NewVec(5.1, 0), NewVec(5, 5)),
			exp:   false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			collide := tc.rect1.CollideRect(tc.rect2)
			if collide != tc.exp {
				t.Fatalf("exp collide %t", tc.exp)
			}
		})
	}
}
