package types

import (
	"gRocketeer/core"
	"gRocketeer/core/entity"
	"github.com/jakecoffman/cp"
)

type IGame interface {
	GetSys() *entity.System
	Screen() core.Vec
	GetSpace() *cp.Space
}
