package utils

import "testing"

func TestIn(t *testing.T) {
	haystack := []int{1, 2, 3, 4}
	needle := 1
	if !In(haystack, needle) {
		t.Error("fail")
	}
	t.Log("success")
}
