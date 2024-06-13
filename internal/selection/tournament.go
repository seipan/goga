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
		winnerIndex := rand.Intn(len(indi))
		winner := indi[winnerIndex]
		for j := 1; j < t.TournamentSize; j++ {
			nextIndex := rand.Intn(len(indi))
			next := indi[nextIndex]
			if next.Fitness > winner.Fitness {
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
