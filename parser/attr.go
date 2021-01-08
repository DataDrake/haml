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
	"github.com/DataDrake/haml/lexer"
)

func parseAttr(p *Parser) parseFn {
	tok := p.peek()
	if tok == nil {
		return parseUnexpectedEOL(p)
	}
	if tok.Kind != lexer.AttrName {
		return parseUnexpectedToken(p, tok)
	}
	name, ok := parseProperName(p)
	if !ok {
		return nil
	}
	if !parseAttrAssign(p) {
		return nil
	}
	value, ok := parseAttrValue(p)
	if !ok {
		return nil
	}
	p.tag.Attrs[name] = value
	done := p.advance()
	if done {
		return parseUnexpectedEOL(p)
	}
	tok = p.peek()
	switch tok.Kind {
	case lexer.AttrSep:
		return parseAttrSeparator
	case lexer.AttrsEnd:
		return parseAttrsEnd
	default:
		return parseUnexpectedToken(p, tok)
	}
}

func parseAttrAssign(p *Parser) (ok bool) {
	tok := p.next()
	if tok == nil {
		parseUnexpectedEOL(p)
		return
	}
	if tok.Kind != lexer.AttrAssign {
		parseUnexpectedToken(p, tok)
		return
	}
	ok = true
	return
}

func parseAttrSeparator(p *Parser) parseFn {
	tok := p.next()
	if tok == nil {
		return parseUnexpectedEOL(p)
	}
	if tok.Kind != lexer.AttrSep {
		return parseUnexpectedToken(p, tok)
	}
	tok = p.peek()
	if tok == nil {
		return parseUnexpectedEOL(p)
	}
	switch tok.Kind {
	case lexer.AttrName:
		return parseAttr
	case lexer.AttrsEnd:
		return parseAttrsEnd
	default:
		return parseUnexpectedToken(p, tok)
	}
}

func parseAttrValue(p *Parser) (value string, ok bool) {
	tok := p.next()
	if tok == nil {
		parseUnexpectedEOL(p)
		return
	}
	if tok.Kind != lexer.AttrValue {
		parseUnexpectedToken(p, tok)
		return
	}
	if len(tok.Content) < 2 {
		parseErrorLocated(p, errors.New("attribute values must be quoted strings"), tok.Line, tok.Col)
		return
	}
	value = string(tok.Content)
	ok = true
	return
}
