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
	"errors"
	"fmt"
)

func lexError(l *Lexer, err error) lexFn {
	l.out <- Token{
		Kind:  Error,
		Error: fmt.Errorf("%s on line %d at col %d", err, l.line, l.col),
		Line:  l.line,
		Col:   l.col,
	}
	return nil
}

func lexUnexpectedEOF(l *Lexer) lexFn {
	return lexError(l, errors.New("input unexpectedly ended"))
}

func lexUnexpectedRune(l *Lexer, r rune) lexFn {
	return lexError(l, fmt.Errorf("unexpected rune '%v'", r))
}
