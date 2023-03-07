package str

func Substring(str string, start int, end int) string {
	genStr := ""
	strLen := len(str)

	if strLen == 0 || end <= 0 {
		return genStr
	}

	if start < 0 {
		start = 0
	}

	if end < strLen {
		strLen = end
	}

	if start >= strLen {
		return genStr
	}

	for i := start; i < strLen; i++ {
		genStr += string(str[i])
	}

	return genStr

}
