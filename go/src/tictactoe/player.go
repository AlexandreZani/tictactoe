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
	Play(b board, p playerR) gameMove
}

type gameMove uint8

type randomPlayer struct{}

func (p randomPlayer) Play(b board, _ playerR) gameMove {
	return gameMove(rand.Intn(8) % 9)
}

func NewRandomPlayer() randomPlayer {
	return randomPlayer{}
}

type randomValidPlayer struct{}

func (p randomValidPlayer) Play(b board, _ playerR) gameMove {
	empty := []int{}
	for i, s := range b {
		if s == EMPTY {
			empty = append(empty, i)
		}
	}
	return gameMove(empty[rand.Int()%len(empty)])
}

func NewRandomValidPlayer() randomValidPlayer {
	return randomValidPlayer{}
}

type nextAvailablePlayer struct{}

func (p nextAvailablePlayer) Play(b board, _ playerR) gameMove {
	for m, s := range b {
		if s == EMPTY {
			return gameMove(m)
		}
	}
	panic("This should never happen.")
}
