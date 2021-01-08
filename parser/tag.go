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

// Tag is the Parser's representation of an HTML tag
type Tag struct {
	Indent   int
	Name     string
	ID       string
	Attrs    map[string]string
	Content  string
	Children []Tag
	Error    error
}

func parseTag(p *Parser) parseFn {
	tok := p.peek()
	switch tok.Kind {
	case lexer.ID:
		return parseID
	case lexer.Element:
		return parseElement
	case lexer.Class:
		return parseClasses
	case lexer.AttrsStart:
		return parseAttrs
	default:
		return parseText
	}
}
