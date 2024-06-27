package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/seipan/goga"
	"github.com/seipan/goga/internal/printer"
	"github.com/seipan/goga/internal/selection"
)

type Item struct {
	Weight, Value int
}

type Knapsack struct {
	Items    []Item
	Capacity int
}

type KnapsackGenome struct {
	Genes      []bool
	Knapsack   Knapsack
	FitnessVal float64
}

func (kg KnapsackGenome) Initialization() goga.Genome {
	rand.Seed(time.Now().UnixNano())
	genome := make([]bool, len(kg.Knapsack.Items))
	currentWeight := 0

	for i := range genome {
		if rand.Float64() < 0.5 {
			if currentWeight+kg.Knapsack.Items[i].Weight <= kg.Knapsack.Capacity {
				genome[i] = true
				currentWeight += kg.Knapsack.Items[i].Weight
			} else {
				genome[i] = false
			}
		} else {
			genome[i] = false
		}
	}
	return KnapsackGenome{Genes: genome, Knapsack: kg.Knapsack}
}

func (kg KnapsackGenome) Fitness() float64 {
	totalWeight, totalValue := 0, 0
	for i, included := range kg.Genes {
		if included {
			totalWeight += kg.Knapsack.Items[i].Weight
			totalValue += kg.Knapsack.Items[i].Value
		}
	}
	if totalWeight > kg.Knapsack.Capacity {
		kg.FitnessVal = 0
		kg.FitnessVal = float64(totalValue)
	}
	return kg.FitnessVal
}

func (kg KnapsackGenome) Mutation() {
	mutationRate := 0.01
	for i := range kg.Genes {
		if rand.Float64() < mutationRate {
			kg.Genes[i] = !kg.Genes[i]
		}
	}
}

func (kg KnapsackGenome) Crossover(other goga.Genome) goga.Genome {
	partner := other.(KnapsackGenome)
	child1Genes, _ := TwoPointCrossover(kg.Genes, partner.Genes)
	return KnapsackGenome{Genes: child1Genes, Knapsack: kg.Knapsack}
}

func main() {
	items := []Item{
		{Weight: 10, Value: 60},
		{Weight: 20, Value: 100},
		{Weight: 25, Value: 120},
		{Weight: 10, Value: 10},
		{Weight: 20, Value: 140},
		{Weight: 40, Value: 80},
		{Weight: 20, Value: 70},
		{Weight: 40, Value: 130},
		{Weight: 10, Value: 60},
		{Weight: 30, Value: 50},
		{Weight: 30, Value: 120},
	}

	capacity := 150
	knapsack := Knapsack{Items: items, Capacity: capacity}
	initialGenome := KnapsackGenome{Knapsack: knapsack}.Initialization()
	selector := selection.NewTournamentSelector(2)
	printer, err := printer.NewCSVPrinter()
	if err != nil {
		log.Fatal(err)
	}
	ga := goga.NewGA(goga.GAConfig{
		PopulationSize: 30,
		NGenerations:   100,
		CrossoverRate:  0.8,
		MutationRate:   0.01,
	}, selector, printer)
	if err := ga.Minimize(initialGenome); err != nil {
		log.Fatal(err)
	}
}
