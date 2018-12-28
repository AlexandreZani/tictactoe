package tictactoe

type playbackBuffer interface {
	AddPlayback(p [19]uint8)
}

type memPlaybackBuffer struct {
	buf [][19]uint8
}

func (m *memPlaybackBuffer) AddPlayback(p [19]uint8) {
	m.buf = append(m.buf, p)
}
