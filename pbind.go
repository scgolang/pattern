package pattern

// Pbind generates events from a map of FloatGen's.
type Pbind struct {
	Instruments StringGen
	Controls    map[string]FloatGen
}

// Next gets the next event in the pattern.
func (pbind *Pbind) Next() (*Event, error) {
	inst, err := pbind.Instruments.Next()
	if err != nil {
		return nil, err
	}

	event, err := NewEvent(inst)
	if err != nil {
		return nil, err
	}

	for key, gen := range pbind.Controls {
		ctrl, err := gen.Next()
		if err != nil {
			return nil, err
		}
		event.AddCtrl(key, ctrl)
	}
	return event, nil
}
