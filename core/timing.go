package core

func SecondsLen(s float64) int {
	return int(s * 60)
}

func PerSecond(s float64) float64 {
	return s / 60
}
