package utils

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

// Copy 复制
func Copy(toValue, fromValue interface{}) {
	if err := copier.Copy(toValue, fromValue); err != nil {
		logrus.Error("copy error: ", err)
	}
}
