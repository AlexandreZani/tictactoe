package tictactoe

import "testing"

func TestMiniMax(t *testing.T) {
	b := newBoard()
	assertEq(t, DRAW, miniMax(b, XP))
	assertEq(t, DRAW, miniMax(b, OP))

	b[0] = X
	b[1] = X
	assertEq(t, X_WIN, miniMax(b, XP))
}

func TestPerfectVsSelf(t *testing.T) {
	g := NewGame(perfectPlayer{}, perfectPlayer{})
	g.Play()
	assertEq(t, DRAW, g.result)
}
