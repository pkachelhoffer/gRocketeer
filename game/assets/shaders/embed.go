package shaders

import (
	_ "embed"
)

var (
	//go:embed testshader.go
	ShaderTest []byte

	//go:embed circle.go
	ShaderCircle []byte
)
