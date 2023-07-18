package examples

func SimpleInterest(p, t, r int) float64 {
	return (float64(p*t*r) / 100) // typecasting p*t*r (int) -> float64
}
