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

	gold := [][19]uint8{
		{1, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 0, 0, 0, 0,
			0, 1, 0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0, 0,
			0, 1, 0, 1, 0, 1, 0, 0, 0, 1},
	}
	assertEq(t, gold, buf.buf)
}
