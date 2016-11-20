package parser

import (
	"bytes"
	"io"
	"strconv"

	"fmt"

	"github.com/davecgh/go-spew/spew"
	"os"
	"errors"

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
	case "IO wait":
		e = EventIOWait
	case "chan receive":
		e = EventChanReceive
	case "chan send":
		e = EventChanSend
	case "select":
		e = EventSelect
	case "sleep":
		e = EventSleep
	case "semacquire":
		e = EventSemAcquire
	case "runnable":
		e = EventRunnable
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
	var err error
	// scanning steps
	stepLoop:
	for {
		tok, _= p.scan()
		if tok == NewLine {

			currStep, err = p.scanStep()
			if err != nil {
				// TODO: if error then loop eternally
				spew.Dump(err)
				os.Exit(1)
			}

			if currStep != nil {
				r.Stacktrace = append(r.Stacktrace, currStep)
			}

			tok, _ = p.scan()

			if tok == NewLine {
				tok, lit = p.scan()
				// end of trace or created by mention
				if tok != Text || lit == "created" {
					break stepLoop
				}
				p.unscan()
			}
			p.unscan()

		}

	}

	var buf bytes.Buffer

	if tok == Text && lit == "created" {
		//
		tok, lit = p.scanWithoutSpaces()
		if tok != Text && lit != "by" {
			return nil
		}

		// whitespace
		p.scan()
		cb := &CreatedBy{}

		createdLoop:
		for {
			tok, lit = p.scan()
			if tok != Text && tok != Dot {
				p.unscan()
				break createdLoop
			}
			buf.WriteString(lit)
		}

		cb.Method = buf.String()
		buf.Reset()

		cb.Location, cb.Line = p.scanLocation()
		if cb.Location == "" {
			return nil
		}

		r.CreatedBy = cb

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
			tok,lit = p.scan()
			if tok == Pointer || tok == ClosingParenthese{
				p.unscan()
				p.unscan()
				break
			}
//			if tok == ClosingParenthese {
//				p.un
//			}
			buf.WriteString("(")
		}
		buf.WriteString(lit)
	}

	// TODO: problem with func 018
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


	f.Location, f.Line = p.scanLocation()
	if f.Location == "" || f.Line == 0 {
		return nil, errors.New("unable to scan location")
	}
	spew.Dump(f)

	return f, nil

}

func (p *Parser) scanLocation() (location string, line int) {
	var buf bytes.Buffer
	var tok Token
	var lit string
	tok, lit = p.scan()
	if tok != NewLine {
		return
	}

	// scanning location
	for {
		tok, lit = p.scanWithoutSpaces()
		// wtf ?
		if tok == NewLine {
			return
		}

		if tok == Colon {
			break
		}
		buf.WriteString(lit)
	}

	location = buf.String()


	buf.Reset()

	tok, lit = p.scan()
	if tok != Integer {
		return
	}

	line, _ = strconv.Atoi(lit)

	// space
	p.scan()
	// scan +
	p.scan()
	// scan pointer
	p.scan()

	return

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
