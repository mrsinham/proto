package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

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
	var err error
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
			var r *Routine
			r, err = p.parseRoutine()
			if err != nil {
				line, col := p.s.GetPosition()
				log.Fatalf("error found while parsing trace at line %v col %v : %v", line, col, err)
			}
			spew.Dump(r)
			//os.Exit(1)
		}

	}
	return s, nil
}

func (p *Parser) parseRoutine() (r *Routine, err error) {

	var tok Token
	var lit string
	tok, _ = p.scan()
	r = &Routine{}

	if tok != Goroutine {
		p.unscan()
		return nil, errors.New("goroutine keyword not found")
	}

	tok, lit = p.scanWithoutSpaces()
	if tok != Integer {
		p.unscan()
		return nil, errors.New("goroutine id not found")
	}

	// we already know its an integer
	// scan the routine ID
	r.ID, _ = strconv.Atoi(lit)

	tok, lit = p.scanWithoutSpaces()
	if tok != OpeningBracket {
		p.unscan()
		return nil, errors.New("opening event bracket not found")
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
		return nil, fmt.Errorf("event not standard %q", event.String())
	}

	r.Event = e

	tok, lit = p.scan()
	if tok != Colon {
		return nil, errors.New("colon was expected after event")
	}

	var currStep *Step
	// scanning steps
stepLoop:
	for {
		tok, _ = p.scan()
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
		tok, lit = p.scanWithoutSpaces()
		if tok != Text && lit != "by" {
			return nil, errors.New("was expecting 'created by'")
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

		cb.Location, cb.Line, err = p.scanLocation()
		if err != nil {
			return nil, err
		}

		r.CreatedBy = cb

	}

	spew.Dump(r)

	return r, nil

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

	s := &Step{}
	buf.WriteString(lit)

	for {
		tok, lit = p.scan()
		if tok == OpeningParenthese {
			tok, lit = p.scan()
			if tok == Pointer || tok == ClosingParenthese {
				p.unscan()
				p.unscan()
				break
			}
			buf.WriteString("(")
		}
		buf.WriteString(lit)
	}

	// TODO: problem with func 018
	s.Method = buf.String()

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

	s.Args = args

	var err error
	s.Location, s.Line, err = p.scanLocation()
	if err != nil {
		return nil, err
	}

	return s, nil

}

func (p *Parser) scanLocation() (location string, line int, err error) {
	var buf bytes.Buffer
	var tok Token
	var lit string
	tok, lit = p.scan()
	if tok != NewLine {
		err = fmt.Errorf("expected new line, found %v", tok.String())
		return
	}

	// scanning location
	for {
		tok, lit = p.scanWithoutSpaces()
		// wtf ?
		if tok == NewLine {
			err = fmt.Errorf("expected new line, found %v", tok.String())
			return
		}

		if tok == Colon {
			fmt.Errorf("expected Colon, found %v", tok.String())
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
