package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/scgolang/pattern"
	"github.com/scgolang/sc"
)

var (
	DefSine = sc.NewSynthdef("testSine", func(p sc.Params) sc.Ugen {
		bus := sc.C(0)
		gate, freq := p.Add("gate", 1), p.Add("freq", 440)
		gain, release := p.Add("gain", 1), p.Add("release", 0.2)
		envgen := sc.EnvGen{
			Env:        sc.EnvPerc{Release: release},
			Gate:       gate,
			LevelScale: gain,
			Done:       sc.FreeEnclosing,
		}.Rate(sc.KR)

		sig := sc.SinOsc{Freq: freq}.Rate(sc.AR).Mul(envgen)

		// make it stereo
		sig = sc.Multi(sig, sig)

		return sc.Out{bus, sig}.Rate(sc.AR)
	})

	DefBlip = sc.NewSynthdef("testBlip", func(p sc.Params) sc.Ugen {
		bus := sc.C(0)
		gate, freq := p.Add("gate", 1), p.Add("freq", 440)
		gain, release := p.Add("gain", 1), p.Add("release", 0.2)
		envgen := sc.EnvGen{
			Env:        sc.EnvPerc{Release: release},
			Gate:       gate,
			LevelScale: gain,
			Done:       sc.FreeEnclosing,
		}.Rate(sc.KR)

		sig := sc.Blip{Freq: freq}.Rate(sc.AR).Mul(envgen)

		// make it stereo
		sig = sc.Multi(sig, sig)

		return sc.Out{bus, sig}.Rate(sc.AR)
	})

	DefSaw = sc.NewSynthdef("testSaw", func(p sc.Params) sc.Ugen {
		bus := sc.C(0)
		gate, freq := p.Add("gate", 1), p.Add("freq", 440)
		gain, release := p.Add("gain", 1), p.Add("release", 0.2)
		envgen := sc.EnvGen{
			Env:        sc.EnvPerc{Release: release},
			Gate:       gate,
			LevelScale: gain,
			Done:       sc.FreeEnclosing,
		}.Rate(sc.KR)

		sig := sc.Saw{Freq: freq}.Rate(sc.AR).Mul(envgen)

		// make it stereo
		sig = sc.Multi(sig, sig)

		return sc.Out{bus, sig}.Rate(sc.AR)
	})

	Controls = map[string]pattern.CtrlFunc{
		"freq": randomNotes(),
		"gain": func() float32 {
			return float32(0.5)
		},
		"release": pattern.Rand(pattern.Inf, []float32{0.1, 0.2, 0.5}),
		"timbre":  pattern.Rand(pattern.Inf, []float32{0, 1, 2}),
	}

	Defs = []string{"testSine", "testBlip", "testSaw"}

	Events = pattern.Pbind{
		Instruments: func() string {
			return Defs[rand.Intn(len(Defs))]
		},
		Controls: Controls,
	}

	durations = []time.Duration{
		62 * time.Millisecond,
		125 * time.Millisecond,
		187 * time.Millisecond,
		250 * time.Millisecond,
	}

	Durations = func() time.Duration {
		return durations[rand.Intn(len(durations))]
	}
)

func main() {
	// Initialize sc client and send the synthdef.
	client, err := sc.DefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	if err := client.SendDef(DefSine); err != nil {
		log.Fatal(err)
	}
	if err := client.SendDef(DefBlip); err != nil {
		log.Fatal(err)
	}
	if err := client.SendDef(DefSaw); err != nil {
		log.Fatal(err)
	}

	// Start playing the pattern.
	player, err := pattern.NewPlayer(Durations, Events)
	if err != nil {
		log.Fatal(err)
	}
	if err := player.Play(); err != nil {
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

func randomNotes() pattern.CtrlFunc {
	i, scale := 0, scales[rand.Intn(len(scales))]

	return func() float32 {
		if i%32 == 0 {
			scale = scales[rand.Intn(len(scales))]
		}
		i++
		return sc.Midicps(int(scale[rand.Intn(7)]) + 12*(rand.Intn(3)+3))
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
