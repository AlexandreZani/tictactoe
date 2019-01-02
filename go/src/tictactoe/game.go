package tictactoe

type Game struct {
	px, po      Player
	b           board
	result      gameResult
	illegalMove bool
	moves       []gameMove
}

func NewGame(px, po Player) Game {
	return Game{
		px:          px,
		po:          po,
		b:           newBoard(),
		result:      UNFINISHED,
		illegalMove: false,
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
		g.illegalMove = true
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

func (g Game) IllegalMove() bool {
	return g.illegalMove
}

func (g Game) AppendPlayback(p playerR, buf PlaybackWriter) {
	if g.result == UNFINISHED {
		panic("Can not generate playbacks from an unfinished game.")
	}

	playback := [28]bool{}

	// If the specified player did not lose, byte 27 is 1.
	playback[27] = true
	if winning(!p) == g.result {
		playback[27] = false
	}

	pb := playback[0:9]
	ob := playback[9:18]
	move := playback[18:27]

	cur := XP
	for _, m := range g.moves {
		if cur == p {
			move[m] = true
			buf.AddPlayback(playback)
			move[m] = false
			pb[m] = true
		} else {
			ob[m] = true
		}
		cur = !cur
	}
}
