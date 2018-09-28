package arrays

import (
	"testing"
)

func TestCompressString(t *testing.T) {

	str1 := "bbbaacxxxxx"
	str1Out := "b3a2cx5"

	str2 := "necessary"

	if CompressString(str1) != str1Out {
		t.Errorf("wrong answer from %s", str1)
	}

	if CompressString(str2) != str2 {
		t.Errorf("wrong answer from %s", str2)
	}
}
