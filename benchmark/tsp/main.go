package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/seipan/goga"
	"github.com/seipan/goga/internal/selection"
)

type City struct {
	X, Y float64
}

func Distance(c1, c2 City) float64 {
	return math.Sqrt((c1.X-c2.X)*(c1.X-c2.X) + (c1.Y-c2.Y)*(c1.Y-c2.Y))
}

type Variables []City

func (v Variables) Initialization() goga.Genome {
	rand.Seed(time.Now().UnixNano())
	genome := make(Variables, len(v))
	copy(genome, v)
	rand.Shuffle(len(genome), func(i, j int) {
		genome[i], genome[j] = genome[j], genome[i]
	})
	return genome
}

func (v Variables) Fitness() float64 {
	totalDistance := 0.0
	for i := 0; i < len(v)-1; i++ {
		totalDistance += Distance(v[i], v[i+1])
	}
	totalDistance += Distance(v[len(v)-1], v[0]) // Return to the starting city
	return -totalDistance
}

func (v Variables) Mutation() {
	mutationRate := 0.01
	for i := range v {
		if rand.Float64() < mutationRate {
			j := rand.Intn(len(v))
			v[i], v[j] = v[j], v[i]
		}
	}
}

func (v Variables) Crossover(other goga.Genome) goga.Genome {
	partner := other.(Variables)
	child1, _ := OrderedCrossover(v, partner)
	return Variables(child1)
}

func main() {
	v := Variables{
		{X: 60, Y: 200}, {X: 180, Y: 200}, {X: 80, Y: 180}, {X: 140, Y: 180},
		{X: 20, Y: 160}, {X: 100, Y: 160}, {X: 200, Y: 160}, {X: 140, Y: 140},
		{X: 40, Y: 120}, {X: 100, Y: 120}, {X: 180, Y: 100}, {X: 60, Y: 80},
		{X: 120, Y: 80}, {X: 180, Y: 60}, {X: 20, Y: 40}, {X: 100, Y: 40},
		{X: 200, Y: 40}, {X: 20, Y: 20}, {X: 60, Y: 20}, {X: 160, Y: 20},
	}

	selector := selection.NewTournamentSelector(2)
	ga := goga.NewGA(goga.GAConfig{
		PopulationSize: 30,
		NGenerations:   200,
		CrossoverRate:  0.8,
		MutationRate:   0.01,
	}, selector)
	if err := ga.Minimize(v); err != nil {
		log.Fatal(err)
	}
}
