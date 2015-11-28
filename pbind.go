package pattern

import "time"

const (
	durKey     = "dur"
	defaultDur = float32(1.0)
)

type Durations func() time.Duration

type CtrlFunc func() float32

type InstFunc func() string

type Pbind struct {
	Instruments InstFunc
	Controls    map[string]CtrlFunc
}

func (pbind Pbind) Next() Event {
	e := Event{
		Instrument: pbind.Instruments(),
		Controls:   map[string]float32{},
	}
	for key, ctrlFunc := range pbind.Controls {
		e.Controls[key] = ctrlFunc()
	}
	return e
}
