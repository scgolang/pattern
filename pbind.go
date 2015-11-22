package pattern

import "time"

const (
	durKey     = "dur"
	defaultDur = float32(1.0)
)

type Durations func() time.Duration

type Pattern func() float32

type Pbind struct {
	durations Durations
	patterns  map[string]Pattern
}

func (p *Pbind) Play() error {
	for dur := p.durations(); true; dur = p.durations() {
		// generate patterns
		time.Sleep(dur)
	}
	return nil
}
