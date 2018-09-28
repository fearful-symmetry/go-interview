package arrays

import "testing"

func TestFindPermutationSort(t *testing.T) {

	str1 := "god"
	str2 := "dog"
	str3 := "desserts"
	str4 := "stressed"
	str5 := "cat"

	if FindPermutationSort(str1, str2) != true {
		t.Fatalf("Failed pair %s, %s", str1, str2)
	}
	if FindPermutationSort(str3, str4) != true {
		t.Fatalf("Failed pair %s, %s", str3, str4)
	}
	if FindPermutationSort(str1, str5) != false {
		t.Fatalf("Failed pair %s, %s", str1, str5)
	}

}

func TestFindPermutationMap(t *testing.T) {

	str1 := "god"
	str2 := "dog"
	str3 := "desserts"
	str4 := "stressed"
	str5 := "cat"

	if FindPermutationMap(str1, str2) != true {
		t.Fatalf("Failed pair %s, %s", str1, str2)
	}
	if FindPermutationMap(str3, str4) != true {
		t.Fatalf("Failed pair %s, %s", str3, str4)
	}
	if FindPermutationMap(str1, str5) != false {
		t.Fatalf("Failed pair %s, %s", str1, str5)
	}

}

func TestFindPermutationArr(t *testing.T) {

	str1 := "god"
	str2 := "dog"
	str3 := "desserts"
	str4 := "stressed"
	str5 := "cat"

	if FindPermutationArr(str1, str2) != true {
		t.Fatalf("Failed pair %s, %s", str1, str2)
	}
	if FindPermutationArr(str3, str4) != true {
		t.Fatalf("Failed pair %s, %s", str3, str4)
	}
	if FindPermutationArr(str1, str5) != false {
		t.Fatalf("Failed pair %s, %s", str1, str5)
	}

}
