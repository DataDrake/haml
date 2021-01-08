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

// lexText reads the remainder of a line as text
func lexText(l *Lexer) lexFn {
	r, eof := l.peek()
	for !eof && r != '\n' {
		if eof = l.advance(); eof {
			break
		}
		r, eof = l.peek()
	}
	l.emit(Text)
	return lexNewLine
}
