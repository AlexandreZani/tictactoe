package tictactoe

import (
	"testing"
)

func TestPlay(t *testing.T) {
	g := NewGame(nextAvailablePlayer{}, nextAvailablePlayer{})
	g.Play()
	assertEq(t, X_WIN, g.result)
	assertEq(t, 7, len(g.moves))
}
