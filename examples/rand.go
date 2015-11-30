package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/scgolang/pattern"
	"github.com/scgolang/sc"
)

const (
	DefName   = "pattern_examples_rand"
	OctaveMin = 3
	OctaveMax = 5
)

var (
	Def = sc.NewSynthdef(DefName, func(p sc.Params) sc.Ugen {
		bus := sc.C(0)
		gate, freq := p.Add("gate", 1), p.Add("freq", 440)
		gain, release := p.Add("gain", 1), p.Add("release", 0.2)
		timbre := p.Add("timbre", 0)

		envgen := sc.EnvGen{
			Env:        sc.EnvPerc{Release: release},
			Gate:       gate,
			LevelScale: gain,
			Done:       sc.FreeEnclosing,
		}.Rate(sc.KR)

		sine := sc.SinOsc{Freq: freq}.Rate(sc.AR).Mul(envgen)
		blip := sc.Blip{Freq: freq}.Rate(sc.AR).Mul(envgen)
		saw := sc.Saw{Freq: freq}.Rate(sc.AR).Mul(envgen)

		sig := sc.Select{
			Which:  timbre,
			Inputs: []sc.Input{sine, blip, saw},
		}.Rate(sc.AR)

		// make it stereo
		sig = sc.Multi(sig, sig)

		return sc.Out{bus, sig}.Rate(sc.AR)
	})

	Controls = map[string]pattern.FloatGen{
		"freq":    &RandomNotes{},
		"gain":    pattern.F(0.5),
		"release": &pattern.Frand{Values: []float32{0.1, 0.2, 0.5}},
		"timbre":  &pattern.Frand{Values: []float32{0, 1, 2}},
	}

	Events = &pattern.Pbind{
		Instruments: pattern.S(DefName),
		Controls:    Controls,
	}

	Durations = RandomDur([]time.Duration{
		128 * time.Millisecond,
		256 * time.Millisecond,
	})
)

func main() {
	// Initialize sc client and send the synthdef.
	client, err := sc.DefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	if err := client.SendDef(Def); err != nil {
		log.Fatal(err)
	}

	// Start playing the pattern.
	player, err := pattern.NewPlayer(Durations)
	if err != nil {
		log.Fatal(err)
	}
	if err := player.Play(Events); err != nil {
		log.Fatal(err)
	}
}

var scales = [][7]float32{
	pattern.Major,
	pattern.NaturalMinor,
	pattern.Dorian,
	pattern.Phrygian,
	pattern.Ionian,
	pattern.Mixolydian,
}

type RandomNotes struct {
	idx int
}

func (pat *RandomNotes) Next() (float32, error) {
	scale := scales[rand.Intn(len(scales))]
	if pat.idx%128 == 0 {
		scale = scales[rand.Intn(len(scales))]
		pat.idx = 0
	}
	pat.idx++
	return sc.Midicps(int(scale[rand.Intn(7)]) + 12*(rand.Intn(OctaveMax)+OctaveMin)), nil
}

type RandomDur []time.Duration

func (rd RandomDur) Next() (time.Duration, error) {
	return rd[rand.Intn(len(rd))], nil
}

func init() {
	rand.Seed(time.Now().Unix())
}
