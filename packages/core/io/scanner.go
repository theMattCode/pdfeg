package io

func IsWhitespace(ch rune) bool {
	return ch == SP || ch == CR || ch == LF || ch == HT || ch == FF || ch == NUL
}
