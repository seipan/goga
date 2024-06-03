package mutation

import (
	"math/rand"
	"time"
)

func InversionMutation(individual []int) []int {
	n := len(individual)
	if n < 2 {
		return individual
	}

	rand.Seed(time.Now().UnixNano())
	start := rand.Intn(n)
	end := rand.Intn(n)

	if start > end {
		start, end = end, start
	}

	for i, j := start, end; i < j; i, j = i+1, j-1 {
		individual[i], individual[j] = individual[j], individual[i]
	}

	return individual
}
