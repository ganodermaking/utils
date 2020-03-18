package utils

import (
	"reflect"
)

// In 元素是否在切片或数组内
func In(haystack interface{}, needle interface{}) bool {
	refVal := reflect.ValueOf(haystack)
	kind := refVal.Kind()

	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < refVal.Len(); i++ {
			if refVal.Index(i).Interface() == needle {
				return true
			}
		}
	}

	return false
}
