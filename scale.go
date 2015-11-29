package pattern

// Major intervals           2 2 1 2 2 2 1
var Major = [7]float32{0, 2, 4, 5, 7, 9, 11}

// NaturalMinor intervals    2 1 2 2 1 2 2
var NaturalMinor = [7]float32{0, 2, 3, 5, 7, 8, 10}

// HarmonicMinor intervals   2 1 2 2 1 3 1
var HarmonicMinor = [7]float32{0, 2, 3, 5, 7, 8, 11}

// Ionian intervals          2 1 2 2 2 2 1
var Ionian = [7]float32{0, 2, 3, 5, 7, 9, 11}

// Dorian intervals          2 1 2 2 2 1 2
var Dorian = [7]float32{0, 2, 3, 5, 7, 9, 10}

// Phrygian intervals        1 2 2 2 1 2 2
var Phrygian = [7]float32{0, 1, 3, 5, 7, 8, 10}

// Lydian intervals          2 2 2 1 2 2 1
var Lydian = [7]float32{0, 2, 4, 6, 7, 9, 11}

// Mixolydian intervals      2 2 1 2 2 1 2
var Mixolydian = [7]float32{0, 2, 4, 5, 7, 9, 10}

// Locrian intervals         1 2 2 1 2 2 2
var Locrian = [7]float32{0, 1, 3, 5, 6, 8, 10}
