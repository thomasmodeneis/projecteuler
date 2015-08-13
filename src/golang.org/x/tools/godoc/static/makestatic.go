// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// Command bake reads a set of files and writes a Go source file to "static.go"
// that declares a map of string constants containing contents of the input files.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

var files = []string{
	"analysis/call3.png",
	"analysis/call-eg.png",
	"analysis/callers1.png",
	"analysis/callers2.png",
	"analysis/chan1.png",
	"analysis/chan2a.png",
	"analysis/chan2b.png",
	"analysis/error1.png",
	"analysis/help.html",
	"analysis/ident-def.png",
	"analysis/ident-field.png",
	"analysis/ident-func.png",
	"analysis/ipcg-func.png",
	"analysis/ipcg-pkg.png",
	"analysis/typeinfo-pkg.png",
	"analysis/typeinfo-src.png",
	"callgraph.html",
	"codewalk.html",
	"codewalkdir.html",
	"dirlist.html",
	"error.html",
	"example.html",
	"godoc.html",
	"godocs.js",
	"images/minus.gif",
	"images/plus.gif",
	"images/treeview-black-line.gif",
	"images/treeview-black.gif",
	"images/treeview-default-line.gif",
	"images/treeview-default.gif",
	"images/treeview-gray-line.gif",
	"images/treeview-gray.gif",
	"implements.html",
	"jquery.js",
	"jquery.treeview.css",
	"jquery.treeview.edit.js",
	"jquery.treeview.js",
	"methodset.html",
	"opensearch.xml",
	"package.html",
	"package.txt",
	"play.js",
	"playground.js",
	"search.html",
	"search.txt",
	"searchcode.html",
	"searchdoc.html",
	"searchtxt.html",
	"style.css",
}

func main() {
	if err := bake(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func bake() error {
	f, err := os.Create("static.go")
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "%v\n\npackage static\n\n", warning)
	fmt.Fprintf(w, "var Files = map[string]string{\n")
	for _, fn := range files {
		b, err := ioutil.ReadFile(fn)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "\t%q: ", fn)
		if utf8.Valid(b) {
			fmt.Fprintf(w, "`%s`", sanitize(b))
		} else {
			fmt.Fprintf(w, "%q", b)
		}
		fmt.Fprintln(w, ",\n")
	}
	fmt.Fprintln(w, "}")
	if err := w.Flush(); err != nil {
		return err
	}
	return f.Close()
}

// sanitize prepares a valid UTF-8 string as a raw string constant.
func sanitize(b []byte) []byte {
	// Replace ` with `+"`"+`
	b = bytes.Replace(b, []byte("`"), []byte("`+\"`\"+`"), -1)

	// Replace BOM with `+"\xEF\xBB\xBF"+`
	// (A BOM is valid UTF-8 but not permitted in Go source files.
	// I wouldn't bother handling this, but for some insane reason
	// jquery.js has a BOM somewhere in the middle.)
	return bytes.Replace(b, []byte("\xEF\xBB\xBF"), []byte("`+\"\\xEF\\xBB\\xBF\"+`"), -1)
}

const warning = "// DO NOT EDIT ** This file was generated with the bake tool ** DO NOT EDIT //"
