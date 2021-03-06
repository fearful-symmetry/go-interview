package arrays

import "strings"

//QUESTION: determine if a string is one edit away from another string

//check to see if the change is a replacement
//assumes strings are the same length
//Returns true if *a single char* has been replaced, false otherwise
func isReplacement(s1, s2 []rune) bool {
	//convert to rune arrays
	//keep track of changes. If there are more than one, bail out
	var foundDiff = false
	for index := 0; index < len(s1); index++ {
		if s1[index] != s2[index] {
			if foundDiff {
				return false
			}
			foundDiff = true
		}
	}

	return foundDiff
}

//because this isn just theory, we're skipping error handling. Please make sure `longer` is longer.
func isInsertChange(longer, shorter []rune) bool {
	//This is...kind of a confusing algo
	li, si := 0, 0
	for li < len(longer) && si < len(shorter) {
		if longer[li] != shorter[si] {
			if li != si {
				return false
			}
			li++
		} else {
			li++
			si++
		}
	}
	return true
}

//IsOneEdit checks three conditions: character replacement, insertion, or removal
func IsOneEdit(s1, s2 string) bool {

	//since we're spending all our time iterating through stuff, just convert our strings to rune slices
	r1 := []rune(s1)
	r2 := []rune(s2)
	if len(r1) == len(r2) {
		return isReplacement(r1, r2)
	} else if len(r1)+1 == len(r2) {
		return isInsertChange(r2, r1)
	} else if len(r2)+1 == len(r1) {
		return isInsertChange(r1, r2)
	}

	return false
}

// Question: Check to see if a string is a rotation of another string. For example, obcatb is a rotation of bobcat, rotated after the first b. Other example: secretarybird and tarybirdsecre, rotated after the "secre"

//CheckIsRotationContains takes two strings, s1 and s2, and determines if s1 is a rotation of s2
func CheckIsRotationContains(s1, s2 string) bool {

	//this is the 'clever' way to do it: cat the substring together, see if the check string contains the cat'ed string
	s1Cat := s1 + s1
	return strings.Contains(s1Cat, s2)
}
