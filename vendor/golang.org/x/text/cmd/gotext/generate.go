// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/text/message/pipeline"
	"golang.org/x/tools/go/loader"
=======
	"golang.org/x/text/message/pipeline"
>>>>>>> v0.0.4
)

func init() {
	out = cmdGenerate.Flag.String("out", "", "output file to write to")
}

<<<<<<< HEAD
var (
	out *string
)

=======
>>>>>>> v0.0.4
var cmdGenerate = &Command{
	Run:       runGenerate,
	UsageLine: "generate <package>",
	Short:     "generates code to insert translated messages",
}

<<<<<<< HEAD
var transRe = regexp.MustCompile(`messages\.(.*)\.json`)

func runGenerate(cmd *Command, args []string) error {

	prog, err := loadPackages(&loader.Config{}, args)
	if err != nil {
		return wrap(err, "could not load package")
	}

	pkgs := prog.InitialPackages()
	if len(pkgs) != 1 {
		return fmt.Errorf("more than one package selected: %v", pkgs)
	}
	pkg := pkgs[0].Pkg.Name()

	// TODO: add in external input. Right now we assume that all files are
	// manually created and stored in the textdata directory.

	// Build up index of translations and original messages.
	extracted := pipeline.Locale{}
	translations := []*pipeline.Locale{}

	err = filepath.Walk(*dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return wrap(err, "loading data")
		}
		if f.IsDir() {
			return nil
		}
		if f.Name() == extractFile {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return wrap(err, "read file failed")
			}
			if err := json.Unmarshal(b, &extracted); err != nil {
				return wrap(err, "unmarshal source failed")
			}
			return nil
		}
		if f.Name() == outFile {
			return nil
		}
		if !strings.HasSuffix(path, gotextSuffix) {
			return nil
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return wrap(err, "read file failed")
		}
		var locale pipeline.Locale
		if err := json.Unmarshal(b, &locale); err != nil {
			return wrap(err, "parsing translation file failed")
		}
		translations = append(translations, &locale)
		return nil
	})
	if err != nil {
		return err
	}

	w := os.Stdout
	if *out != "" {
		w, err = os.Create(*out)
		if err != nil {
			return wrap(err, "create file failed")
		}
	}

	_, err = pipeline.Generate(w, pkg, &extracted, translations...)
	return err
=======
func runGenerate(cmd *Command, config *pipeline.Config, args []string) error {
	config.Packages = args
	s, err := pipeline.Extract(config)
	if err != nil {
		return wrap(err, "extraction failed")
	}
	if err := s.Import(); err != nil {
		return wrap(err, "import failed")
	}
	return wrap(s.Generate(), "generation failed")
>>>>>>> v0.0.4
}
