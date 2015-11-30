package pattern

import (
	"time"

	"github.com/scgolang/sc"
)

// Player defines a type that can play events.
type Player interface {
	Play(EventGen) error
}

// player can play Events using supercollider.
type player struct {
	client *sc.Client
	durs   DurGen
}

// Play plays a series of events.
func (p *player) Play(eg EventGen) error {
	var (
		event *Event
		err   error
	)
	for event, err = eg.Next(); err == nil; event, err = eg.Next() {
		dur, ed := p.durs.Next()
		if ed != nil {
			if ed == End {
				return nil
			}
			return ed
		}

		var (
			sid    = p.client.NextSynthID()
			gid    = sc.DefaultGroupID
			action = sc.AddToTail
			inst   = event.Instrument()
			ctrls  = event.Controls()
		)

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
func NewPlayer(durs DurGen) (Player, error) {
	client, err := sc.DefaultClient()
	if err != nil {
		return nil, err
	}
	return &player{client: client, durs: durs}, nil
}
