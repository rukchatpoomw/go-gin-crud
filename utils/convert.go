package utils

import "strconv"

func ConvertStringToInt(str string, defaultValue ...int) int {

	num, err := strconv.Atoi(str)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return num
}

func ConvertStringToInt64(str string, defaultValue ...int64) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return num
}
