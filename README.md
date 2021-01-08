# haml
HAML Parser in pure Go

[![Go Report Card](https://goreportcard.com/badge/github.com/DataDrake/haml)](https://goreportcard.com/report/github.com/DataDrake/haml) [![license](https://img.shields.io/github/license/DataDrake/haml.svg)]()

## Motivation

All of the existing Go implementations of HAML have gone by the wayside. This makes sense given that HAML started in Ruby and there are many other templating and markup languages that have popped up since. HAML still works well for me and my purposes and it looks like I'll need to step up and build a parser for myself.

## Goals

 * Support parsing HAML formatted documents
 * Adapt HAML to use the Go template language
 * Generate well-formed HTML5 documents
 * Provide a simple CLI tool for working with HAML files and templates
 * A+ Rating on [Report Card](https://goreportcard.com/report/github.com/DataDrake/sup)
 
## Limitations

This library is not intended to be 100% compatible with Ruby implementations of HAML. The following features are planned for support:

- [ ] Tags
	- [ ] `#id-of-tag`
	- [ ] `%name-of-element`
	- [ ] `.class.lists`
	- [ ] `{attributes: "one", two:"three"}` (Ruby 1.9)
	- [ ] Auto-closing (e.g. br, hr, img)
- [ ] Text
	- [ ] Multiline strings
		- [ ] nbsp if missing whitespace
		- [ ] no preservation of leading whitespace

Because the Go templating language is very capable and this library will ultimately rely on `html/template`, many features are either not needed or complicate parsing. The following features are intentionally omitted:

- Embedded Ruby code
- Ruby hash-style attribute lists
- HTML-style attribute lists
- Prefixed attributes
- Filters
- Helper methods
- HAML Comments
- HTML Comments
- Doctypes
- Explicit multi-line strings
- Whitespace preservation
- Interpolated strings

## Usage

```
WORK IN PROGRESS
```

## License
 
Copyright 2021 Bryan T. Meyers <root@datadrake.com>
 
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 
http://www.apache.org/licenses/LICENSE-2.0
 
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
