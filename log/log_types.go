package log

import (
	"fmt"
	"strings"
)

func W(message ...string) {
	logTxt := fmt.Sprintf("%sWarning:%s %s", FgYellow, Reset, strings.Join(message, " "))
	fmt.Println(logTxt)
}

func E(message ...string) {
	logTxt := fmt.Sprintf("%sError:%s %s", fgRed, Reset, strings.Join(message, " "))
	fmt.Println(logTxt)
}
