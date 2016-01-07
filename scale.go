package pattern

const ScaleLen = 7

// Major intervals           2 2 1 2 2 2 1
var Major = [ScaleLen]float32{0, 2, 4, 5, 7, 9, 11}

// NaturalMinor intervals    2 1 2 2 1 2 2
var NaturalMinor = [ScaleLen]float32{0, 2, 3, 5, 7, 8, 10}

// HarmonicMinor intervals   2 1 2 2 1 3 1
var HarmonicMinor = [ScaleLen]float32{0, 2, 3, 5, 7, 8, 11}

// Ionian intervals          2 1 2 2 2 2 1
var Ionian = [ScaleLen]float32{0, 2, 3, 5, 7, 9, 11}

// Dorian intervals          2 1 2 2 2 1 2
var Dorian = [ScaleLen]float32{0, 2, 3, 5, 7, 9, 10}

// Phrygian intervals        1 2 2 2 1 2 2
var Phrygian = [ScaleLen]float32{0, 1, 3, 5, 7, 8, 10}

// Lydian intervals          2 2 2 1 2 2 1
var Lydian = [ScaleLen]float32{0, 2, 4, 6, 7, 9, 11}

// Mixolydian intervals      2 2 1 2 2 1 2
var Mixolydian = [ScaleLen]float32{0, 2, 4, 5, 7, 9, 10}

// Locrian intervals         1 2 2 1 2 2 2
var Locrian = [ScaleLen]float32{0, 1, 3, 5, 6, 8, 10}
