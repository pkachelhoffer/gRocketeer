package game

import (
	"gRocketeer/core/config"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"math"
)

var (
	ResWidth  = 320
	ResHeight = 180

	ViewWidth  = 960
	ViewHeight = 540

	G *Game

	whiteImage = ebiten.NewImage(3, 3)
)

func init() {
	whiteImage.Fill(color.White)
}

type Game struct {
}

func (g Game) Update() error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	vertices := []ebiten.Vertex{
		{
			DstX:   0,
			DstY:   50,
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 0,
			ColorB: 0,
			ColorA: 1,
		},
		{
			DstX:   50,
			DstY:   0,
			SrcX:   0,
			SrcY:   0,
			ColorR: 0,
			ColorG: 1,
			ColorB: 0,
			ColorA: 1,
		},
		{
			DstX:   100,
			DstY:   50,
			SrcX:   0,
			SrcY:   0,
			ColorR: 0,
			ColorG: 0,
			ColorB: 1,
			ColorA: 1,
		},
	}
	// Define triangle indices
	indices := []uint16{0, 1, 2} // Indices of the three vertices

	// Draw triangle
	screen.DrawTriangles(vertices, indices, whiteImage, nil)

}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ResWidth, ResHeight
}

func Run() {
	ebiten.SetWindowSize(ViewWidth, ViewHeight)
	ebiten.SetWindowTitle("Rocketeer")
	ebiten.SetTPS(90)

	config.Debug = true

	G = &Game{}

	if err := ebiten.RunGame(G); err != nil {
		log.Fatal(err)
	}
}

func rotateVertex(x, y, angle float32) (float32, float32) {
	// Convert angle from degrees to radians
	radians := angle * (math.Pi / 180.0)

	// Compute sine and cosine of the angle
	sinTheta := float32(math.Sin(float64(radians)))
	cosTheta := float32(math.Cos(float64(radians)))

	// Apply rotation transformation
	newX := x*cosTheta - y*sinTheta
	newY := x*sinTheta + y*cosTheta

	return newX, newY
}
