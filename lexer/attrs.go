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
	"fmt"
)

// lexAttrsStart reads the start of an attributes list, skipping whitespace after
func lexAttrsStart(l *Lexer) lexFn {
	r, eof := l.next()
	if eof {
		return lexUnexpectedEOF(l)
	}
	if r != '{' {
		return lexError(l, fmt.Errorf("unexpected rune '%v'", r))
	}
	l.emit(AttrsStart)
	switch r {
	case '}':
		return lexAttrsEnd
	default:
		return lexAttrName
	}
}

// lexAttrName reads the name of an attribute, expecting a ':' to terminate it
func lexAttrName(l *Lexer) lexFn {
	if _, eof := l.skipSpaces(); eof {
		return lexUnexpectedEOF(l)
	}
	ok, eof := l.properName()
	if eof {
		return lexUnexpectedEOF(l)
	}
	if !ok {
		r, _ := l.peek()
		return lexUnexpectedRune(l, r)
	}
	l.emit(AttrName)
	return lexAttrAssign
}

// lexAttrAssign reads an attribute assignment, skipping whitespace after
func lexAttrAssign(l *Lexer) lexFn {
	r, eof := l.next()
	if eof {
		return lexUnexpectedEOF(l)
	}
	if r != ':' {
		return lexUnexpectedRune(l, r)
	}
	l.emit(AttrAssign)
	if _, eof = l.skipSpaces(); eof {
		return lexUnexpectedEOF(l)
	}
	return lexAttrValue
}

// lexAttrValue reads the vaue of an attribute, expecting it to be a double-quoted string
func lexAttrValue(l *Lexer) lexFn {
	r, eof := l.next()
	if eof {
		return lexUnexpectedEOF(l)
	}
	if r != '\'' {
		return lexUnexpectedRune(l, r)
	}
	r, eof = l.peek()
	for !eof && r != '\'' && r != '\n' {
		if eof = l.advance(); eof {
			return lexUnexpectedEOF(l)
		}
		r, eof = l.peek()
	}
	if r != '\'' {
		return lexUnexpectedRune(l, r)
	}
	l.emit(AttrValue)
	if r, eof = l.skipSpaces(); eof {
		return lexUnexpectedEOF(l)
	}
	switch r {
	case ',':
		return lexAttrSep
	case '}':
		return lexAttrsEnd
	default:
		return lexUnexpectedRune(l, r)
	}
}

// lexAttrSep reads the separator in a list of attributes
func lexAttrSep(l *Lexer) lexFn {
	r, eof := l.next()
	if eof {
		return lexUnexpectedEOF(l)
	}
	if r != ',' {
		return lexUnexpectedRune(l, r)
	}
	l.emit(AttrSep)
	if r, eof = l.skipSpaces(); eof {
		return lexUnexpectedEOF(l)
	}
	switch r {
	case '}':
		return lexAttrsEnd
	default:
		return lexAttrName
	}
}

// lexAttrsEnd reads the end of an attribute list, skipping any trailing space
func lexAttrsEnd(l *Lexer) lexFn {
	r, eof := l.next()
	if eof {
		return lexUnexpectedEOF(l)
	}
	if r != '}' {
		return lexUnexpectedRune(l, r)
	}
	l.emit(AttrsEnd)
	if _, eof = l.skipSpaces(); eof {
		return nil
	}
	return lexText
}
