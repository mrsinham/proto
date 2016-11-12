package parser

import (
	"bufio"
	"bytes"
	"io"
)

type Token int

const (

	// Misc
	Whitespace Token = iota
	Integer
	Text
	Unknown
	EOF
	Slash

	// symbols
	OpeningBracket
	ClosingBracket
	OpeningParenthese
	ClosingParenthese
	Colon
	Dot
	Equal

	// Keywords
	Goroutine
	Panic
	Recovered
	Runtime
	Error
	Created
	By

	// Event
	Running
	IOWait
	Chan
	Receive
	Send
	Syscall
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

var eof = rune(0)

func (s *Scanner) read() rune {

	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	}

	if isLetter(ch) {
		s.unread()
		return s.scanIdentifiers()
	}

	if isInteger(ch) {
		s.unread()
		return s.scanInteger()
	}

	if ch == eof {
		return EOF, ""
	}

	if ch == '/' {
		return Slash, string(ch)
	}

	if ch == '[' {
		return OpeningBracket, string(ch)
	}

	if ch == ']' {
		return ClosingBracket, string(ch)
	}

	if ch == '(' {
		return OpeningParenthese, string(ch)
	}
	if ch == ')' {
		return ClosingParenthese, string(ch)
	}
	if ch == ':' {
		return Colon, string(ch)
	}
	if ch == '.' {
		return Dot, string(ch)
	}
	if ch == '=' {
		return Equal, string(ch)
	}

	return Text, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Whitespace, buf.String()
}

func (s *Scanner) scanInteger() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); !isInteger(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return Integer, buf.String()
}

func (s *Scanner) scanIdentifiers() (tok Token, lit string) {
	var buf bytes.Buffer
	// we put the current char into it
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); !isLetter(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}

	}

	st := buf.String()

	switch st {
	case "panic":
		return Panic, st
	case "recovered":
		return Recovered, st
	case "error":
		return Error, st
	case "goroutine":
		return Goroutine, st
	case "created":
		return Created, st
	case "by":
		return By, st
	case "runtime":
		return Runtime, st
	case "syscall":
		return Syscall, st
	case "running":
		return Running, st
	case "chan":
		return Chan, st
	case "receive":
		return Receive, st
	}

	return Text, st

}
