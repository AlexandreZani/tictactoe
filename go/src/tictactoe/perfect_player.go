package tictactoe

type perfectPlayer struct{}

func (p perfectPlayer) Id() uint64 {
	return 1
}

func (r perfectPlayer) Play(b board, p boardSquare) gameMove {
	best := X_WIN
	next := O
	if p == O {
		best = O_WIN
		next = X
	}

	cur := 0
	for m, s := range b {
		if s != EMPTY {
			continue
		}

		bc := b
		bc[m] = p
		res := miniMax(bc, next)
		if res == best {
			return gameMove(m)
		}
		if res == DRAW {
			cur = m
		}
	}
	return gameMove(cur)
}

// TODO(azani): Memoize
func miniMax(b board, p boardSquare) gameResult {
	if res := b.Evaluate(); res != UNFINISHED {
		return res
	}

	best := X_WIN
	cur := O_WIN
	next := O
	if p == O {
		best = O_WIN
		cur = X_WIN
		next = X
	}

	for m, s := range b {
		if s == EMPTY {
			bc := b
			bc[m] = p
			res := miniMax(bc, next)
			if res == DRAW {
				cur = DRAW
			} else if res == best {
				return res
			}
		}
	}

	return cur
}
