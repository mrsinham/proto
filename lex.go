package parser

import (
	"bufio"
	"bytes"
	"io"
)

// Token is an analyzed part of a content
type Token int

const (

	// Misc

	// Whitespace is space or tab
	Whitespace Token = iota

	// NewLine is \n
	NewLine
	// Integer is 0-9+
	Integer
	// Text is the default value
	Text
	// EOF is the end of the content
	EOF
	// Slash well is Slash
	Slash
	// Unknown is when the analysis didnt succeed
	Unknown

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

	// Info
	Pointer
)

// Scanner gives you a scanner capable of dividing the content of the underlying Reader
// into tokens
type Scanner struct {
	r *bufio.Reader
}

// NewScanner gives you a new scanner
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

// Scan gives the next token that it fetchs into the io.Reader
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

	if ch == '0' {
		s.unread()
		tok, lit = s.scanPointer()
		if tok == Pointer {
			return
		}
	}

	if isInteger(ch) {
		s.unread()
		return s.scanInteger()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '\n':
		return NewLine, ""
	case '/':
		return Slash, string(ch)
	case '[':
		return OpeningBracket, string(ch)
	case ']':
		return ClosingBracket, string(ch)
	case '(':
		return OpeningParenthese, string(ch)
	case ')':
		return ClosingParenthese, string(ch)
	case ':':
		return Colon, string(ch)
	case '.':
		return Dot, string(ch)
	case '=':
		return Equal, string(ch)
	}

	return Text, string(ch)
}

func (s *Scanner) scanPointer() (tok Token, lit string) {

	var buf bytes.Buffer
	// be sure of the beginning
	ch := s.read()
	if ch != '0' {
		s.unread()
		return Unknown, ""
	}

	buf.WriteRune(ch)

	if ch = s.read(); ch != 'x' {
		// return this char
		s.unread()
		//  return 0
		s.unread()
		return Unknown, ""
	}

	buf.WriteRune(ch)

	for {
		ch = s.read()
		if !isHexa(ch) {
			s.unread()
			break
		}
		buf.WriteRune(ch)
	}

	return Pointer, buf.String()
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
