package jcosmos

import (
	"fmt"
)

type loglevel int

const (
	LogLevelTrace loglevel = 5
	LogLevelDebug loglevel = 4
	LogLevelInfo  loglevel = 3
	LogLevelWarn  loglevel = 2
	LogLevelError loglevel = 1
	LogLevelNone  loglevel = 0
)

func getLogLevelLabel(l loglevel) string {
	switch l {
	case LogLevelTrace:
		return "Trace"
	case LogLevelDebug:
		return "Debug"
	case LogLevelInfo:
		return "Info"
	case LogLevelWarn:
		return "Warn"
	case LogLevelError:
		return "Error"
	case LogLevelNone:
		return "None"
	default:
		return "unknown"
	}
}

func (c Jcosmos) logReq(rl, pk, method, body string, headers map[string]string) {
	var hString string
	for k, v := range headers {
		hString += fmt.Sprintf("%s:%s\n", k, v)
	}
	msg := fmt.Sprintf("%s %s[%s] \n%s\n%s", method, rl, pk, body, hString)
	c.log(LogLevelTrace, msg)
}

func (c Jcosmos) log(level loglevel, msg string) {
	if c.loglevel == LogLevelNone {
		return
	}
	if level <= c.loglevel {
		c.logger.Printf("[%s]\t%s", getLogLevelLabel(level), msg)
	}

}
