package arrays

import "testing"

func TestUniqueCharInString(t *testing.T) {

	str1 := "The bearded vulture lives on a diet of 80%-90% bones."
	str2 := "subdermatoglyphic"
	str3 := "Mr. Jock TV quiz PhD, bags few lynx"
	str4 := "Heizölrückstoßabdämpfung"
	str5 := "🔥 🌊"

	if UniqueCharInString(str1, false) {
		t.Fatalf("Failed string %s", str1)
	}

	if !UniqueCharInString(str2, false) {
		t.Fatalf("Failed string %s", str2)
	}

	if !UniqueCharInString(str3, true) {
		t.Fatalf("Failed string %s", str3)
	}

	if !UniqueCharInString(str4, false) {
		t.Fatalf("Failed string %s", str4)
	}

	if !UniqueCharInString(str5, false) {
		t.Fatalf("Failed string %s", str5)
	}

}

func TestUniqueCharInStringArr(t *testing.T) {
	str1 := "The bearded vulture lives on a diet of 80%-90% bones."
	str2 := "subdermatoglyphic"
	str3 := "Mr. Jock TV quiz PhD, bags few lynx"
	str4 := "Heizölrückstoßabdämpfung"
	str5 := "🔥 🌊"

	if UniqueCharInStringArr(str1, false) {
		t.Fatalf("Failed string %s", str1)
	}

	if !UniqueCharInStringArr(str2, false) {
		t.Fatalf("Failed string %s", str2)
	}

	if !UniqueCharInStringArr(str3, true) {
		t.Fatalf("Failed string %s", str3)
	}

	if !UniqueCharInStringArr(str4, false) {
		t.Fatalf("Failed string %s", str4)
	}

	if !UniqueCharInString(str5, false) {
		t.Fatalf("Failed string %s", str5)
	}
}
