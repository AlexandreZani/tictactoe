package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"tictactoe"
)

var n = flag.Int("n", 1, "number of games to run")
var summary = flag.Bool("summary", false, "display summary statistics")
var player1 = flag.String("player1", "perfect", "")
var player2 = flag.String("player2", "perfect", "")
var p1PlaybackFile = flag.String("p1_file", "", "")
var p2PlaybackFile = flag.String("p2_file", "", "")

func getPlayer(s string) tictactoe.Player {
	switch s {
	case "perfect":
		return tictactoe.NewPerfectPlayer()
	case "random":
		return tictactoe.NewRandomPlayer()
	}
	log.Fatalf("Unknown player: %s", s)
	return nil
}

func getPlaybackWriter(p string) tictactoe.PlaybackWriter {
	if p == "" {
		return tictactoe.DevNullPlaybackWriter{}
	}

	fp, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return tictactoe.NewCsvPlaybackWriter(csv.NewWriter(fp))
}

func main() {
	flag.Parse()

	p1 := getPlayer(*player1)
	p2 := getPlayer(*player2)
	w1 := getPlaybackWriter(*p1PlaybackFile)
	defer w1.Flush()
	w2 := getPlaybackWriter(*p2PlaybackFile)
	defer w2.Flush()

	draw := 0
	p1Win := 0
	p2Win := 0

	for i := 0; i < *n; i++ {
		xp := &p1
		op := &p2
		xWin := &p1Win
		oWin := &p2Win
		xw := &w1
		ow := &w2
		if rand.Int()%2 == 0 {
			xp = &p2
			op = &p1
			xWin = &p2Win
			oWin = &p1Win
			xw = &w2
			ow = &w1
		}

		g := tictactoe.NewGame(*xp, *op)
		switch g.Play() {
		case tictactoe.DRAW:
			draw++
		case tictactoe.X_WIN:
			*xWin++
		case tictactoe.O_WIN:
			*oWin++
		}
		g.AppendPlayback(tictactoe.XP, *xw)
		g.AppendPlayback(tictactoe.OP, *ow)

	}

	if *summary {
		fmt.Printf("%v\n%v\n%v\n", p1Win, draw, p2Win)
	}
}
