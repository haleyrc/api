package log

import (
	"fmt"
	"log"
	"strings"
)

var debug = false

const (
	debugLevel = "[DEBUG]"
)

func SetDebug(val bool) {
	debug = val
}

func Debug(args ...interface{}) {
	if !debug {
		return
	}
	p(debugLevel, args...)
}

func p(level string, args ...interface{}) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s ", level)
	fmt.Fprint(&sb, args...)
	log.Print(sb.String())
}

func pf(level string, format string, args ...interface{}) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s ", level)
	fmt.Fprintf(&sb, format, args...)
	log.Print(sb.String())
}

func Debugf(format string, args ...interface{}) {
	if !debug {
		return
	}
	pf(debugLevel, format, args...)
}
