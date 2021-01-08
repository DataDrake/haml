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

func parseID(p *Parser) parseFn {
	tok := p.peek()
	if tok.Kind != lexer.ID {
		return parseUnexpectedToken(p, tok)
	}
	name, ok := parseProperName(p)
	if !ok {
		return nil
	}
	if len(p.tag.Name) == 0 {
		p.tag.Name = "div"
	}
	p.tag.ID = name
	tok = p.peek()
	if tok == nil {
		return parseNewLine
	}
	switch tok.Kind {
	case lexer.Class:
		return parseClasses
	case lexer.AttrsStart:
		return parseAttrs
	default:
		return parseText
	}
}
