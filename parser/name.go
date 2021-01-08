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
	"errors"
	"fmt"
	"unicode"
)

func parseProperName(p *Parser) (name string, ok bool) {
	tok := p.next()
	if len(tok.Content) == 0 {
		parseErrorLocated(p, errors.New("missing name"), tok.Line, tok.Col)
		return
	}
	if r := []rune(tok.Content)[0]; !unicode.IsLetter(r) {
		err := fmt.Errorf("proper names must start with a letter, found: '%c'", r)
		parseErrorLocated(p, err, tok.Line, tok.Col)
		return
	}
	for i, r := range tok.Content {
		if i == 0 {
			continue
		}
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '-' {
			err := fmt.Errorf("proper names must be letters, numbers, or hyphens, found: %c", r)
			parseErrorLocated(p, err, tok.Line, tok.Col+i)
			return
		}
	}
	name = string(tok.Content)
	ok = true
	return
}
