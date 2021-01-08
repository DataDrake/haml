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

package lexer

import (
	"unicode"
)

// lexFn are funtions used to Lex Tokens
type lexFn func(*Lexer) (f lexFn)

// Lexer generates a stream of Tokens from raw HAML text
type Lexer struct {
	out   chan Token
	data  []rune
	start int
	pos   int
	line  int
	col   int
}

// NewLexer creates a new Lexer for the specified string
func NewLexer(data string) *Lexer {
	return &Lexer{
		data: []rune(data),
		out:  make(chan Token),
		line: 1,
		col:  1,
	}
}

// isEOF checks for an EOF condition
func (l *Lexer) isEOF() bool {
	return l.pos >= len(l.data)
}

// advance to the next character
func (l *Lexer) advance() (eof bool) {
	l.pos++
	eof = l.isEOF()
	return
}

// peek at the next character
func (l *Lexer) peek() (r rune, eof bool) {
	if eof = l.isEOF(); eof {
		return
	}
	r = l.data[l.pos]
	return
}

// next character, advancing after
func (l *Lexer) next() (r rune, eof bool) {
	if l.isEOF() {
		return
	}
	r = l.data[l.pos]
	eof = l.advance()
	return
}

// skip over the next character, reseting token start position
func (l *Lexer) skip() (eof bool) {
	eof = l.advance()
	l.start = l.pos
	l.col++
	return
}

// skipSpaces until a newline or non-space character
func (l *Lexer) skipSpaces() (r rune, eof bool) {
	r, eof = l.peek()
	for !eof && unicode.IsSpace(r) && r != '\n' {
		if eof = l.skip(); eof {
			return
		}
		r, eof = l.peek()
	}
	return
}

// properName advances formware for a proper name
func (l *Lexer) properName() (ok, eof bool) {
	r, eof := l.peek()
	if eof {
		return
	}
	if !unicode.IsLetter(r) {
		return
	}
	for !eof && r != '\n' {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' {
			return
		}
		if eof = l.advance(); eof {
			return
		}
		r, eof = l.peek()
	}
	ok = true
	return
}

// go back a character if not at EOF
func (l *Lexer) backup() (eof bool) {
	if eof = l.isEOF(); eof {
		return
	}
	l.pos--
	return
}

// emit a new Token from the previously read characters
func (l *Lexer) emit(k TokenKind) {
	// Get the contents
	content := l.data[l.start:l.pos]
	// Make and send a new Token
	l.out <- Token{
		Kind:    k,
		Content: string(content),
		Line:    l.line,
		Col:     l.col,
	}
	// Update positions
	l.start = l.pos
	l.col += len(content)
	// Reset line and column counts on NewLine
	if k == NewLine {
		l.line++
		l.col = 1
	}
}

// Lex the raw HAML until EOF or an error
func (l *Lexer) Lex() {
	// start on a new line
	fn := lexIndent
	for fn != nil {
		fn = fn(l)
	}
	close(l.out)
}

// Next reads the next available Token
func (l *Lexer) Next() (t *Token, err error) {
	tok, ok := <-l.out
	if !ok {
		return
	}
	if tok.Kind == Error {
		err = tok.Error
		return
	}
	t = &tok
	return
}

// LexAll characters until EOF or an error
func (l *Lexer) LexAll() (tokens []Token, err error) {
	go l.Lex()
	var token *Token
	for {
		token, err = l.Next()
		if err != nil {
			return
		}
		tokens = append(tokens, *token)
	}
}
