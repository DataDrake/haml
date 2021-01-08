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

// lexIndent reads space or tab characters at the start of a line
func lexIndent(l *Lexer) lexFn {
	r, eof := l.peek()
	for !eof && (r == ' ' || r == '\t') {
		if eof = l.advance(); eof {
			break
		}
		r, eof = l.peek()
	}
	if l.pos != l.start {
		l.emit(Indent)
	}
	if eof {
		return nil
	}
	switch r {
	case '#':
		return lexID
	case '%':
		return lexElement
	case '.':
		return lexClass
	case '{':
		return lexAttrsStart
	default:
		return lexText
	}
}
