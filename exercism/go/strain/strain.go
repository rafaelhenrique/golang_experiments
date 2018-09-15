package strain

// Ints store array of int
type Ints []int

// Lists store array of arrays of int
type Lists [][]int

// Strings store array of strings
type Strings []string

// Keep walks numbers into Ints and keep numbers according predicate
func (ints Ints) Keep(predicate func(int) bool) (result Ints) {
	for _, number := range ints {
		if predicate(number) {
			result = append(result, number)
		}
	}
	return
}

// Discard walks numbers into Ints and discard numbers according predicate
func (ints Ints) Discard(predicate func(int) bool) (result Ints) {
	for _, number := range ints {
		if !predicate(number) {
			result = append(result, number)
		}
	}
	return
}

// Keep walks words into Strings and keep words according predicate
func (strings Strings) Keep(predicate func(string) bool) (result Strings) {
	for _, word := range strings {
		if predicate(word) {
			result = append(result, word)
		}
	}
	return
}

// Keep walks list into Lists and keep list according predicate
func (lists Lists) Keep(predicate func([]int) bool) (result Lists) {
	for _, list := range lists {
		if predicate(list) {
			result = append(result, list)
		}
	}
	return
}
