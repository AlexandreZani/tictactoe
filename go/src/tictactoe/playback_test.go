package tictactoe

import (
	"bytes"
	"encoding/csv"
	"testing"
)

func TestCsvPlaybackWriter(t *testing.T) {
	gold := `0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,1
1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1
1,0,1,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1
1,0,1,0,1,0,0,0,0,0,1,0,1,0,1,0,0,0,0,0,0,0,0,0,1,0,0,1
`

	g := NewGame(nextAvailablePlayer{}, nextAvailablePlayer{})
	g.Play()

	sBuf := bytes.Buffer{}
	csvWriter := csv.NewWriter(&sBuf)
	w := NewCsvPlaybackWriter(csvWriter)
	g.AppendPlayback(XP, w)
	csvWriter.Flush()

	assertEq(t, gold, sBuf.String())
}
