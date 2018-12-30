package tictactoe

import (
	"testing"
)

func TestLoadModel(t *testing.T) {
	m := LoadModelOrDie("testdata/test_model")
	assertNeq(t, nil, m)
}

func TestTestModelVsSelf(t *testing.T) {
	m := LoadModelOrDie("testdata/test_model")
	g := NewGame(NewMlPlayer(m), NewMlPlayer(m))
	g.Play()
}

func TestBoardToSplitFloat(t *testing.T) {
	E := EMPTY
	b := board{
		X, O, O,
		E, X, E,
		E, E, X,
	}

	gold := [27]float32{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,

		0.0, 1.0, 1.0,
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,

		0.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 0.0,
	}

	actual := boardToSplitFloat(b, XP, gameMove(4))

	assertEq(t, gold, actual)
}
