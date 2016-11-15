package parser

import (
	"io"
	"strconv"
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
		return nil
	}

	r := &Routine{}

	tok, lit = p.scanWithoutSpaces()
	if tok != Integer {
		return nil
	}

	// we already know its an integer
	// scan the routine ID
	r.ID, _ = strconv.Atoi(lit)

	return r

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
