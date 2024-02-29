package Algorithm

import (
	"errors"
	"math"
	"sort"
)

var GlobalVec [][]float64

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

// 定义一个结构体来存储索引和距离
type idxDist struct {
	idx      int
	distance float64
}

// 实现向量搜索函数
func FindNearestVectors(vectors [][]float64, targetIdx, n int) []int {
	distances := make([]idxDist, len(vectors))
	for i, vec := range vectors {
		distances[i] = idxDist{i, CosineSim(vectors[targetIdx], vec)}
	}

	// 根据距离对所有向量进行排序
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	// 选择最近的n个向量的索引
	nearestIdxs := make([]int, n)
	for i := 0; i < n; i++ {
		nearestIdxs[i] = distances[i+1].idx // 排除目标向量本身，所以是i+1
	}

	return nearestIdxs
}
