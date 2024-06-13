package mutation

import (
	"math/rand"
	"time"
)

func Mutate(genome []int, rate float64) []int {
	rand.Seed(time.Now().UnixNano())
	mutatedGenome := make([]int, len(genome))
	copy(mutatedGenome, genome)
	for i := range mutatedGenome {
		if rand.Float64() < rate {
			if mutatedGenome[i] == 0 {
				mutatedGenome[i] = 1
			} else {
				mutatedGenome[i] = 0
			}
		}
	}
	return mutatedGenome
}
