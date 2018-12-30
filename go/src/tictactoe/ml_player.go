package tictactoe

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"log"
	"math/rand"
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

func (_ mlPlayer) Id() uint64 {
	return 2
}

func (p mlPlayer) Play(b board, r playerR) gameMove {
	buf := [9][27]float32{}
	for i := 0; i < 9; i++ {
		buf[i] = boardToSplitFloat(b, r, gameMove(i))
	}

	tensor, err := tf.NewTensor(buf)
	if err != nil {
		log.Fatal(err)
	}

	result, err := p.model.Session.Run(
		map[tf.Output]*tf.Tensor{p.input.Output(0): tensor},
		[]tf.Output{p.output.Output(0)},
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	output := result[0].Value().([][]float32)

	sum := float32(0.0)
	for _, a := range output {
		sum += a[0]
	}

	target := rand.Float32() * sum
	for i, a := range output {
		target -= a[0]
		if target < 0 {
			return gameMove(i)
		}
	}

	return gameMove(8)
}

func LoadModelOrDie(p string) *tf.SavedModel {
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

func boardToSplitFloat(b board, p playerR, m gameMove) (buf [27]float32) {
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

	return buf
}
