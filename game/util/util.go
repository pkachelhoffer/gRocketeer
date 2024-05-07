package util

import (
	"gRocketeer/core"
	"github.com/jakecoffman/cp"
)

func ToVec(v core.Vec) cp.Vector {
	return cp.Vector{
		X: v.X,
		Y: v.Y,
	}
}

func FromVec(v cp.Vector) core.Vec {
	return core.NewVec(v.X, v.Y)
}
