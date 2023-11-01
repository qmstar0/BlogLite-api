package utils

import "strconv"

// Offset 计算偏移
func Offset(page string, limit string) (limitInt int, offset int) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	limitInt, err = strconv.Atoi(limit)
	if err != nil {
		limitInt = 13
	}

	return limitInt, (pageInt - 1) * limitInt
}
