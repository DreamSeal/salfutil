package util

import (
	"regexp"
	"strconv"
)

// 转换8进制utf-8字符串到中文
//var s = `\346\200\241`
//convertOctonaryUtf8(s) => 怡
func convertOctonaryUtf8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)

	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {
			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			return []byte{byte(i)}
		})
	return string(out)
}

//字符串翻转
//reverse("string")  => "gnirts"
func reverse(str string) string {
	runes := []rune(str)
	mid := len(runes)/2
	for i:=0;i<mid;i++ {
		runes[i],runes[len(runes)-i-1] = runes[len(runes)-i-1],runes[i]
	}
	return string(runes)
}