package tictactoe

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"log"
	"math/rand"
)

type mlPlayer struct {
	model  *tf.SavedModel
	input  *tf.Operation
	output *tf.Operation
}

func NewMlPlayer(m *tf.SavedModel) mlPlayer {
	i := findLayerWithName(m, "inputLayer_input")
	o := findLayerWithName(m, "outputLayer/Sigmoid")

	return mlPlayer{model: m, input: i, output: o}
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

	return gameMove(sampleWeighted(output, rand.Float32()))
}

func variance(w [][]float32) float32 {
	mean := float32(0.0)
	for _, a := range w {
		mean += (a[0] / float32(len(w)))
	}

	v := float32(0.0)
	for _, a := range w {
		v += (a[0] - mean) * (a[0] - mean) / float32(len(w))
	}

	return v
}

func sampleWeighted(w [][]float32, v float32) int {
	sum := float32(0.0)
	for _, a := range w {
		sum += a[0]
	}

	t := v * sum
	for i, a := range w {
		t -= a[0]
		if t <= 0 {
			return i
		}
	}

	return len(w) - 1
}

func argMax(w [][]float32) int {
	r := 0
	for i := range w {
		if w[i][0] > w[r][0] {
			r = i
		}
	}

	return r
}

func LoadModelOrDie(p string) *tf.SavedModel {
	m, err := tf.LoadSavedModel(p, []string{"tictactoe"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func findLayerWithName(m *tf.SavedModel, name string) (op *tf.Operation) {
	op = nil
	for i, o := range m.Graph.Operations() {
		if o.Name() == name {
			if op != nil {
				log.Fatalf("Found more than one of %s.", name)
			}
			op = &m.Graph.Operations()[i]
		}
	}
	if op == nil {
		log.Fatalf("Could not find %s.", name)
	}
	return op
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
