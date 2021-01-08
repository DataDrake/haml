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

func parseError(p *Parser, err error) parseFn {
	p.tag.Error = err
	p.emit()
	return nil
}

func parseErrorLocated(p *Parser, err error, line, col int) parseFn {
	err = fmt.Errorf("%s on line: %d, col %d", err, line, col)
	p.tag.Error = err
	p.emit()
	return nil
}

func parseUnexpectedEOL(p *Parser) parseFn {
	err := fmt.Errorf("input unexpectedly ended on line: %d", p.line)
	return parseError(p, err)
}

func parseUnexpectedToken(p *Parser, t *lexer.Token) parseFn {
	err := fmt.Errorf("unexpected token '%s'", t.Type())
	return parseErrorLocated(p, err, t.Line, t.Col)
}
