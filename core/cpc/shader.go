//go:build ignore

//kage:unit pixels

package cpc

var Center vec2
var Radius float
var Color vec4

func Fragment(position vec4, _ vec2, _ vec4) vec4 {
	distToCenter := distance(Center, position.xy)
	distToEdge := distToCenter - Radius

	// dist to edge will be negative if we are inside the
	// circle and positive if we are outside, but we want to
	// preserve the circle color if we are inside (multiply
	// by one), and discard it if we are outside (multiply
	// by zero), so we need to change the sign and clamp
	factor := clamp(-distToEdge, 0, 1)
	factor = pow(factor, 1.0/2.2) // gamma correction
	return vec4(Color.r, Color.g, Color.b, Color.a) * factor
}
