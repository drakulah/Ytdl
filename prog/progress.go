package prog

import (
	"fmt"
	"ytdl/log"
)

type Progress struct {
	prevLen int
	total   int64
	current int64
}

func New() Progress {
	return Progress{
		prevLen: 0,
		total:   0,
		current: 1,
	}
}

func (p *Progress) SetTotal(t int64) {
	p.total = t
}

func (p *Progress) SetCurrent(c int64) {
	p.current = c
}

func (p Progress) clearLine() {
	fmt.Printf("\r")
	for i := 0; i < p.prevLen; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("\r")
}

func (p Progress) toProperUnit(bytes int64) string {
	unit := ""
	value := float64(bytes)

	switch {
	case bytes >= 1<<30:
		unit = "GB"
		value = value / (1 << 30)
	case bytes >= 1<<20:
		unit = "MB"
		value = value / (1 << 20)
	case bytes >= 1<<10:
		unit = "KB"
		value = value / (1 << 10)
	default:
		unit = "B"
	}

	return fmt.Sprintf("%.2f%s", value, unit)
}

func (p *Progress) Display() {
	p.clearLine()
	perc := (float64(p.current) / float64(p.total)) * 100.0
	displayStr := fmt.Sprintf("Downloading %s%s%s of %s%s%s, %s%.2f%%%s done", log.FgCyan, p.toProperUnit(p.current), log.Reset, log.FgCyan, p.toProperUnit(p.total), log.Reset, log.FgYellow, perc, log.Reset)
	p.prevLen = len(displayStr)
	fmt.Print(displayStr)
}
