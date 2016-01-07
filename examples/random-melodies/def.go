package main

import "github.com/scgolang/sc"

const (
	DefName = "pattern_examples_rand"
)

var Def = sc.NewSynthdef(DefName, func(p sc.Params) sc.Ugen {
	// Controls
	var (
		bus           = sc.C(0)
		gate, freq    = p.Add("gate", 1), p.Add("freq", 440)
		gain, release = p.Add("gain", 1), p.Add("release", 0.2)
		timbre        = p.Add("timbre", 0)
	)
	envgen := sc.EnvGen{
		Env:        sc.EnvPerc{Release: release},
		Gate:       gate,
		LevelScale: gain,
		Done:       sc.FreeEnclosing,
	}.Rate(sc.KR)

	// Oscillators
	var (
		sine = sc.SinOsc{Freq: freq}.Rate(sc.AR)
		blip = sc.Blip{Freq: freq}.Rate(sc.AR)
		saw  = sc.Saw{Freq: freq}.Rate(sc.AR)
	)
	sig := sc.Select{
		Which:  timbre,
		Inputs: []sc.Input{sine, blip, saw},
	}.Rate(sc.AR).Mul(envgen)

	// Make it stereo.
	sig = sc.Multi(sig, sig)

	return sc.Out{bus, sig}.Rate(sc.AR)
})
