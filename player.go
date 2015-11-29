package pattern

import (
	"errors"
	"time"

	"github.com/scgolang/sc"
)

var (
	ErrDuplicatePlayer = errors.New("duplicate player")
)

// Event defines a generic event.
type Event struct {
	Instrument string
	Controls   map[string]float32
}

// Nexter is implemented by anything that can generate events.
type Nexter interface {
	Next() Event
}

// Player can play Events using supercollider.
type Player struct {
	client    *sc.Client
	durations Durations
	nexter    Nexter
}

// Play plays a pattern.
func (p *Player) Play() error {
	for ev, dur := p.nexter.Next(), p.durations(); true; ev, dur = p.nexter.Next(), p.durations() {
		sid := p.client.NextSynthID()

		if _, err := p.client.Synth(ev.Instrument, sid, sc.AddToTail, sc.DefaultGroupID, ev.Controls); err != nil {
			return err
		}

		time.Sleep(dur)
	}
	return nil
}

// NewPlayer returns a new player.
func NewPlayer(durations Durations, nexter Nexter) (*Player, error) {
	scc, err := sc.DefaultClient()
	if err != nil {
		return nil, err
	}
	return &Player{
		client:    scc,
		durations: durations,
		nexter:    nexter,
	}, nil
}
