package tictactoe

import (
	"reflect"
	"testing"
)

func assertEq(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v)\nReceived %v (type %v)",
			a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

func TestSquare(t *testing.T) {
	b := board{0, 1, 2, 3, 4, 5, 6, 7, 8}
	assertEq(t, boardSquare(0), b.Square(0, 0))
	assertEq(t, boardSquare(1), b.Square(0, 1))
	assertEq(t, boardSquare(2), b.Square(0, 2))

	assertEq(t, boardSquare(3), b.Square(1, 0))
	assertEq(t, boardSquare(4), b.Square(1, 1))
	assertEq(t, boardSquare(5), b.Square(1, 2))

	assertEq(t, boardSquare(6), b.Square(2, 0))
	assertEq(t, boardSquare(7), b.Square(2, 1))
	assertEq(t, boardSquare(8), b.Square(2, 2))
}

func TestEvaluate(t *testing.T) {
	E := EMPTY
	b := board{
		O, O, O,
		E, E, E,
		E, E, E,
	}
	assertEq(t, O_WIN, b.Evaluate())

	b = board{
		X, O, O,
		E, X, E,
		E, E, X,
	}
	assertEq(t, X_WIN, b.Evaluate())

	b = board{
		X, O, O,
		O, X, X,
		O, X, O,
	}
	assertEq(t, DRAW, b.Evaluate())

	b = board{
		X, O, O,
		O, X, X,
		O, X, E,
	}
	assertEq(t, UNFINISHED, b.Evaluate())
}
