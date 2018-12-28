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

func TestAppendPlayback(t *testing.T) {
	g := NewGame(nextAvailablePlayer{}, nextAvailablePlayer{})
	g.Play()
	buf := memPlaybackBuffer{}
	g.AppendPlayback(XP, &buf)

	gold := [][19]bool{
		{true, false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false, false, true},
		{true, false, true, false, false, false, false, false, false,
			false, true, false, false, false, false, false, false, false, true},
		{true, false, true, false, true, false, false, false, false,
			false, true, false, true, false, false, false, false, false, true},
		{true, false, true, false, true, false, true, false, false,
			false, true, false, true, false, true, false, false, false, true},
	}
	assertEq(t, gold, buf.buf)
}
