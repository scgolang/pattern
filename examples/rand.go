package main

import (
	"log"

	"github.com/scgolang/pattern"
)

func main() {
	durations := func() {
		return float32(0.5)
	}
	pat := pattern.Pbind{
		durations: durations,
		patterns: map[string]Pattern{
			"degree": pattern.Rand(pattern.Inf, []float32{
				float32(60),
				float32(62),
				float32(64),
				float32(67),
			}),
		},
	}
	if err := pat.Play(); err != nil {
		log.Fatal(err)
	}
}
