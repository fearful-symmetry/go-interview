package arrays

import "sort"

//Questions regarding permutations of strings

//QUESTION: Given two strings, determine if one is a permutation of the other

type wordSort []rune

func (s wordSort) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s wordSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s wordSort) Len() int {
	return len(s)
}

//FindPermutationSort sorts the two strings, that checks equality
func FindPermutationSort(s1, s2 string) bool {

	//permutations must be of equal length
	if len(s1) != len(s2) {
		return false
	}

	//implent the data interface for sort
	r1 := []rune(s1)
	r2 := []rune(s2)
	sort.Sort(wordSort(r1))
	sort.Sort(wordSort(r2))
	sortedWord1 := string(r1)
	sortedWOrd2 := string(r2)

	return sortedWord1 == sortedWOrd2
}

//FindPermutationMap uses a map to find if two strings are a permutation
func FindPermutationMap(s1, s2 string) bool {

	boolArr := map[rune]int{}

	if len(s1) != len(s2) {
		return false
	}

	//This might look ugly, but there's a reason for it
	//We need to iterate over runes, and not bytes.
	//Because different points in the strings could have runes of different lengths at different points in the string, we need to iterate over each individually
	for _, r := range s1 {
		boolArr[r]++
	}

	for _, r := range s2 {
		boolArr[r]++
	}

	for _, i := range boolArr {
		if i < 2 {
			return false
		}
	}

	return true

}

//FindPermutationArr uses an array to find if two strings are a permutation
func FindPermutationArr(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	//Do something similar to what we did before
	var largest int32
	for _, r := range s1 + s2 {
		if r > largest {
			largest = r
		}
	}
	boolArr := make([]uint8, largest+1)

	//This might look ugly, but there's a reason for it
	//We need to iterate over runes, and not bytes.
	//Because different points in the strings could have runes of different lengths at different points in the string, we need to iterate over each individually
	for _, r := range s1 {
		boolArr[r]++
	}

	for _, r := range s2 {
		boolArr[r]++
	}

	for _, i := range boolArr {
		if i == 1 {
			return false
		}
	}

	return true

}
