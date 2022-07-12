package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type instructionCode int

const (
	aInstruction instructionCode = iota
	cInstruction
	lInstruction
)

type Parser struct {
	reader  *bufio.Reader
	source  string
	current string
	dest    string
	comp    string
	jump    string
}

func newParser(file *os.File) Parser {
	reader := bufio.NewReader(file)

	return Parser{reader, "", "", "", "", ""}
}

func (p *Parser) hasMoreLines() bool {
	_, err := p.reader.Peek(1)
	if err != nil {
		return false
	}
	return true
}

func (p *Parser) advance() {
	b, _, err := p.reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	p.current = string(b)
	if p.current == "" {
		p.advance()
		return
	}

	if len(p.current) <= 0 {
		return
	}
	tokens := strings.SplitN(p.current, "//", 2)
	if len(tokens) > 0 {
		p.current = tokens[0]
	}
	p.current = strings.TrimSpace(p.current)
}

func (p *Parser) instructionType() instructionCode {
	switch p.current[0] {
	case '@':
		return aInstruction // assignment
	case '(':
		return lInstruction // label
	default:
		tokens := strings.Split(p.current, ";")
		if len(tokens) == 2 {
			p.jump = tokens[1]
		} else {
			p.jump = ""
		}

		comp := tokens[0]
		tokens = strings.Split(tokens[0], "=")
		if len(tokens) == 2 {
			p.dest = tokens[0]
			p.comp = tokens[1]
		} else {
			p.dest = ""
			p.comp = comp
		}

		return cInstruction
	}
}

func (p *Parser) symbol() string {
	return strings.TrimPrefix(p.current, "@")
}
