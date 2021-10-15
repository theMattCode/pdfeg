package io

import (
	"bufio"
	"io"
	"math"
	"regexp"
)

const (
	eof = rune(0)
)

func isWhitespace(ch rune) bool {
	return ch == SP || ch == CR || ch == LF || ch == HT || ch == FF || ch == NUL
}

type Scanner struct {
	reader            *bufio.Reader
	currentPosition   uint64
	bodyStartPosition uint64
}

func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{reader: bufio.NewReader(reader), currentPosition: 0}
}

func (s *Scanner) read() rune {
	b, err := s.reader.ReadByte()
	s.currentPosition++
	if err != nil {
		return eof
	}
	return rune(b)
}

func (s *Scanner) unread() {
	_ = s.reader.UnreadByte()
	s.currentPosition--
}

func (s *Scanner) Scan() (token Token, literal string, pos uint64) {
	r := s.read()

	switch r {
	case PERCENT:
		return Percent, string(r), s.currentPosition
	}

	return Unknown, string(r), s.currentPosition
}

var pdfVersionRegex = regexp.MustCompile("%PDF-(1\\.[01234567])|(2\\.0)(\r|\r\n|\n)")
var pdfBlobRegex = regexp.MustCompile("%.{4}(\r|\r\n|\n)")

func (s *Scanner) FindStart() (version string, pos uint64, err error) {
	for {
		bytes, err := s.reader.ReadBytes(byte(LF))

		if err != nil {
			return "", math.MaxUint64, err
		}

		s.currentPosition += uint64(len(bytes))

		if pdfVersionRegex.Match(bytes) {
			blobBytes, err := s.reader.ReadBytes(byte(LF))
			if err != nil {
				return "", math.MaxUint64, err
			}
			s.currentPosition += uint64(len(blobBytes))
			if pdfBlobRegex.Match(blobBytes) {
				s.bodyStartPosition = s.currentPosition
				return string(bytes[:len(bytes)-1]), s.currentPosition, nil
			}
		}
	}
}
