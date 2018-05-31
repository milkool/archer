package archerpackage

import "strconv"

/*
	大寫開頭的function為公開權限
	其他package的程式可以使用
 */
func IntToStr(i int) string {
	str := strconv.Itoa(int(i))

	return str
}