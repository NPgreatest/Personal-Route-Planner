package utils

import (
	"Personal-Route-Planner/Algorithm"
	"fmt"
)

func MakeRateMatrix() {

}

func GetSites() {

	prefs := Algorithm.MakeRatingMatrix([]float64{5, 5, 5, 0, 1,
		0, 0, 0, 4, 1,
		1, 2, 3, 3, 1,
		2, 0, 4, 1, 0,
		5, 2, 0, 1, 0,
		5, 2, 0, 1, 0}, 6, 5)
	products := []string{"Spiderman", "Big Momma's House", "Vanilla Sky", "Pacific Rim", "The Mask"}
	var rate []float64 = []float64{7, 2, 2, 4, 3}
	res := make(map[string]float64)
	for i := 0; i < len(rate); i++ {
		prods, scores, err := Algorithm.GetRecommendations(prefs, i, products)
		if err != nil {
			fmt.Println(err)
			return
		}
		for k := 0; k < len(prods); k++ {
			res[prods[k]] += scores[k] * rate[k]
		}
	}
	fmt.Println(res)
}
