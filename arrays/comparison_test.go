package arrays

import "testing"

func TestIsOneEdit(t *testing.T) {
	str1 := "cat"
	str2 := "bat"

	str3 := "at"

	str4 := "bed"
	str5 := "bend"

	str6 := "🔥🔥"
	str7 := "🔥🐯🔥"

	str8 := "tiger"
	str9 := "fire"

	if !IsOneEdit(str1, str2) {
		t.Errorf("wrong result: %s and %s", str1, str2)
	}

	if !IsOneEdit(str1, str3) {
		t.Errorf("wrong result: %s and %s", str1, str3)
	}

	if !IsOneEdit(str4, str5) {
		t.Errorf("wrong result: %s and %s", str4, str5)
	}

	if !IsOneEdit(str6, str7) {
		t.Errorf("wrong result: %s and %s", str6, str7)
	}

	if IsOneEdit(str8, str9) {
		t.Errorf("wrong result: %s and %s", str8, str9)
	}
}
