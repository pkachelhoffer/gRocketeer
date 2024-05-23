package entities

import (
	"gRocketeer/core/cpc"
	"gRocketeer/core/entity"
	"gRocketeer/game/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

var (
	imgBullet *ebiten.Image
)

func init() {
	imgBullet = ebiten.NewImage(3, 3)
	imgBullet.Fill(types.ColorWhite)
}

type Bullet01 struct {
	entity.BaseEntity

	csp cpc.CSprite

	g types.IGame
}

func NewBullet01(g types.IGame) *Bullet01 {
	csp := cpc.NewCSprite(imgBullet, g.GetSpace(), 0.1, cp.INFINITY)
	csp.Shape.SetFilter(cp.ShapeFilter{Categories: types.CatBullet, Mask: types.CatBullet})

	return &Bullet01{
		csp: csp,
		g:   g,
	}
}

func (b *Bullet01) Update() bool {
	b.csp.Update()
	return false
}

func (b *Bullet01) Draw(target *ebiten.Image) {
	b.csp.Draw(target)
}

func (b *Bullet01) Terminate() {
	b.csp.Terminate()
}
