package game

import (
	"gRocketeer/core"
	"gRocketeer/core/config"
	"gRocketeer/core/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

var (
	ResWidth  = 320
	ResHeight = 180

	ViewWidth  = 960
	ViewHeight = 540

	G *Game
)

type Game struct {
	System *entity.System
	Con    *Controller
}

func (g *Game) Update() error {
	g.System.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.System.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ResWidth, ResHeight
}

func (g *Game) Sys() *entity.System {
	return g.System
}

func (g *Game) Screen() core.Vec {
	return core.NewVec(float64(ResWidth), float64(ResHeight))
}

func Run() {
	ebiten.SetWindowSize(ViewWidth, ViewHeight)
	ebiten.SetWindowTitle("Rocketeer")
	ebiten.SetTPS(90)

	config.Debug = true

	G = &Game{
		System: entity.NewSystem(),
	}

	G.Con = NewController(G)
	G.Con.ShowGame()

	if err := ebiten.RunGame(G); err != nil {
		log.Fatal(err)
	}
}
