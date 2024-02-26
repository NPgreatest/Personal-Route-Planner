package Algorithm

import (
	"errors"
	"fmt"
	. "github.com/skelterjohn/go.matrix"
	"math"
	"sort"
	"strconv"
)

func TimesMatrix(a float64, b []float64) []float64 {
	for i := 0; i < len(b); i++ {
		b[i] += a * b[i]
	}
	return b
}

// 向量电积
func DotProduct(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return float64(0), errors.New("Cannot dot vectors of different length")
	}
	prod := float64(0)
	for i := 0; i < len(a); i++ {
		prod += a[i] * b[i]
	}
	return prod, nil
}

// 平方
func NormSquared(a []float64) float64 {
	sum := float64(0)
	for i := 0; i < len(a); i++ {
		sum += a[i] * a[i]
	}
	return math.Sqrt(sum)
}

// 余弦相似度
func CosineSim(a, b []float64) float64 {
	dp, err := DotProduct(a, b)
	if err != nil {
		return 0
	}
	a_squared := NormSquared(a)
	b_sqaured := NormSquared(b)
	return dp / (a_squared * b_sqaured)
}

// A并B/A交B
func Jaccard(a, b []float64) float64 {
	intersection := float64(0)
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			intersection += 1
		}
	}
	union := float64(0)
	for i := 0; i < len(a); i++ {
		if a[i] > 0 || b[i] > 0 {
			union += 1
		}
	}
	return intersection / union
}

// 更换NA
func replaceNA(prefs *DenseMatrix) *DenseMatrix {
	arr := prefs.Array()
	for i := 0; i < len(arr); i++ {
		if math.IsNaN(arr[i]) {
			arr[i] = float64(0)
		}
	}
	return MakeDenseMatrix(arr, prefs.Rows(), prefs.Cols())
}

// GetRecommendations Gets Recommendations for a user (row index) based on the prefs matrix.
// Uses cosine similarity for rating scale, and jaccard similarity if binary
func GetRecommendations(preference *DenseMatrix, user int, products []string) ([]string, []float64, error) {
	// make sure user is in the preference matrix
	if user >= preference.Rows() {
		return nil, nil, errors.New("user index out of range")
	}
	preference = replaceNA(preference)
	// item ratings
	ratings := make(map[int]float64, 0)
	sims := make(map[int]float64, 0)
	// Get user row from preference matrix
	userRatings := preference.GetRowVector(user).Array()
	for i := 0; i < preference.Rows(); i++ {
		// don't compare row to itself.
		if i != user {
			// get cosine similarity for other scores.
			other := preference.GetRowVector(i).Array()
			cosSim := CosineSim(userRatings, other)
			// get product recs for neighbors
			for idx, val := range other {
				if (userRatings[idx] == 0 || math.IsNaN(userRatings[idx])) && val != 0 {
					weightedRating := val * cosSim
					ratings[idx] += weightedRating
					sims[idx] += cosSim
				}
			}
		}
	}
	recs, values := calculateWeightedMean(ratings, sims, products)
	return recs, values, nil
}

func calculateWeightedMean(ratings, sims map[int]float64, products []string) (recommends []string, values []float64) {
	recommendations := make(map[float64]string, 0)
	for k, v := range ratings {
		meanProductRating := v / sims[k]
		fmt.Println(meanProductRating)
		if products != nil {
			recommendations[meanProductRating] = products[k]
		} else {
			recommendations[meanProductRating] = strconv.Itoa(k)
		}
	}
	recommends, values = sortMap(recommendations)
	return
}

// Sorts a map of floats -> strings to get the best recommendations. Probably a better way to do this.
func sortMap(recs map[float64]string) ([]string, []float64) {
	vals := make([]float64, 0)
	for k, _ := range recs {
		vals = append(vals, k)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))
	prods := make([]string, 0)
	for _, val := range vals {
		prods = append(prods, recs[val])
	}
	return prods, vals
}
func sum(x []float64) float64 {
	sum := float64(0)
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

func GetBinaryRecommendations(prefs *DenseMatrix, user int, products []string) ([]string, []float64, error) {
	// make sure user is in the preference matrix
	if user >= prefs.Rows() {
		return nil, nil, errors.New("user index out of range")
	}
	prefs = replaceNA(prefs)
	// item ratings
	ratings := make(map[float64]string)
	// Get user row from prefs matrix
	user_ratings := prefs.GetRowVector(user).Array()

	for ii := 0; ii < prefs.Cols(); ii++ {
		if user_ratings[ii] == float64(0) {
			num_liked := sum(prefs.GetColVector(ii).Array())
			num_disliked := float64(prefs.Rows()) - num_liked
			jaccard_liked := make([]float64, 0)
			jaccard_disliked := make([]float64, 0)
			for i := 0; i < prefs.Rows(); i++ {
				if i != user {
					other := prefs.GetRowVector(i).Array()
					if other[ii] == float64(0) {
						jaccard_disliked = append(jaccard_disliked, Jaccard(user_ratings, other))
					} else {
						jaccard_liked = append(jaccard_liked, Jaccard(user_ratings, other))
					}
				}
			}
			rating := (sum(jaccard_liked) - sum(jaccard_disliked)) / (num_disliked + num_liked)
			if products != nil {
				ratings[rating] = products[ii]
			} else {
				ratings[rating] = strconv.Itoa(ii)
			}
		}
	}
	prods, scores := sortMap(ratings)
	return prods, scores, nil
}

func MakeRatingMatrix(ratings []float64, rows, cols int) *DenseMatrix {
	return MakeDenseMatrix(ratings, rows, cols)
}
