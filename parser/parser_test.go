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
//"testing"
)

/*
func TestParseBreak(t *testing.T) {
	haml := "\n"
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_BREAK {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_BREAK, l.Kind)
	}
}

func TestParseText(t *testing.T) {
	haml := []rune("definitely not anything other than text\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TEXT {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TEXT, l.Kind)
	}
	if l.Content != "definitely not anything other than text" {
		t.Errorf("Expected content '%s', found ID: '%s'", "definitely not anything other than text", l.Content)
	}
}

func TestParseTemplate(t *testing.T) {
	haml := []rune("{{go template here}}\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TEMPLATE {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TEMPLATE, l.Kind)
	}
	if l.Content != "{{go template here}}" {
		t.Errorf("Expected content '%s', found ID: '%s'", "{{go template here}}", l.Content)
	}
}

func TestParseIDTag(t *testing.T) {
	haml := []rune("#validID\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TAG {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TAG, l.Kind)
	}
	if l.Tag.ID != "validID" {
		t.Errorf("Expected ID '%s', found ID: '%s'", "validID", l.Tag.ID)
	}
	if l.Tag.Name != "div" {
		t.Errorf("Expected name '%s', found name: '%s'", "div", l.Tag.Name)
	}
}

func TestParseIndentedTag(t *testing.T) {
	haml := []rune("  #validID\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 2 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TAG {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TAG, l.Kind)
	}
	if l.Tag.ID != "validID" {
		t.Errorf("Expected ID '%s', found ID: '%s'", "validID", l.Tag.ID)
	}
	if l.Tag.Name != "div" {
		t.Errorf("Expected name '%s', found name: '%s'", "div", l.Tag.Name)
	}
}

func TestParseIDAttrsTag(t *testing.T) {
	haml := []rune("#validID(   name: \"this is a value=2\"   )\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TAG {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TAG, l.Kind)
	}
	if l.Tag.ID != "validID" {
		t.Errorf("Expected ID '%s', found ID: '%s'", "validID", l.Tag.ID)
	}
	if l.Tag.Name != "div" {
		t.Errorf("Expected name '%s', found name: '%s'", "div", l.Tag.Name)
	}
	if l.Tag.Attrs["name"] != "this is a value=2" {
		t.Errorf("Expected attr 'name' value '%s', found value: '%s'", "this is a value=2", l.Tag.Attrs["name"])
	}
}

func TestParseIDMultipleAttrsTag(t *testing.T) {
	haml := []rune("#validID(name1:\"value1\", name2: \"value2\"  )\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TAG {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TAG, l.Kind)
	}
	if l.Tag.ID != "validID" {
		t.Errorf("Expected ID '%s', found ID: '%s'", "validID", l.Tag.ID)
	}
	if l.Tag.Name != "div" {
		t.Errorf("Expected name '%s', found name: '%s'", "div", l.Tag.Name)
	}
	if l.Tag.Attrs["name1"] != "value1" {
		t.Errorf("Expected attr '%s' value '%s', found value: '%s'", "name1", "value1", l.Tag.Attrs["name1"])
	}
	if l.Tag.Attrs["name2"] != "value2" {
		t.Errorf("Expected attr '%s' value '%s', found value: '%s'", "name2", "value2", l.Tag.Attrs["name2"])
	}
}

func TestParseElementTag(t *testing.T) {
	haml := []rune("%element-name\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TAG {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TAG, l.Kind)
	}
	if l.Tag.Name != "element-name" {
		t.Errorf("Expected ID '%s', found ID: '%s'", "element-name", l.Tag.ID)
	}
}

func TestParseClassTag(t *testing.T) {
	haml := []rune(".class-name\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TAG {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TAG, l.Kind)
	}
	if l.Tag.Name != "div" {
		t.Errorf("Expected name '%s', found name: '%s'", "div", l.Tag.Name)
	}
	if l.Tag.Attrs["class"] != "class-name" {
		t.Errorf("Expected class name '%s', found class name: '%s'", "class-name", l.Tag.Attrs["class"])
	}
}

func TestParseClassedIDTag(t *testing.T) {
	haml := []rune("#validID.class-name-1.class-name-2\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TAG {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TAG, l.Kind)
	}
	if l.Tag.ID != "validID" {
		t.Errorf("Expected ID '%s', found ID: '%s'", "validID", l.Tag.ID)
	}
	if l.Tag.Name != "div" {
		t.Errorf("Expected name '%s', found name: '%s'", "div", l.Tag.Name)
	}
	if l.Tag.Attrs["class"] != "class-name-1 class-name-2" {
		t.Errorf("Expected class names '%s', found class name: '%s'", "class-name-1 class-name-2", l.Tag.Attrs["class"])
	}
}

func TestParseClassedElementTag(t *testing.T) {
	haml := []rune("%element-name.class-name-1.class-name-2\n")
	l, i, err := parseLine(haml, 0)
	if err != nil {
		t.Errorf("Expected no error, found: '%s'", err.Error())
	}
	if i != len(haml) {
		t.Errorf("Expected index '%d', found index '%d'", len(haml), i)
	}
	if l.Indent != 0 {
		t.Errorf("Expected indent '%d', found indent: '%d'", 0, l.Indent)
	}
	if l.Kind != LINE_TAG {
		t.Errorf("Expected line type '%d', found type: '%d'", LINE_TAG, l.Kind)
	}
	if l.Tag.ID != "" {
		t.Errorf("Expected ID '%s', found ID: '%s'", "", l.Tag.ID)
	}
	if l.Tag.Name != "element-name" {
		t.Errorf("Expected name '%s', found name: '%s'", "element-name", l.Tag.Name)
	}
	if l.Tag.Attrs["class"] != "class-name-1 class-name-2" {
		t.Errorf("Expected class names '%s', found class name: '%s'", "class-name-1 class-name-2", l.Tag.Attrs["class"])
	}
}

const TestHaml = `
%section.container
  %h1= post.title
  %h2= post.subtitle
  .content
    = post.content
  #id(attr:"bob")
`

func testParseLine(t *testing.T) {

	text := []rune(TestHaml)
	l, i, err := parseLine(text, 0)
	t.Errorf("Line: %#v\n", l)
	t.Errorf("i: %d\n", i)
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
	}
	l, i, err = parseLine(text, i)
	t.Errorf("Line: %#v\n", l)
	t.Errorf("i: %d\n", i)
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
	}
	l, i, err = parseLine(text, i)
	t.Errorf("Line: %#v\n", l)
	t.Errorf("i: %d\n", i)
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
	}
	l, i, err = parseLine(text, i)
	t.Errorf("Line: %#v\n", l)
	t.Errorf("i: %d\n", i)
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
	}
	l, i, err = parseLine(text, i)
	t.Errorf("Line: %#v\n", l)
	t.Errorf("i: %d\n", i)
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
	}
	l, i, err = parseLine(text, i)
	t.Errorf("Line: %#v\n", l)
	t.Errorf("i: %d\n", i)
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
	}
	l, i, err = parseLine(text, i)
	t.Errorf("Line: %#v\n", l)
	t.Errorf("i: %d\n", i)
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
	}
	l, i, err = parseLine(text, i)
	t.Errorf("Line: %#v\n", l)
	t.Errorf("i: %d\n", i)
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
	}
}
*/
