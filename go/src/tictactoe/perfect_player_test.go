package tictactoe

import "testing"

func TestMiniMax(t *testing.T) {
	b := newBoard()
	assertEq(t, DRAW, miniMax(b, X))
	assertEq(t, DRAW, miniMax(b, O))

	b[0] = X
	b[1] = X
	assertEq(t, X_WIN, miniMax(b, X))
}

func TestPerfectVsSelf(t *testing.T) {
	g := NewGame(perfectPlayer{}, perfectPlayer{})
	g.Play()
	assertEq(t, DRAW, g.result)
}
