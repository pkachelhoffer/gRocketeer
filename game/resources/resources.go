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

	ImgShip01     *ebiten.Image
	ImgShipHull01 *ebiten.Image
)

func init() {
	ShaderTest = core.LoadShaderBytes(shaders.ShaderTest)
	ShaderCircle = core.LoadShaderBytes(shaders.ShaderCircle)

	ImgShip01 = core.LoadImageBytes(art.Ship01)
	ImgShipHull01 = core.LoadImageBytes(art.ShipHull01)
}
