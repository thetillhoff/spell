package spell

// Returns a string for each character of the input string
// Spaces return an empty string
func SpellGerman(input string) []string {
	var (
		spelling = []string{}

		spellMap = map[rune]string{
			'a': "anton",
			'b': "berta",
			'c': "cäsar",
			'd': "dora",
			'e': "emil",
			'f': "friedrich",
			'g': "gustav",
			'h': "heinrich",
			'i': "ida",
			'j': "julius",
			'k': "kaufmann",
			'l': "ludwig",
			'm': "martha",
			'n': "nordpol",
			'o': "otto",
			'p': "paula",
			'q': "quelle",
			'r': "richard",
			's': "siegfried",
			't': "theodor",
			'u': "ulrich",
			'v': "viktor",
			'w': "wilhelm",
			'x': "xanthippe",
			'y': "ypsilon",
			'z': "zeppelin",
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
