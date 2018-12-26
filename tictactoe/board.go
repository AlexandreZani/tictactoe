package tictactoe

type boardSquare int8

const (
	X     = boardSquare(0)
	O     = boardSquare(1)
	EMPTY = boardSquare(2)
)

type board [9]boardSquare

type gameMove uint8

type player interface {
	Id() uint32
	Play(b board) gameMove
}

func (b board) Square(row, col int) boardSquare {
	if row < 0 || row > 2 {
		panic("Invalid row value")
	}
	if col < 0 || col > 2 {
		panic("Invalid column value")
	}
	return b[row*3+col]
}

type GameResult int8

const (
	X_WIN      = GameResult(0)
	O_WIN      = GameResult(1)
	DRAW       = GameResult(2)
	UNFINISHED = GameResult(4)
)

func (b board) row3(row0, col0, rowStep, colStep int) GameResult {
	s0 := b.Square(row0, col0)
	s1 := b.Square(row0+rowStep, col0+colStep)
	s2 := b.Square(row0+2*rowStep, col0+2*colStep)
	if s0 == s1 && s0 == s2 {
		if s0 == X {
			return X_WIN
		}
		if s0 == O {
			return O_WIN
		}
	}
	return UNFINISHED
}

func (b board) Evaluate() GameResult {
	for r := 0; r < 3; r++ {
		if res := b.row3(r, 0, 0, 1); res != UNFINISHED {
			return res
		}
	}

	for c := 0; c < 3; c++ {
		if res := b.row3(0, c, 1, 0); res != UNFINISHED {
			return res
		}
	}

	if res := b.row3(0, 0, 1, 1); res != UNFINISHED {
		return res
	}

	if res := b.row3(2, 0, -1, 1); res != UNFINISHED {
		return res
	}

	for _, s := range b {
		if s == EMPTY {
			return UNFINISHED
		}
	}

	return DRAW
}
