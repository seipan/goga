package selection

import (
	"errors"
	"math/rand"
	"time"

	selection "github.com/seipan/goga"
)

type RouletteSelector struct{}

func (rs RouletteSelector) Select(population selection.Individuals) (selection.Individuals, error) {
	if len(population) == 0 {
		return nil, errors.New("集団が空です")
	}

	// 集団の総フィットネスを計算
	var totalFitness float64
	for _, ind := range population {
		totalFitness += ind.Fitness
	}

	if totalFitness == 0 {
		return nil, errors.New("集団のフィットネス合計が0です")
	}

	rand.Seed(time.Now().UnixNano())
	selected := make(selection.Individuals, 0, len(population))

	for i := 0; i < len(population); i++ {
		randomValue := rand.Float64()
		var cumulativeSum float64

		for _, ind := range population {
			cumulativeSum += ind.Fitness
			if randomValue < cumulativeSum {
				selected = append(selected, ind)
				break
			}
		}
	}

	return selected, nil
}
