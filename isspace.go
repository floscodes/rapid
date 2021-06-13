package main

import "unicode"

func Checkspace(s string) bool {
	space := true

	for _, x := range s {
		if !unicode.IsSpace(x) {
			space = false
			break
		}
	}

	return space
}
