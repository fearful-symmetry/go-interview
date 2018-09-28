package arrays

import (
	"fmt"
)

//QUESTION: Implement a compression algorithim that replaces repeated characters in a word with number counts. Example: xxxcaaa becames x3ca3. If the length remains unchanged, return the original string

//CompressString compresses a string, or returns the original string if its length can't be reduced.
func CompressString(str string) string {

	var repeats int

	//This will be easier to manipulate as a rune slice
	strSlice := []rune(str)
	var comp string
	//iterate over runes
	for iter, c := range strSlice {

		//short out if we're on the last element
		if iter+1 == len(strSlice) {
			if repeats > 0 {
				comp = fmt.Sprintf("%s%c%d", comp, c, repeats+1)
			} else {
				comp = fmt.Sprintf("%s%c", comp, c)
			}
			continue
		}

		if strSlice[iter+1] != c {
			if repeats > 0 {
				comp = fmt.Sprintf("%s%c%d", comp, c, repeats+1)
			} else {
				comp = fmt.Sprintf("%s%c", comp, c)
			}
			repeats = 0
		} else {
			repeats++
		}

	}

	//no repeated characters
	if len(comp) == 0 {
		return str
	}

	//our compressed string is longer or the same length
	if len(comp) >= len(str) {
		return str
	}

	return comp

}
