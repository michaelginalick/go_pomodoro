package digits


// GetDigits : returns the ansii representation of the time
func GetDigits(d string) [5]string {

	var bigChars = map[string][5]string{
		"p": {
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
		"0": {
			"XXX",
			"X X",
			"X X",
			"X X",
			"XXX",
		},
		"1": {
			" X ",
			"XX ",
			" X ",
			" X ",
			"XXX",
		},
		"2": {
			"XXX",
			"  X",
			"XXX",
			"X  ",
			"XXX",
		},
		"3": {
			"XXX",
			"  X",
			"XXX",
			"  X",
			"XXX",
		},
		"4": {
			"X X",
			"X X",
			"XXX",
			"  X",
			"  X",
		},
		"5": {
			"XXX",
			"X  ",
			"XXX",
			"  X",
			"XXX",
		},
		"6": {
			"XXX",
			"X  ",
			"XXX",
			"X X",
			"XXX",
		},
		"7": {
			"XXX",
			"  X",
			"  X",
			" X ",
			" X ",
		},
		"8": {
			"XXX",
			"X X",
			"XXX",
			"X X",
			"XXX",
		},
		"9": {
			"XXX",
			"X X",
			"XXX",
			"  X",
			"  X",
		},
		":": {
			"   ",
			" X ",
			"   ",
			" X ",
			"   ",
		},
		".": {
			"   ",
			"   ",
			"   ",
			" XX",
			" XX",
		},
	}

	return bigChars[d]
}

