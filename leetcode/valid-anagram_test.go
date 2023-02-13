package algorithms

import "testing"

func TestIsAnagram(t *testing.T) {
	var str1 string
	var str2 string

	str1 = "anagram"
	str2 = "nagaram"
	if ok := IsAnagram(str1, str2); ok != true {
		t.Errorf("IsAnagram(%s, %s) expected: true, got: %t", str1, str2, ok)
	}

	str1 = "anagr"
	str2 = "nagaram"
	if ok := IsAnagram(str1, str2); ok != false {
		t.Errorf("IsAnagram(%s, %s) expected: false, got: %t", str1, str2, ok)
	}

	str1 = ""
	str2 = ""
	if ok := IsAnagram(str1, str2); ok != true {
		t.Errorf("IsAnagram(%s, %s) expected: true, got: %t", str1, str2, ok)
	}
}
