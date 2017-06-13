// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aah-cb/minify source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package minify

import (
	"io"

	"aahframework.org/aah.v0"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

const htmlContentType = "text/html"

var m *minify.M

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Global methods
//___________________________________

// HTML method calls the minify library func to minify content. It reads data
// from reader and writes it to the given writer.
func HTML(contentType string, w io.Writer, r io.Reader) error {
	return m.Minify(htmlContentType, w, r)
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Unexported methods
//___________________________________

// addingHTMLMinifer event func to initilize the minify instance and
// adding minifier into framework
func addingHTMLMinifer(e *aah.Event) {
	m = minify.New()
	cfg := aah.AppConfig()

	m.Add(htmlContentType, &html.Minifier{
		KeepConditionalComments: cfg.BoolDefault("render.minify.keep.conditional_comments", true),
		KeepDocumentTags:        cfg.BoolDefault("render.minify.keep.document_tags", true),
		KeepWhitespace:          cfg.BoolDefault("render.minify.keep.whitespace", true),
		KeepDefaultAttrVals:     cfg.BoolDefault("render.minify.keep.default_attr_vals", false),
		KeepEndTags:             cfg.BoolDefault("render.minify.keep.end_tags", false),
	})

	// set `HTML` minify func.
	aah.SetMinifier(HTML)
}

func init() {
	// register into aah framework
	aah.OnInit(addingHTMLMinifer)
}
