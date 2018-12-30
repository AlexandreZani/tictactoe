package tictactoe

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"log"
	"strings"
)

type mlPlayer struct {
	model  *tf.SavedModel
	input  *tf.Operation
	output *tf.Operation
}

func NewMlPlayer(m *tf.SavedModel) mlPlayer {
	i := findLayerWithNamePart(m, "inputLayer_input")
	o := findLayerWithNamePart(m, "outputLayer/Sigmoid")

	return mlPlayer{model: m, input: i, output: o}
}

func loadModelOrDie(p string) *tf.SavedModel {
	m, err := tf.LoadSavedModel(p, []string{"tictactoe"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func findLayerWithNamePart(m *tf.SavedModel, part string) *tf.Operation {
	for i, o := range m.Graph.Operations() {
		if strings.Contains(o.Name(), part) {
			return &m.Graph.Operations()[i]
		}
	}
	log.Fatalf("Could not find %s.", part)
	return nil
}

func boardToSplitFloat(b board, p playerR, m gameMove, buf *[27]float32) {
	ps := square(p)
	opp := square(!p)
	buf[int(m)+18] = 1.0

	for i, s := range b {
		if s == ps {
			buf[i] = 1.0
		}
		if s == opp {
			buf[i+9] = 1.0
		}
	}
}
