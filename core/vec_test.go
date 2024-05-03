package core

import (
	"testing"
)

//func TestIntersect(t *testing.T) {
//	tcs := []struct {
//		name string
//		r1   Rect
//		r2   Rect
//		exp  Rect
//	}{
//		{
//			name: "1",
//			r1:   NewRectMM(NewVec(0, 0), NewVec(10, 10)),
//			r2:   NewRectMM(NewVec(-2, -2), NewVec(10, 10)),
//			exp:  NewRectMM(NewVec(-1, -1), NewVec(8, 8)),
//		},
//		{
//			name: "2",
//			r1:   NewRectMM(NewVec(0, 0), NewVec(10, 10)),
//			r2:   NewRectMM(NewVec(2, -2), NewVec(10, 10)),
//			exp:  NewRectMM(NewVec(1, -1), NewVec(8, 8)),
//		},
//		{
//			name: "3",
//			r1:   NewRectMM(NewVec(0, 0), NewVec(10, 10)),
//			r2:   NewRectMM(NewVec(2, 2), NewVec(10, 10)),
//			exp:  NewRectMM(NewVec(1, 1), NewVec(8, 8)),
//		},
//		{
//			name: "4",
//			r1:   NewRectMM(NewVec(0, 0), NewVec(10, 10)),
//			r2:   NewRectMM(NewVec(-2, 2), NewVec(10, 10)),
//			exp:  NewRectMM(NewVec(-1, 1), NewVec(8, 8)),
//		},
//		{
//			name: "5",
//			r1:   NewRectMM(NewVec(0, 0), NewVec(10, 10)),
//			r2:   NewRectMM(NewVec(1, 1), NewVec(2, 2)),
//			exp:  NewRectMM(NewVec(1, 1), NewVec(2, 2)),
//		},
//		{
//			name: "6",
//			r1:   NewRectMM(NewVec(0, 0), NewVec(2, 2)),
//			r2:   NewRectMM(NewVec(5, 5), NewVec(2, 2)),
//			exp:  NewRectMM(NewVec(2.5, 2.5), NewVec(-3, -3)),
//		},
//		{
//			name: "7",
//			r1:   NewRectMM(NewVec(0, 0), NewVec(2, 2)),
//			r2:   NewRectMM(NewVec(2, 0), NewVec(2, 2)),
//			exp:  NewRectMM(NewVec(1, 0), NewVec(0, 2)),
//		},
//	}
//
//	for _, tc := range tcs {
//		t.Run(tc.name, func(t *testing.T) {
//			inter := tc.r1.Intersect(tc.r2)
//			if inter != tc.exp {
//				t.Errorf("intersect not as expected, exp: %s, act: %s", tc.exp, inter)
//			}
//
//			inter = tc.r2.Intersect(tc.r1)
//			if inter != tc.exp {
//				t.Errorf("intersect not as expected, exp: %s, act: %s", tc.exp, inter)
//			}
//		})
//	}
//}

func TestDistance(t *testing.T) {
	vec1 := NewVec(0, 0)
	vec2 := NewVec(40, 30)

	dist := vec1.Distance(vec2)
	if dist != 50 {
		t.Fatalf("fail!")
	}
}
