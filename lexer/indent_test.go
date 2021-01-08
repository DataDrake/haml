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
	"testing"
)

func TestIndent2SP(t *testing.T) {
	l := NewLexer("  ")
	go l.Lex()
	tok, err := l.Next()
	if err != nil {
		t.Fatalf("Should be no error, found: %s", err)
	}
	if tok == nil {
		t.Fatal("token should not be nil")
	}
	if tok.Kind != Indent {
		t.Errorf("expected '%d', found '%d'", Indent, tok.Kind)
	}
	if tok.Content != "  " {
		t.Errorf("expected '%s', found '%s'", "  ", tok.Content)
	}
	if tok.Line != 1 {
		t.Errorf("expected '%d', found '%d'", 1, tok.Line)
	}
	if tok.Col != 1 {
		t.Errorf("expected '%d', found '%d'", 1, tok.Col)
	}
	tok, err = l.Next()
	if err != nil {
		t.Fatalf("Should be no error, found: %s", err)
	}
	if tok != nil {
		t.Fatal("Should be the last token")
	}
}

func TestIndent3T(t *testing.T) {
	l := NewLexer("\t\t\t")
	go l.Lex()
	tok, err := l.Next()
	if err != nil {
		t.Fatalf("Should be no error, found: %s", err)
	}
	if tok == nil {
		t.Fatal("token should not be nil")
	}
	if tok.Kind != Indent {
		t.Errorf("expected '%d', found '%d'", Indent, tok.Kind)
	}
	if tok.Content != "\t\t\t" {
		t.Errorf("expected '%s', found '%s'", "\t\t\t", tok.Content)
	}
	if tok.Line != 1 {
		t.Errorf("expected '%d', found '%d'", 1, tok.Line)
	}
	if tok.Col != 1 {
		t.Errorf("expected '%d', found '%d'", 1, tok.Col)
	}
	tok, err = l.Next()
	if err != nil {
		t.Fatalf("Should be no error, found: %s", err)
	}
	if tok != nil {
		t.Fatal("Should be the last token")
	}
}

func TestIndent2SP3T(t *testing.T) {
	l := NewLexer("  \t\t\t")
	go l.Lex()
	tok, err := l.Next()
	if err != nil {
		t.Fatalf("Should be no error, found: %s", err)
	}
	if tok == nil {
		t.Fatal("token should not be nil")
	}
	if tok.Kind != Indent {
		t.Errorf("expected '%d', found '%d'", Indent, tok.Kind)
	}
	if tok.Content != "  \t\t\t" {
		t.Errorf("expected '%s', found '%s'", "  \t\t\t", tok.Content)
	}
	if tok.Line != 1 {
		t.Errorf("expected '%d', found '%d'", 1, tok.Line)
	}
	if tok.Col != 1 {
		t.Errorf("expected '%d', found '%d'", 1, tok.Col)
	}
	tok, err = l.Next()
	if err != nil {
		t.Fatalf("Should be no error, found: %s", err)
	}
	if tok != nil {
		t.Fatal("Should be the last token")
	}
}
