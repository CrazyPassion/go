package helloworld

import (
	"bytes"
	"strings"
)

func Basename(s string) string {
	if pos := strings.LastIndex(s, "/"); pos >= 0 {
		s = s[pos+1:]
		if pos = strings.LastIndex(s, "."); pos >= 0 {
			s = s[:pos]
			pos = 15
		}
	}
	return s
}

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func Comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	loop := (n - 1) / 3
	for i := 0; i < loop; i++ {
		s = s[:n-3*(i+1)] + "," + s[n-3*(i+1):]
	}
	return s
}

func Comma2(s string) string {
	buf := &bytes.Buffer{}
	pos := 0

	comma := []byte{','}
	if len(s)%3 != 0 {
		pos += len(s) % 3
		buf.WriteString(s[:pos])
		buf.Write(comma)
	}

	for ; pos < len(s); pos += 3 {
		buf.WriteString(s[pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}
