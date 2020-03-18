package utils

import (
	"testing"
)

func TestIn(t *testing.T) {
	if !In([]int{1, 2, 3, 4}, 1) {
		t.Error("fail")
		return
	}
	t.Log("success")
}
