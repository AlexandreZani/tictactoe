package tictactoe

import (
	"reflect"
	"testing"
)

func assertEq(t *testing.T, e interface{}, r interface{}) {
	if e != r {
		t.Errorf("Expected %v (type %v)\nReceived %v (type %v)",
			e, reflect.TypeOf(e), r, reflect.TypeOf(r))
	}
}

func assertNeq(t *testing.T, e interface{}, r interface{}) {
	if e == r {
		t.Errorf("Expected not %v (type %v)\nReceived %v (type %v)",
			e, reflect.TypeOf(e), r, reflect.TypeOf(r))
	}
}

func assertTrue(t *testing.T, v bool) {
	assertEq(t, true, v)
}

func assertFalse(t *testing.T, v bool) {
	assertEq(t, false, v)
}
