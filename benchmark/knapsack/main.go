package knapsack

import (
	"math/rand"
	"time"

	"github.com/seipan/goga"
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
	Knapsack   *Knapsack
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
		kg.FitnessVal = float64(totalValue) - float64(totalWeight-kg.Knapsack.Capacity)*10
	} else {
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
	childGenes := make([]bool, len(kg.Genes))
	for i := range kg.Genes {
		if rand.Float64() < 0.5 {
			childGenes[i] = kg.Genes[i]
		} else {
			childGenes[i] = partner.Genes[i]
		}
	}
	return KnapsackGenome{Genes: childGenes, Knapsack: kg.Knapsack}
}
