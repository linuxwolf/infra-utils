package pkg

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	ErrParseEmptyLine   = fmt.Errorf("empty line")
	ErrParseCommentLine = fmt.Errorf("comment line")
	ErrParseInvalid     = fmt.Errorf("invalid line")
)

type Parser struct {
	ptn *regexp.Regexp
	log *log.Logger
}

func NewParser(l *log.Logger) *Parser {
	if l == nil {
		l = log.Default()
	}

	return &Parser{
		ptn: regexp.MustCompile(`^([_a-zA-Z][_a-zA-Z0-9]*)\s*=\s*(.*)\s*$`),
		log: l,
	}
}

func (p *Parser) ProcessReader(r io.Reader) *Env {
	scanner := bufio.NewScanner(r)
	env := NewEnvWith(nil)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		k, v, err := p.parseLine(line)
		if err != nil {
			p.log.Printf("skipping: %s", err)
		} else {
			env.vars[k] = v
		}
	}
	if err := scanner.Err(); err != nil {
		p.log.Printf("error while parsing reader: %v", err)
	}

	return env
}

func (p *Parser) ProcessArray(ins []string) *Env {
	env := NewEnvWith(nil)
	for _, line := range ins {
		line = strings.TrimSpace(line)
		k, v, err := p.parseLine(line)
		if err != nil {
			p.log.Printf("skipping: %s", err)
		} else {
			env.vars[k] = v
		}
	}

	return env
}

func (p *Parser) parseLine(line string) (string, string, error) {
	line = strings.TrimSpace(line)

	// skip comment lines
	if line == "" {
		return "", "", ErrParseEmptyLine
	}
	if strings.HasPrefix(line, "#") {
		return "", "", ErrParseCommentLine
	}
	// parse into key and value, skipping (with warning) if invalid
	parts := p.ptn.FindStringSubmatch(line)
	if parts == nil || len(parts) < 3 {
		return "", "", ErrParseInvalid
	}

	return parts[1], parts[2], nil
}

func NewEnvsFromEnviron() *Env {
	parser := NewParser(nil)
	env := parser.ProcessArray(os.Environ())

	return env
}
