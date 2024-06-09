package crossover

import (
	"math/rand"
	"time"
)

func TwoPointCrossover(parent1, parent2 []int) ([]int, []int) {
	if len(parent1) != len(parent2) {
		panic("両方の親の長さは同じでなければなりません")
	}

	length := len(parent1)
	child1 := make([]int, length)
	child2 := make([]int, length)

	rand.Seed(time.Now().UnixNano())
	crossPoint1 := rand.Intn(length)
	crossPoint2 := rand.Intn(length)

	if crossPoint1 > crossPoint2 {
		crossPoint1, crossPoint2 = crossPoint2, crossPoint1
	}

	for i := 0; i < length; i++ {
		if i < crossPoint1 || i >= crossPoint2 {
			child1[i] = parent1[i]
			child2[i] = parent2[i]
		} else {
			child1[i] = parent2[i]
			child2[i] = parent1[i]
		}
	}

	return child1, child2
}
