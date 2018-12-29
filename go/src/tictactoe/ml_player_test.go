package tictactoe

import (
	//tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"testing"
)

func TestLoadModel(t *testing.T) {
	m := loadModelOrDie("testdata/test_model")
	assertNeq(t, nil, m)

	/*
		p := NewMlPlayer(m)

		tensor, _ := tf.NewTensor([2][18]float32{})
		result, err := p.model.Session.Run(
			map[tf.Output]*tf.Tensor{p.input.Output(0): tensor},
			[]tf.Output{p.output.Output(0)},
			nil,
		)

		if err != nil {
			t.Error(err)
		}

		t.Error(result[0].Value())
	*/
}

func TestBoardToSplitFloat(t *testing.T) {
	E := EMPTY
	b := board{
		X, O, O,
		E, X, E,
		E, E, X,
	}

	gold := [18]float32{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,

		0.0, 1.0, 1.0,
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,
	}

	actual := [18]float32{}

	boardToSplitFloat(b, XP, &actual)

	assertEq(t, gold, actual)
}
