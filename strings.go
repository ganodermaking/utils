package utils

// Substr 截取字符串
func Substr(str string, sl ...int) string {
	var start = 0
	var end = 0
	var length = len(str)

	if len(sl) == 1 {
		start = sl[0]
	}

	if len(sl) == 2 {
		start = sl[0]

		if sl[1] <= length {
			end = start + sl[1]
		}

		if sl[1] > length {
			end = length
		}
	}

	newStr := str[start:end]
	return newStr
}

// MbSubstr 截取字符串（包含中文字符）
func MbSubstr(str string, sl ...int) string {
	var start = 0
	var end = 0
	var length = len(str)

	if len(sl) == 1 {
		start = sl[0]
	}

	if len(sl) == 2 {
		start = sl[0]

		if sl[1] <= length {
			end = start + sl[1]
		}

		if sl[1] > length {
			end = length
		}
	}

	newStr := string([]rune(str)[start:end])
	return newStr
}
