package core

func Lerp(from float64, to float64, val float64) float64 {
	rng := to - from
	return from + (rng * val)
}

func LerpInt(from int, to int, val float64) int {
	rng := to - from
	return from + int(float64(rng)*val)
}
