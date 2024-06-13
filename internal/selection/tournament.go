package selection

import (
	"errors"
	"math/rand"
	"time"

	selection "github.com/seipan/goga"
)

type TournamentSelector struct {
	TournamentSize int
}

func (t TournamentSelector) Select(indi selection.Individuals) (selection.Individuals, error) {
	if len(indi) < t.TournamentSize {
		return nil, errors.New("invalid NSelection: Too large NSelection")
	}
	selected := make(selection.Individuals, len(indi))
	rand.Seed(time.Now().UnixNano())
	for i := range selected {
		winner := indi[rand.Intn(len(indi))]
		for j := 0; j < t.TournamentSize; j++ {
			next := indi[rand.Intn(len(indi))]
			if winner.Fitness < next.Fitness {
				winner = next
			}
		}
		selected[i] = winner
	}

	return selected, nil
}

func NewTournamentSelector(tournamentSize int) TournamentSelector {
	return TournamentSelector{TournamentSize: tournamentSize}
}
