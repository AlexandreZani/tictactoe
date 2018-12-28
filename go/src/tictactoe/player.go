package tictactoe

import (
	"math/rand"
)

type playerR bool

const (
	XP = playerR(true)
	OP = playerR(false)
)

func winning(p playerR) gameResult {
	if p == XP {
		return X_WIN
	}

	return O_WIN
}

func square(p playerR) boardSquare {
	if p == XP {
		return X
	}

	return O
}

type Player interface {
	Id() uint64
	Play(b board, p playerR) gameMove
}

type gameMove uint8

type randomPlayer struct{}

func (p randomPlayer) Id() uint64 {
	return 0
}

func (p randomPlayer) Play(b board, _ playerR) gameMove {
	return gameMove(rand.Intn(8) % 9)
}

func NewRandomPlayer() randomPlayer {
	return randomPlayer{}
}

type nextAvailablePlayer struct{}

func (p nextAvailablePlayer) Id() uint64 {
	return 0
}

func (p nextAvailablePlayer) Play(b board, _ playerR) gameMove {
	for m, s := range b {
		if s == EMPTY {
			return gameMove(m)
		}
	}
	panic("This should never happen.")
}
