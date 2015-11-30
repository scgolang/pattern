package pattern

import (
	"errors"
	"time"

	"github.com/scgolang/sc"
)

var (
	ErrDuplicatePlayer = errors.New("duplicate player")
)

// Player can play Events using supercollider.
type Player struct {
	client    *sc.Client
	durations DurGen
	events    EventGen
}

// Play plays a pattern.
func (p *Player) Play() error {
	var (
		event *Event
		err   error
	)
	for event, err = p.events.Next(); err == nil; event, err = p.events.Next() {
		dur, ed := p.durations.Next()
		if ed != nil {
			if ed == End {
				return nil
			}
			return ed
		}

		sid, gid, action := p.client.NextSynthID(), int32(sc.DefaultGroupID), int32(sc.AddToTail)
		inst, ctrls := event.Instrument(), event.Controls()

		if _, err := p.client.Synth(inst, sid, action, gid, ctrls); err != nil {
			return err
		}

		time.Sleep(dur)
	}
	if err != End {
		return err
	}
	return nil
}

// NewPlayer returns a new player.
func NewPlayer(durations DurGen, events EventGen) (*Player, error) {
	scc, err := sc.DefaultClient()
	if err != nil {
		return nil, err
	}
	return &Player{
		client:    scc,
		durations: durations,
		events:    events,
	}, nil
}
