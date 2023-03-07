package conv

func StringToInt(str string) int {
	number, strLen := 0, len(str)
	for i := 0; i < strLen; i++ {
		remainder := str[i]
		if remainder == '1' {
			number = (number * 10) + 1
		} else if remainder == '2' {
			number = (number * 10) + 2
		} else if remainder == '3' {
			number = (number * 10) + 3
		} else if remainder == '4' {
			number = (number * 10) + 4
		} else if remainder == '5' {
			number = (number * 10) + 5
		} else if remainder == '6' {
			number = (number * 10) + 6
		} else if remainder == '7' {
			number = (number * 10) + 7
		} else if remainder == '8' {
			number = (number * 10) + 8
		} else if remainder == '9' {
			number = (number * 10) + 9
		} else {
			number = number * 10
		}
	}
	return number
}

func StringToInt64(str string) int64 {
	number, strLen := 0, len(str)
	for i := 0; i < strLen; i++ {
		remainder := str[i]
		if remainder == '1' {
			number = (number * 10) + 1
		} else if remainder == '2' {
			number = (number * 10) + 2
		} else if remainder == '3' {
			number = (number * 10) + 3
		} else if remainder == '4' {
			number = (number * 10) + 4
		} else if remainder == '5' {
			number = (number * 10) + 5
		} else if remainder == '6' {
			number = (number * 10) + 6
		} else if remainder == '7' {
			number = (number * 10) + 7
		} else if remainder == '8' {
			number = (number * 10) + 8
		} else if remainder == '9' {
			number = (number * 10) + 9
		} else {
			number = number * 10
		}
	}
	return int64(number)
}
