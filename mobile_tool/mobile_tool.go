package mobile_tool

import "regexp"

// 手机号真伪判断
func MobileVerification(mobile string) bool {
	return regexp.MustCompile(`^1(3|4|5|6|7|8|9)\d{9}$`).MatchString(mobile)
}

// 手机号中间中间四位标*
func MobileNote(mobile string) string {
	return string(mobile[:3]) + "****" + string(mobile[7:])
}
