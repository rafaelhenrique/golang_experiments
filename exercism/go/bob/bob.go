// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"regexp"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	expression := regexp.MustCompile("^[^a-z]+$")
	match := expression.MatchString(remark)

	fmt.Println()

	if match {
		return "Whoa, chill out!"
	} else if remark[len(remark)-1:] == "?" && match {
		return "Calm down, I know what I'm doing!"
	} else if remark[len(remark)-1:] == "?" {
		return "Sure."
	}
	return "Whatever."
}
