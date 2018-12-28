package tictactoe

type Game struct {
	px, po Player
	b      board
	result gameResult
	moves  []gameMove
}

func NewGame(px, po Player) Game {
	return Game{
		px:     px,
		po:     po,
		b:      newBoard(),
		result: UNFINISHED,
	}
}

func (g *Game) playTurn() {
	p := g.po
	pr := OP
	if len(g.moves)%2 == 0 {
		p = g.px
		pr = XP
	}
	m := p.Play(g.b, pr)
	g.moves = append(g.moves, m)

	// Illegal moves result in opponent victory.
	if !g.b.ApplyMove(m, square(pr)) {
		g.result = winning(!pr)
		return
	}

	g.result = g.b.Evaluate()
}

func (g *Game) Play() gameResult {
	for g.result == UNFINISHED {
		g.playTurn()
	}
	return g.result
}

func (g Game) AppendPlayback(p playerR, buf PlaybackWriter) {
	if g.result == UNFINISHED {
		panic("Can not generate playbacks from an unfinished game.")
	}

	playback := [19]bool{}

	// If the specified player did not lose, byte 18 is 1.
	playback[18] = true
	if winning(!p) == g.result {
		playback[18] = false
	}

	x := playback[0:9]
	o := playback[9:18]
	if p == OP {
		o = playback[0:9]
		x = playback[9:18]
	}

	cur := XP
	for _, m := range g.moves {
		if cur == XP {
			x[m] = true
		}
		if cur == OP {
			o[m] = true
		}
		if cur == p {
			buf.AddPlayback(playback)
		}
		cur = !cur
	}
}
