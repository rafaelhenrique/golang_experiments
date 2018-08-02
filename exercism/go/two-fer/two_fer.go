// Twofer package has only one function responsible to create a sentence of the form "One for X, one for me."
package twofer

// ShareWith create a sentence of the form "One for X, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
