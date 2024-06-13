package crossover

import (
	"math/rand"
	"time"
)

func TwoPointCrossover(parent1, parent2 []int) ([]int, []int) {
	length := len(parent1)
	if length != len(parent2) {
		panic("Parents must be of the same length")
	}

	rand.Seed(time.Now().UnixNano())

	point1 := rand.Intn(length)
	point2 := rand.Intn(length)
	for point1 == point2 {
		point2 = rand.Intn(length)
	}

	if point1 > point2 {
		point1, point2 = point2, point1
	}

	child1 := make([]int, length)
	child2 := make([]int, length)

	copy(child1, parent1)
	copy(child2, parent2)

	for i := point1; i <= point2; i++ {
		child1[i], child2[i] = child2[i], child1[i]
	}

	return child1, child2
}
