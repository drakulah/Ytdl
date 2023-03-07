package log

import (
	"fmt"
	"strings"
)

func W(message ...string) {
	logTxt := fmt.Sprintf("%s%s W %s %s", fgWhite, bgYellow, reset, strings.Join(message, " "))
	fmt.Println(logTxt)
}

func I(message ...string) {
	logTxt := fmt.Sprintf("%s%s I %s %s", fgWhite, bgCyan, reset, strings.Join(message, " "))
	fmt.Println(logTxt)
}

func E(message ...string) {
	logTxt := fmt.Sprintf("%s%s E %s %s", fgWhite, bgRed, reset, strings.Join(message, " "))
	fmt.Println(logTxt)
}

func S(message ...string) {
	logTxt := fmt.Sprintf("%s%s E %s %s", fgWhite, bgGreen, reset, strings.Join(message, " "))
	fmt.Println(logTxt)
}
