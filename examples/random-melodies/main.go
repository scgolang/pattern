package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/scgolang/pattern"
	"github.com/scgolang/sc"
)

const (
	octaveMin = 3
	octaveMax = 5
)

var (
	controls = map[string]pattern.FloatGen{
		"freq":    &RandomNotes{},
		"gain":    pattern.F(0.5),
		"release": &pattern.Frand{Values: []float32{0.1, 0.2, 0.5}},
		"timbre":  &pattern.Frand{Values: []float32{0, 1, 2}},
	}

	events = &pattern.Pbind{
		Instruments: pattern.S(defName),
		Controls:    controls,
	}

	durations = RandomDur([]time.Duration{
		64 * time.Millisecond,
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
	if err := client.SendDef(def); err != nil {
		log.Fatal(err)
	}

	// Start playing the pattern.
	player, err := pattern.NewPlayer(durations)
	if err != nil {
		log.Fatal(err)
	}
	if err := player.Play(events); err != nil {
		log.Fatal(err)
	}
}

var (
	scales = [][pattern.ScaleLen]float32{
		pattern.Major,
		// pattern.NaturalMinor,
		// pattern.Dorian,
		// pattern.Phrygian,
		// pattern.Ionian,
		// pattern.Mixolydian,
	}
	numScales = len(scales)
)

type randomNotes struct {
	idx int
}

func (pat *randomNotes) Next() (float32, error) {
	scale := scales[rand.Intn(len(scales))]
	if pat.idx%128 == 0 {
		scale = scales[rand.Intn(numScales)]
		pat.idx = 0
	}
	pat.idx++
	return sc.Midicps(int(scale[rand.Intn(pattern.ScaleLen)]) + 12*(rand.Intn(octaveMax)+octaveMin)), nil
}

type randomDur []time.Duration

func (rd randomDur) Next() (time.Duration, error) {
	return rd[rand.Intn(len(rd))], nil
}

func init() {
	rand.Seed(time.Now().Unix())
}
