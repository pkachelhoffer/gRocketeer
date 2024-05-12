package game

import (
	"gRocketeer/core"
	"gRocketeer/core/config"
	"gRocketeer/core/entity"
	"gRocketeer/game/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
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
	Space  *cp.Space
	Con    *Controller
}

func (g *Game) GetSpace() *cp.Space {
	return g.Space
}

func (g *Game) Update() error {
	g.Space.Step(1.0 / float64(ebiten.TPS()))
	g.System.Update()
	g.Con.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(types.ColorSpace)
	g.System.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ResWidth, ResHeight
}

func (g *Game) GetSys() *entity.System {
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
		Space:  cp.NewSpace(),
	}

	G.Space.SetGravity(cp.Vector{X: 0, Y: 0})
	G.Space.SleepTimeThreshold = 0.5
	G.Space.SetDamping(0.6)

	G.Con = NewController(G)
	G.Con.LoadLevel()

	if err := ebiten.RunGame(G); err != nil {
		log.Fatal(err)
	}
}
