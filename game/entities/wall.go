package entities

import (
	"gRocketeer/core"
	"gRocketeer/core/entity"
	"gRocketeer/game/types"
	"gRocketeer/game/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"image/color"
)

var (
	whiteImage = ebiten.NewImage(1, 1)
)

func init() {
	whiteImage.Fill(color.White)
}

type Wall struct {
	entity.BaseEntity
	g  types.IGame
	sp *core.Sprite

	Col color.Color
}

func NewWall(g types.IGame, rect core.Rectangle, col color.Color) *Wall {
	w := &Wall{
		Col: col,
		sp:  core.NewSprite(whiteImage),
		g:   g,
	}

	w.sp.Rect = rect
	w.sp.Scale = rect.Bounds

	return w
}

func (w *Wall) Init() {
	b := cp.NewKinematicBody()
	b.SetPosition(util.ToVec(w.sp.Rect.Pos))
	s := cp.NewBox(b, w.sp.Rect.Bounds.X, w.sp.Rect.Bounds.Y, 0.5)
	s.SetElasticity(0.1)
	s.SetFriction(0.1)
	w.g.GetSpace().AddShape(s)
}

func (w *Wall) Update() bool {
	return false
}

func (w *Wall) Draw(target *ebiten.Image) {
	w.sp.Draw(target)
}
