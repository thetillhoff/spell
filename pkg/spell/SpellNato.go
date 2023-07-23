package spell

// Returns a string for each character of the input string
// Spaces return an empty string
func SpellNato(input string) []string {
	var (
		spelling = []string{}

		spellMap = map[rune]string{
			'a': "alfa",
			'b': "bravo",
			'c': "charlie",
			'd': "delta",
			'e': "echo",
			'f': "foxtrott",
			'g': "golf",
			'h': "hotel",
			'i': "india",
			'j': "juliett",
			'k': "kilo",
			'l': "lima",
			'm': "mike",
			'n': "november",
			'o': "oscar",
			'p': "papa",
			'q': "quebec",
			'r': "romeo",
			's': "sierra",
			't': "tango",
			'u': "uniform",
			'v': "victor",
			'w': "whiskey",
			'x': "x-ray",
			'y': "yankee",
			'z': "zulu",
		}
	)

	for _, character := range input {

		if _, word := spellMap[character]; word { // If character is contained in list

			spelling = append(spelling, spellMap[character])

		} else { // If character is not contained in list

			switch character {
			case ' ': // If character is a space char
				spelling = append(spelling, "") // Add empty spelling
			}

		}
	}

	return spelling
}
