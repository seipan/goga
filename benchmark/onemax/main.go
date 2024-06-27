package main

import (
	"log"
	"math/rand"
	"time"

	goga "github.com/seipan/goga"
	crossover "github.com/seipan/goga/internal/crossover"
	mutation "github.com/seipan/goga/internal/mutation"
	"github.com/seipan/goga/internal/printer"
	selection "github.com/seipan/goga/internal/selection"
)

type Variables []int

func (bg Variables) Initialization() goga.Genome {
	rand.Seed(time.Now().UnixNano())
	genome := make(Variables, len(bg))
	for i := range genome {
		if rand.Float64() < 0.5 {
			genome[i] = 1
		} else {
			genome[i] = 0
		}
	}
	return genome
}

func (bg Variables) Fitness() float64 {
	var fitness float64
	for _, gene := range bg {
		fitness += float64(gene)
	}
	return fitness
}

func (V Variables) Mutation() {
	mutation.Mutate(V, 0.01)
}

func (bg Variables) Crossover(other goga.Genome) goga.Genome {
	partner := other.(Variables)
	child1, _ := crossover.TwoPointCrossover(bg, partner)
	return Variables(child1)
}

func main() {
	genomeLength := 20 // Length of the binary genome
	var v Variables = make([]int, genomeLength)
	selector := selection.NewTournamentSelector(2)
	printer, err := printer.NewCSVPrinter()
	if err != nil {
		log.Fatal(err)
	}

	ga := goga.NewGA(goga.GAConfig{
		PopulationSize: 30,
		NGenerations:   40,
		CrossoverRate:  0.8,
		MutationRate:   0.01,
	}, selector, printer)
	if err := ga.Minimize(v); err != nil {
		log.Fatal(err)
	}

}
