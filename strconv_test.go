package utils

import (
	"reflect"
	"testing"
)

func TestParseInt(t *testing.T) {
	var str = "110"
	i := ParseInt(str)
	refI := reflect.ValueOf(i)
	if refI.Kind() != reflect.Int32 {
		t.Error("fail")
		return
	}
	t.Log("success")
}

func TestAtoi(t *testing.T) {
	var str = "110"
	i := Atoi(str)
	refI := reflect.ValueOf(i)
	if refI.Kind() != reflect.Int32 {
		t.Error("fail")
		return
	}
	t.Log("success")
}
