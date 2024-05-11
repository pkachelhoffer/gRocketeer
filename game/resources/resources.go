package resources

import (
	"gRocketeer/core"
	"gRocketeer/game/assets/art"
	"gRocketeer/game/assets/shaders"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ShaderTest   *ebiten.Shader
	ShaderCircle *ebiten.Shader

	ImgShip       *ebiten.Image
	ImgShipHull01 *ebiten.Image
)

func init() {
	ShaderTest = core.LoadShaderBytes(shaders.ShaderTest)
	ShaderCircle = core.LoadShaderBytes(shaders.ShaderCircle)

	ImgShip = core.LoadImageBytes(art.Ship)
	ImgShipHull01 = core.LoadImageBytes(art.ShipHull01)
}
