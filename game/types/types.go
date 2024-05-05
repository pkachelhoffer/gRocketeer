package types

import (
	"gRocketeer/core"
	"gRocketeer/core/entity"
)

type IGame interface {
	Sys() *entity.System
	Screen() core.Vec
}
