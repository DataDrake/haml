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

// TokenKind specifies the type of a Token
type TokenKind int

const (
	// NewLine are found at the end of each textual line
	NewLine TokenKind = iota
	// Indent contain all of the leading whitespace of a line
	Indent
	// Text contains all of the trailing text on a line
	Text
	// ID specifies the HTML tag ID for the current Div tag
	ID
	// Class specifies an HTML class for the current tag
	Class
	// Element specifies kind of tag for this line
	Element
	// AttrsStart specifies the beginning of an attribute list
	AttrsStart
	// AttrName is the name of an attribute
	AttrName
	// AttrAssign is the assignment operator for Name:Value pairs of attributes
	AttrAssign
	// AttrValue is the value assigned to an attribute, as a quoted string
	AttrValue
	// AttrSep is the list separator for attributes
	AttrSep
	// AttrsEnd marks the end of an attribute list
	AttrsEnd
	// Error is a special Token used to communicate errors to the Parser
	Error
)

var tokenKindNames = map[TokenKind]string{
	NewLine:    "Newline",
	Indent:     "Indent",
	Text:       "Text",
	ID:         "ID",
	Class:      "Class",
	Element:    "Element",
	AttrsStart: "Start Attributes",
	AttrName:   "Attribute Name",
	AttrAssign: "Attribute Assignment",
	AttrValue:  "Attribute Value",
	AttrSep:    "Attribute Separator",
	AttrsEnd:   "End Attributes",
	Error:      "Error",
}

// Token is a lexer Token
type Token struct {
	Kind    TokenKind
	Content string
	Line    int
	Col     int
	Error   error
}

// Type returns the name of the Kind of this Token
func (t Token) Type() string {
	return tokenKindNames[t.Kind]
}
