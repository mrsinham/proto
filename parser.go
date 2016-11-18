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

	var currStep *Step
	// scanning steps
	stepLoop:
	for {
		tok, _= p.scan()
		if tok == NewLine {

			currStep, _ = p.scanStep()
			spew.Dump(currStep)

			if currStep != nil {
				r.Stacktrace = append(r.Stacktrace, currStep)
			}

			tok, _ = p.scan()


			spew.Dump(tok)
			if tok == NewLine {
				tok, _ = p.scan()
				if tok != Text {
					break stepLoop
				}
				p.unscan()
			}
			p.unscan()

		}

	}
	spew.Dump(r)


	return r

}

func (p *Parser) scanStep() (*Step, error) {

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
		if tok == OpeningParenthese {
			p.unscan()
			break
		}
		buf.WriteString(lit)
	}

	f.Method = buf.String()

	buf.Reset()

	tok, lit = p.scan()
	if tok != OpeningParenthese {
		return nil, nil
	}

	var args []string

	// scanning args
	for {
		tok, lit = p.scan()
		if tok == ClosingParenthese {
			break
		}
		if tok == Pointer {
			args = append(args, lit)
		}
	}

	f.Args = args

	tok, lit = p.scan()
	if tok != NewLine {
		return nil, nil
	}

	tok, lit = p.scan()
	if tok != Tab {
		return nil, nil
	}


	// scanning location
	for {
		tok, lit = p.scan()
		// wtf ?
		if tok == NewLine {
			return nil, nil
		}

		if tok == Colon {
			break
		}
		buf.WriteString(lit)
	}

	f.Location = buf.String()

	buf.Reset()

	tok, lit = p.scan()
	if tok != Integer {
		return nil, nil
	}

	f.Line, _ = strconv.Atoi(lit)

	// space
	p.scan()
	// scan +
	p.scan()
	// scan pointer
	p.scan()

	return f, nil

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
