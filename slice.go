package utils

import (
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"
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

// ParseInt 字符串转int64
func ParseInt(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logrus.Error(err)
	}
	return i
}

// Atoi 字符串转int32
func Atoi(str string) int32 {
	i, err := strconv.Atoi(str)
	if err != nil {
		logrus.Error(err)
	}
	return int32(i)
}
