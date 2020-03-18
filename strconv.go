package utils

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

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
