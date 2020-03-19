package utils

import (
	"testing"
)

func TestIn(t *testing.T) {
	if !InArray([]int{1, 2, 3, 4}, 1) {
		t.Error("fail1")
		return
	}

	if !InArray([]int{1, 2, 3, 4}, []int{1, 3}) {
		t.Error("fail2")
		return
	}
	t.Log("success")
}
