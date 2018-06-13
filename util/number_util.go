package util

import (
	"regexp"
	"strconv"
)

func ZhNumToEn(zhNum string) int {
	zhRegexp, _ := regexp.Compile("[\\d|.]+万")
	numRegexp, _ := regexp.Compile("[\\d|.]+")
	number := 0
	if zhRegexp.MatchString(zhNum) {
		strNum := numRegexp.FindAllString(zhNum, 1)[0]
		num, err := strconv.ParseFloat(strNum, 64)
		if err == nil {
			number = int(num * 10000)
		}
	} else if numRegexp.MatchString(zhNum) {
		num, err := strconv.ParseFloat(zhNum, 64)
		if err == nil {
			number = int(num)
		}
	}
	return number
}
