package parser

import (
	"bytes"
	"io"
	"strconv"

	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Parser struct {
	s   *Scanner
	buf struct {
		last []Token
		lit  []string
		n    int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		s: NewScanner(r),
	}
}

func (p *Parser) Parse() (*State, error) {

	s := &State{}
parsingLoop:
	for {
		var tok Token
		//var lit string
		tok, _ = p.scan()
		switch tok {
		case EOF:
			break parsingLoop
		case Goroutine:
			p.unscan()
			p.parseRoutine()
			//os.Exit(1)
		}

	}
	return s, nil
}

func (p *Parser) parseRoutine() *Routine {

	var tok Token
	var lit string
	tok, _ = p.scan()

	if tok != Goroutine {
		p.unscan()
		return nil
	}

	r := &Routine{}

	tok, lit = p.scanWithoutSpaces()
	if tok != Integer {
		p.unscan()
		return nil
	}

	// we already know its an integer
	// scan the routine ID
	r.ID, _ = strconv.Atoi(lit)

	tok, lit = p.scanWithoutSpaces()
	if tok != OpeningBracket {
		p.unscan()
		return nil
	}

	var event bytes.Buffer

	for {
		tok, lit = p.scan()
		if tok == ClosingBracket {
			break
		}
		event.WriteString(lit)
	}

	var e Event
	switch event.String() {
	case "running":
		e = EventRunning
	case "syscall":
		e = EventSyscall
	case "IO Wait":
		e = EventIOWait
	case "chan receive":
		e = EventChanReceive
	case "chan send":
		e = EventChanSend
	case "select":
		e = EventSelect
	default:
		return nil
	}

	r.Event = e

	tok, lit = p.scan()
	if tok != Colon {
		return nil
	}
	spew.Dump(r)

	p.scanFrame()

	return r

}

func (p *Parser) scanFrame() (*Step, error) {

	var buf bytes.Buffer
	// get the first text
	var tok Token
	var lit string
	tok, lit = p.scanWithoutSpaces()
	if tok != Text {
		return nil, fmt.Errorf("waiting text, received %v", tok.String())
	}

	f := &Step{}
	buf.WriteString(lit)

	for {
		tok, lit = p.scan()
		if tok != Text && tok != Dot {
			p.unscan()
			break
		}
		buf.WriteString(lit)
	}

	f.Method = buf.String()

	buf.Reset()

	tok, lit = p.scan()
	if tok != OpeningParenthese {
		return nil
	}

	spew.Dump(f)
	return nil, nil

}

func (p *Parser) scanWithoutSpaces() (tok Token, lit string) {
	for {
		tok, lit = p.scan()
		if tok != Whitespace && tok != NewLine {
			break
		}
	}
	return
}

func (p *Parser) scan() (tok Token, lit string) {

	if p.buf.n > 0 {
		tok, lit = p.buf.last[len(p.buf.last)-p.buf.n], p.buf.lit[len(p.buf.lit)-p.buf.n]
		p.buf.n--
		return
	}

	tok, lit = p.s.Scan()

	p.buf.last = append(p.buf.last, tok)
	p.buf.lit = append(p.buf.lit, lit)

	return
}

func (p *Parser) unscan() {
	if p.buf.n+1 <= len(p.buf.last) {
		p.buf.n++
	}
}
