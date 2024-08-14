package main

func main() {

}

func sum(xs [8]int) int {
	sum := 0
	for _, v := range xs {
		sum += v
	}
	return sum
}

func average(xs [8]int) float64 {
	sum := 0
	for _, v := range xs {
		sum += v
	}
	return float64(sum) / float64(len(xs))
}

func averageFloat(xs [8]float64) float64 {
	sum := 0.0
	for _, v := range xs {
		sum += v
	}
	return sum / float64(len(xs))
}

func reverse(xs [8]int) [8]int {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
	return xs
}
