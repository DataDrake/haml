//
// Copyright 2021 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package parser

import (
	"github.com/DataDrake/haml/lexer"
)

// parseFn are functions used to Parse Tags
type parseFn func(*Parser) parseFn

// Parser generates a stream of Tags from a stream of Tokens
type Parser struct {
	lex  *lexer.Lexer
	out  chan Tag
	curr *lexer.Token
	tag  *Tag
	done bool
	line int
}

// NewParser creates a new Parser for the specified string
func NewParser(data string) *Parser {
	return &Parser{
		lex:  lexer.NewLexer(data),
		out:  make(chan Tag),
		line: 1,
	}
}

// IsDone returns true if there are no more Tokens to read
func (p *Parser) IsDone() bool {
	return p.done
}

// advance to the next Token
func (p *Parser) advance() (done bool) {
	if p.IsDone() {
		done = true
		return
	}
	t, err := p.lex.Next()
	p.curr = t
	if err != nil {
		done = true
		p.done = done
		parseError(p, err)
		return
	}
	return
}

// peek at the next Token
func (p *Parser) peek() *lexer.Token {
	return p.curr
}

// next Token, advancing after
func (p *Parser) next() (t *lexer.Token) {
	t = p.curr
	p.advance()
	return
}

// skip over the next Token
func (p *Parser) skip() (done bool) {
	done = p.advance()
	return
}

// emit a new Tag from the previously read Tokens
func (p *Parser) emit() {
	p.out <- *p.tag
	p.tag = nil
}

// Parse the Tokens from the Lexer until EOF or an error
func (p *Parser) Parse() {
	// start on a new line
	fn := parseLine
	for fn != nil {
		fn = fn(p)
	}
	close(p.out)
}

// Next reads the next available Tag
func (p *Parser) Next() (t *Tag, err error) {
	tag, ok := <-p.out
	if !ok {
		return
	}
	if tag.Error != nil {
		err = tag.Error
		return
	}
	t = &tag
	return
}

// ParseAll Tags until EOF or an error
func (p *Parser) ParseAll() (tags []Tag, err error) {
	go p.Parse()
	var tag *Tag
	for {
		tag, err = p.Next()
		if err != nil {
			return
		}
		tags = append(tags, *tag)
	}
}
