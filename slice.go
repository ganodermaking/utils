package utils

import (
	"reflect"
)

// InArray 元素、切片或数组是否在切片或数组内
func InArray(haystack, needle interface{}) bool {
	hVal := reflect.ValueOf(haystack)
	hKind := hVal.Kind()

	nVal := reflect.ValueOf(needle)
	nKind := nVal.Kind()

	if hKind != reflect.Slice && hKind != reflect.Array {
		return false
	}

	var flag bool
	switch nKind {
	case reflect.Slice:
		flag = sliInArray(hVal, nVal)
		break
	case reflect.Array:
		flag = sliInArray(hVal, nVal)
		break
	default:
		flag = eleInArray(hVal, needle)
	}

	return flag
}

// sliInArray ...
func sliInArray(haystack, needle reflect.Value) bool {
	needleInMap := map[int]bool{}
	for i := 0; i < needle.Len(); i++ {
		needleInMap[i] = false
		for j := 0; j < haystack.Len(); j++ {
			if haystack.Index(j).Interface() == needle.Index(i).Interface() {
				needleInMap[i] = true
				break
			}
		}
	}

	for _, item := range needleInMap {
		if item == false {
			return false
		}
	}

	return true
}

// eleInArray ...
func eleInArray(haystack reflect.Value, needle interface{}) bool {
	for i := 0; i < haystack.Len(); i++ {
		if haystack.Index(i).Interface() == needle {
			return true
		}
	}
	return false
}
