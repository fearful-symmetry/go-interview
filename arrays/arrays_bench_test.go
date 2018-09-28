package arrays

import "testing"

//Benchmarks for the array questions

func BenchmarkUniqueCharInString(b *testing.B) {
	str := "Mr. Jock TV quiz PhD, bags few lynx"

	for index := 0; index < b.N; index++ {
		UniqueCharInString(str, true)
	}
}

func BenchmarkUniqueCharInStringArr(b *testing.B) {
	str := "Mr. Jock TV quiz PhD, bags few lynx"

	for index := 0; index < b.N; index++ {
		UniqueCharInStringArr(str, true)
	}
}

func BenchmarkFindPermutationMap(b *testing.B) {
	str3 := "desserts"
	str4 := "stressed"
	for index := 0; index < b.N; index++ {
		FindPermutationMap(str3, str4)
	}
}

func BenchmarkFindPermutationSort(b *testing.B) {
	str3 := "desserts"
	str4 := "stressed"
	for index := 0; index < b.N; index++ {
		FindPermutationSort(str3, str4)
	}
}

func BenchmarkFindPermutationArr(b *testing.B) {
	str3 := "desserts"
	str4 := "stressed"
	for index := 0; index < b.N; index++ {
		FindPermutationArr(str3, str4)
	}
}
