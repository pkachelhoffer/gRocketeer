//go:build ignore

//kage:unit pixels

package shaders

var Time float
var Cursor vec2

func Fragment(position vec4, srcPos vec2, color vec4) vec4 {
	xy := floor(position.x) + floor(position.y)
	if mod(xy, 2) == 0 {
		return vec4(1)
	}
	return vec4(0, 0, 0, 1)
}
