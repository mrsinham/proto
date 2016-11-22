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
	// Tab is \t
	Tab
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

	// Meta components
	Frame
)

// Scanner gives you a scanner capable of dividing the content of the underlying Reader
// into tokens
type Scanner struct {
	r    *bufio.Reader
	line int
	col  int
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
	if ch == '\n' {
		s.line++
		s.col = 0
	}
	s.col++
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
		return s.scanInteger([]rune(lit))
	}

	switch ch {
	case eof:
		return EOF, ""
	case '\t':
		return Tab, ""
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
		//spew.Dump(buf.String(), string(ch), string(s.read()))
		//os.Exit(1)
		return Unknown, buf.String()
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

func (s *Scanner) scanInteger(b []rune) (tok Token, lit string) {
	var buf bytes.Buffer
	for i := range b {
		buf.WriteRune(b[i])
	}
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); !isInteger(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	i := buf.String()
	if len(i) > 0 && i[0] == '0' {
		return Text, i
	}
	return Integer, i
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
	case "error":
		return Error, st
	case "goroutine":
		return Goroutine, st
	case "panic":
		ch := s.read()
		if ch != ':' {
			s.unread()
		} else {
			return Panic, st
		}
	}

	return Text, st

}

func (s *Scanner) GetPosition() (line int, col int) {
	line = s.line
	col = s.col
	return
}
