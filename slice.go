package utils

import (
	"reflect"
)

// In 元素是否在切片或数组内
func In(haystack interface{}, needle interface{}) bool {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()

	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == needle {
				return true
			}
		}
	}

	return false
}
