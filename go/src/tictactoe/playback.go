package tictactoe

import (
	"encoding/csv"
)

type PlaybackWriter interface {
	AddPlayback(p [19]bool)
	Flush()
}

type memPlaybackBuffer struct {
	buf [][19]bool
}

func (m *memPlaybackBuffer) AddPlayback(p [19]bool) {
	m.buf = append(m.buf, p)
}

func (m memPlaybackBuffer) Flush() {}

type csvPlaybackWriter struct {
	w *csv.Writer
}

func NewCsvPlaybackWriter(w *csv.Writer) *csvPlaybackWriter {
	return &csvPlaybackWriter{w: w}
}

func (w *csvPlaybackWriter) AddPlayback(p [19]bool) {
	r := [19]string{}
	for i := range p {
		if p[i] {
			r[i] = "1"
		} else {
			r[i] = "0"
		}
	}
	w.w.Write(r[:])
}

func (w *csvPlaybackWriter) Flush() {
	w.w.Flush()
}

type DevNullPlaybackWriter struct{}

func (w DevNullPlaybackWriter) AddPlayback(_ [19]bool) {}
func (w DevNullPlaybackWriter) Flush()                 {}
