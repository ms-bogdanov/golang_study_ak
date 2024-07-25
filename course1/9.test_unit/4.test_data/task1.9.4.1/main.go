package main

func main() {

}

func average(xs []float64) float64 {

	var sum float64 = 0
	for _, v := range xs {
		sum += v
	}

	return sum / float64(len(xs))
}
