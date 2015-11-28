package main

import (
	"log"
	"time"

	"github.com/scgolang/pattern"
	"github.com/scgolang/sc"
)

var (
	Release = sc.C(0.5)

	Notes = []float32{
		sc.Midicps(48),
		sc.Midicps(50),
		sc.Midicps(52),
		sc.Midicps(53),
		sc.Midicps(55),
		sc.Midicps(60),
		sc.Midicps(62),
		sc.Midicps(64),
		sc.Midicps(67),
		sc.Midicps(69),
		sc.Midicps(71),
		sc.Midicps(72),
		sc.Midicps(79),
		sc.Midicps(81),
		sc.Midicps(83),
	}

	Def = sc.NewSynthdef("test", func(p sc.Params) sc.Ugen {
		bus := sc.C(0)
		gate, freq := p.Add("gate", 1), p.Add("freq", 440)
		gain := p.Add("gain", 1)
		envgen := sc.EnvGen{
			Env:        sc.EnvPerc{Release: Release},
			Gate:       gate,
			LevelScale: gain,
			Done:       sc.FreeEnclosing,
		}.Rate(sc.KR)
		sig := sc.Blip{Freq: freq}.Rate(sc.AR).Mul(envgen)
		sig = sc.Multi(sig, sig)
		return sc.Out{bus, sig}.Rate(sc.AR)
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
	controls := map[string]pattern.CtrlFunc{
		"freq": pattern.Rand(pattern.Inf, Notes),
		"gain": func() float32 {
			return float32(0.5)
		},
	}
	events := pattern.Pbind{
		Instruments: func() string { return "test" },
		Controls:    controls,
	}
	durations := func() time.Duration {
		return 125 * time.Millisecond
	}
	player, err := pattern.NewPlayer(durations, events)
	if err != nil {
		log.Fatal(err)
	}
	if err := player.Play(); err != nil {
		log.Fatal(err)
	}
}
