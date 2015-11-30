package pattern

// Fdrain drains a FloatGen and returns the resulting array.
func Fdrain(gen FloatGen) ([]float32, error) {
	var (
		vals = []float32{}
		val  float32
		err  error
	)
	for val, err = gen.Next(); err == nil; val, err = gen.Next() {
		vals = append(vals, val)
	}
	if err == End {
		return vals, nil
	}
	return nil, err
}

// Sdrain drains a StringGen and returns the resulting array.
func Sdrain(gen StringGen) ([]string, error) {
	var (
		vals = []string{}
		val  string
		err  error
	)
	for val, err = gen.Next(); err == nil; val, err = gen.Next() {
		vals = append(vals, val)
	}
	if err == End {
		return vals, nil
	}
	return nil, err
}
