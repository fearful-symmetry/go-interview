package arrays

import (
	"unicode"
)

//This includes questions dealing with arrays, including strings.
//There are sometimes multiple answers per question, dependng on the question
//=======================================================================================

//QUESTION: Given a string, determine if there are any repeating characters

//UniqueCharInString returns true if all characters in the string are unique are unique, false otherwise
//The second argument allows you to ignore whitespace.
func UniqueCharInString(s string, ignoreWhitespace bool) bool {
	//This is one of those things that's made more complicated by unicode, as opposed to UTF-8 where we can consider everything just an array of all possible ascii characters
	boolArr := map[rune]bool{}

	//iterating like this will force go to iterate over runes instead of individual codepoints
	for _, r := range s {

		if ignoreWhitespace && unicode.IsSpace(r) {
			continue
		}

		if boolArr[r] {
			//is true
			return false
		}
		boolArr[r] = true
	}
	return true
}

//UniqueCharInStringArr is the same as above, but uses a large array instead of a map
func UniqueCharInStringArr(s string, ignoreWhitespace bool) bool {

	//runes are int32
	//This is wildly slow if we just allocated MaxInt32
	//So find the largest known value. Despite the fact that we're iterating twice, this will be faster, as we have no hashmaps.
	var largest int32
	for _, r := range s {
		if r > largest {
			largest = r
		}
	}
	boolArr := make([]bool, largest+1)

	for _, r := range s {
		if ignoreWhitespace && unicode.IsSpace(r) {
			continue
		}

		if boolArr[r] {
			return false
		}
		boolArr[r] = true
	}

	return true
}
