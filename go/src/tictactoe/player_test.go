package tictactoe

import "testing"

func TestPlayerR(t *testing.T) {
	assertEq(t, OP, !XP)
	assertEq(t, XP, !OP)
	assertEq(t, OP, !!OP)
	assertEq(t, XP, !!XP)
}

func TestPlayerRWinning(t *testing.T) {
	assertEq(t, X_WIN, winning(XP))
	assertEq(t, O_WIN, winning(OP))
}
