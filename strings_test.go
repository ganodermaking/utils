package utils

import "testing"

func TestSubstr(t *testing.T) {
	want := "5678"
	res := Substr("0123456789X", 5, 4)
	if res != want {
		t.Errorf("res:%s, want:%s", res, want)
	}
}

func TestMbSubstr(t *testing.T) {
	want := "中国67"
	res := MbSubstr("01234中国6789X", 5, 4)
	if res != want {
		t.Errorf("res:%s, want:%s", res, want)
	}
}
