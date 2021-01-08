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

// lexID reads an ID tag
func lexID(l *Lexer) lexFn {
	return lexTag(l, '#', ID)
}

// lexElement reads an Element tag
func lexElement(l *Lexer) lexFn {
	return lexTag(l, '%', Element)
}

// lexClass reads a Class tag
func lexClass(l *Lexer) lexFn {
	return lexTag(l, '.', Class)
}

// lexTag reads a tag, pivoting to further tags
func lexTag(l *Lexer, start rune, kind TokenKind) lexFn {
	r, eof := l.next()
	if eof {
		return lexUnexpectedEOF(l)
	}
	if r != start {
		return lexUnexpectedRune(l, r)
	}
	ok, eof := l.properName()
	if eof {
		return lexUnexpectedEOF(l)
	}
	if !ok {
		r, _ := l.peek()
		return lexUnexpectedRune(l, r)
	}
	l.emit(kind)
	r, eof = l.peek()
	if eof {
		return nil
	}
	if kind == '%' && r == '#' {
		return lexID
	}
	switch r {
	case '.':
		return lexClass
	case '(':
		return lexAttrsStart
	default:
		return lexText
	}
}
