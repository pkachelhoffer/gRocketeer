package art

import (
	_ "embed"
)

var (
	//go:embed ship.png
	Ship []byte

	//go:embed ShipHull01.png
	ShipHull01 []byte

	//go:embed Ship01.png
	Ship01 []byte
)
