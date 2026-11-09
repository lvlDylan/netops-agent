package collector

func roundFloat(val float64, precision uint) float64 {
	ratio := 1.0
	for i := uint(0); i < precision; i++ {
		ratio *= 10
	}
	return float64(int(val*ratio+0.5)) / ratio
}
