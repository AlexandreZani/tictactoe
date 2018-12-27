package tictactoe

type perfectPlayer struct{}

func (p perfectPlayer) Id() uint64 {
	return 1
}

func (r perfectPlayer) Play(b board, p playerR) gameMove {
	cur := 0
	for m, s := range b {
		if s != EMPTY {
			continue
		}

		bc := b
		bc.ApplyMove(gameMove(m), square(p))
		res := miniMax(bc, !p)
		if res == winning(p) {
			return gameMove(m)
		}
		if res == DRAW {
			cur = m
		}
	}
	return gameMove(cur)
}

// TODO(azani): Memoize
func miniMax(b board, p playerR) gameResult {
	if res := b.Evaluate(); res != UNFINISHED {
		return res
	}

	cur := winning(!p)

	for m, s := range b {
		if s == EMPTY {
			bc := b
			bc.ApplyMove(gameMove(m), square(p))
			res := miniMax(bc, !p)
			if res == DRAW {
				cur = DRAW
			} else if res == winning(p) {
				return res
			}
		}
	}

	return cur
}
