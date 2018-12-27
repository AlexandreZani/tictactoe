package tictactoe

type Game struct {
	px, po player
	b      board
	result gameResult
	moves  []gameMove
}

func NewGame(px, po player) Game {
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

func (g *Game) Play() {
	for g.result == UNFINISHED {
		g.playTurn()
	}
}
