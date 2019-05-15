// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

<<<<<<< HEAD
//go:generate gotext extract --lang=de,zh
=======
//go:generate gotext -srclang=en update -out=catalog_gen.go -lang=en,zh
>>>>>>> v0.0.4

import (
	"net/http"

	"golang.org/x/text/cmd/gotext/examples/extract_http/pkg"
)

func main() {
	http.Handle("/generize", http.HandlerFunc(pkg.Generize))
}
