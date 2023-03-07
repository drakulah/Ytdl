package str

func StartsWith(src string, startsWith string) bool {
	startsWithLen, ret := len(startsWith), false
	if srcLen := len(src); srcLen == 0 || startsWithLen == 0 || srcLen < startsWithLen {
		return ret
	}
	if srcFirstLetters := src[0:startsWithLen]; startsWith == srcFirstLetters {
		ret = true
	}
	return ret
}
