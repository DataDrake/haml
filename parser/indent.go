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
	"fmt"
	"github.com/DataDrake/haml/lexer"
)

func parseIndent(p *Parser) parseFn {
	tok := p.next()
	if tok == nil {
		return parseNewLine
	}
	if tok.Kind != lexer.Indent {
		return parseUnexpectedToken(p, tok)
	}
	if len(tok.Content) == 0 {
		return parseTag
	}
	space := false
	for i, r := range tok.Content {
		if space {
			if r != ' ' {
				return parseErrorLocated(p, fmt.Errorf("unexpected rune '%c'", r), tok.Line, tok.Col+i)
			}
			p.tag.Indent++
			space = false
			continue
		}
		switch r {
		case ' ':
			space = true
		case '\t':
			p.tag.Indent++
		default:
			if r != ' ' {
				return parseErrorLocated(p, fmt.Errorf("unexpected rune '%c'", r), tok.Line, tok.Col+i)
			}
		}
	}
	return parseTag
}
