package jcosmos

import (
	"bytes"
	"io"
)

func bodyToStr(body io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	return buf.String()
}
