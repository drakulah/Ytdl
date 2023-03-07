package conv

func IntToString(number int) string {
	str := ""
	for number != 0 {
		remainder := number % 10
		number = int(number / 10)
		if remainder == 1 {
			str = "1" + str
		} else if remainder == 2 {
			str = "2" + str
		} else if remainder == 3 {
			str = "3" + str
		} else if remainder == 4 {
			str = "4" + str
		} else if remainder == 5 {
			str = "5" + str
		} else if remainder == 6 {
			str = "6" + str
		} else if remainder == 7 {
			str = "7" + str
		} else if remainder == 8 {
			str = "8" + str
		} else if remainder == 9 {
			str = "9" + str
		} else {
			str = "0" + str
		}
	}
	return str
}

func Int64ToString(number int64) string {
	str := ""
	for number != 0 {
		remainder := number % 10
		number = int64(number / 10)
		if remainder == 1 {
			str = "1" + str
		} else if remainder == 2 {
			str = "2" + str
		} else if remainder == 3 {
			str = "3" + str
		} else if remainder == 4 {
			str = "4" + str
		} else if remainder == 5 {
			str = "5" + str
		} else if remainder == 6 {
			str = "6" + str
		} else if remainder == 7 {
			str = "7" + str
		} else if remainder == 8 {
			str = "8" + str
		} else if remainder == 9 {
			str = "9" + str
		} else {
			str = "0" + str
		}
	}
	return str
}
