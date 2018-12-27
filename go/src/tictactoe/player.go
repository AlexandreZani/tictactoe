package tictactoe

import (
	"math/rand"
)

type player interface {
	Id() uint64
	Play(b board, p boardSquare) gameMove
}

type gameMove uint8

type randomPlayer struct{}

func (p randomPlayer) Id() uint64 {
	return 0
}

func (p randomPlayer) Play(b board, _ boardSquare) gameMove {
	return gameMove(rand.Intn(8) % 9)
}

func NewRandomPlayer() randomPlayer {
	return randomPlayer{}
}

type nextAvailablePlayer struct{}

func (p nextAvailablePlayer) Id() uint64 {
	return 0
}

func (p nextAvailablePlayer) Play(b board, _ boardSquare) gameMove {
	for m, s := range b {
		if s == EMPTY {
			return gameMove(m)
		}
	}
	panic("This should never happen.")
}
