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

func TestSampleWeighted(t *testing.T) {
	w := [][]float32{
		[]float32{0.5},
		[]float32{0.3},
		[]float32{0.7},
		[]float32{0.2},
	}

	sum := 0.5 + 0.3 + 0.7 + 0.2

	assertEq(t, 0, sampleWeighted(w, float32(0.1/sum)))
	assertEq(t, 1, sampleWeighted(w, float32(0.6/sum)))
	assertEq(t, 2, sampleWeighted(w, float32(0.9/sum)))
	assertEq(t, 3, sampleWeighted(w, float32(1.6/sum)))
}

func TestArgMax(t *testing.T) {
	w := [][]float32{
		[]float32{0.5},
		[]float32{0.3},
		[]float32{0.7},
		[]float32{0.2},
	}

	assertEq(t, 2, argMax(w))
}
