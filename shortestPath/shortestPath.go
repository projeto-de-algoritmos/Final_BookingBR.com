package shortestpath

import "math"

type Graph struct {
	to int
	wt float64
}

func floydWashall(g [][]Graph) [][]float64 {
	var results [][]float64 = make([][]float64, len(g))
	for i := range results {
		xs := make([]float64, len(g))
		for j := range xs {
			xs[j] = math.Inf(1)
		}
		xs[i] = 0
		results[i] = xs
	}

	for i, gv := range g {
		for _, value := range gv {
			results[i][value.to] = value.wt
		}
	}

	for k, dk := range results {
		for _, di := range results {
			for j, dij := range di {
				d := di[k] + dk[j]
				if dij > d {
					di[j] = d
				}
			}
		}

	}

	return results
}
