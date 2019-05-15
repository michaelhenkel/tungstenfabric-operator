// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
<<<<<<< HEAD
	"bytes"
=======
>>>>>>> v0.0.4
	"context"
	"encoding/json"
	"flag"
	"fmt"
<<<<<<< HEAD
	"go/types"
	"os"

	guru "golang.org/x/tools/cmd/guru/serial"
	"golang.org/x/tools/internal/lsp/cache"
	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/lsp/xlog"
=======
	"os"
	"strings"

	guru "golang.org/x/tools/cmd/guru/serial"
	"golang.org/x/tools/internal/lsp/protocol"
>>>>>>> v0.0.4
	"golang.org/x/tools/internal/span"
	"golang.org/x/tools/internal/tool"
)

// A Definition is the result of a 'definition' query.
type Definition struct {
	Span        span.Span `json:"span"`        // span of the definition
	Description string    `json:"description"` // description of the denoted object
}

// These constant is printed in the help, and then used in a test to verify the
// help is still valid.
// They refer to "Set" in "flag.FlagSet" from the DetailedHelp method below.
const (
<<<<<<< HEAD
	exampleLine   = 47
	exampleColumn = 47
	exampleOffset = 1359
=======
	exampleLine   = 44
	exampleColumn = 47
	exampleOffset = 1270
>>>>>>> v0.0.4
)

// definition implements the definition noun for the query command.
type definition struct {
	query *query
}

func (d *definition) Name() string      { return "definition" }
func (d *definition) Usage() string     { return "<position>" }
func (d *definition) ShortHelp() string { return "show declaration of selected identifier" }
func (d *definition) DetailedHelp(f *flag.FlagSet) {
	fmt.Fprintf(f.Output(), `
Example: show the definition of the identifier at syntax at offset %[1]v in this file (flag.FlagSet):

$ gopls definition internal/lsp/cmd/definition.go:%[1]v:%[2]v
$ gopls definition internal/lsp/cmd/definition.go:#%[3]v

	gopls query definition flags are:
`, exampleLine, exampleColumn, exampleOffset)
	f.PrintDefaults()
}

// Run performs the definition query as specified by args and prints the
// results to stdout.
func (d *definition) Run(ctx context.Context, args ...string) error {
	if len(args) != 1 {
		return tool.CommandLineErrorf("definition expects 1 argument")
	}
<<<<<<< HEAD
	log := xlog.New(xlog.StdSink{})
	view := cache.NewView(ctx, log, "definition_test", span.FileURI(d.query.app.Config.Dir), &d.query.app.Config)
	from := span.Parse(args[0])
	f, err := view.GetFile(ctx, from.URI())
	if err != nil {
		return err
	}
	converter := span.NewTokenConverter(view.FileSet(), f.GetToken(ctx))
	rng, err := from.Range(converter)
	if err != nil {
		return err
	}
	ident, err := source.Identifier(ctx, view, f, rng.Start)
	if err != nil {
		return fmt.Errorf("%v: %v", from, err)
	}
	if ident == nil {
		return fmt.Errorf("%v: not an identifier", from)
	}
	var result interface{}
	switch d.query.Emulate {
	case "":
		result, err = buildDefinition(ctx, view, ident)
	case emulateGuru:
		result, err = buildGuruDefinition(ctx, view, ident)
=======
	client := &baseClient{}
	server, err := d.query.app.connect(ctx, client)
	if err != nil {
		return err
	}
	from := span.Parse(args[0])
	m, err := client.AddFile(ctx, from.URI())
	if err != nil {
		return err
	}
	loc, err := m.Location(from)
	if err != nil {
		return err
	}
	p := protocol.TextDocumentPositionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: loc.URI},
		Position:     loc.Range.Start,
	}
	locs, err := server.Definition(ctx, &p)
	if err != nil {
		return fmt.Errorf("%v: %v", from, err)
	}

	if len(locs) == 0 {
		return fmt.Errorf("%v: not an identifier", from)
	}
	hover, err := server.Hover(ctx, &p)
	if err != nil {
		return fmt.Errorf("%v: %v", from, err)
	}
	if hover == nil {
		return fmt.Errorf("%v: not an identifier", from)
	}
	m, err = client.AddFile(ctx, span.NewURI(locs[0].URI))
	if err != nil {
		return fmt.Errorf("%v: %v", from, err)
	}
	definition, err := m.Span(locs[0])
	if err != nil {
		return fmt.Errorf("%v: %v", from, err)
	}
	description := strings.TrimSpace(hover.Contents.Value)
	var result interface{}
	switch d.query.Emulate {
	case "":
		result = &Definition{
			Span:        definition,
			Description: description,
		}
	case emulateGuru:
		pos := span.New(definition.URI(), definition.Start(), definition.Start())
		result = &guru.Definition{
			ObjPos: fmt.Sprint(pos),
			Desc:   description,
		}
>>>>>>> v0.0.4
	default:
		return fmt.Errorf("unknown emulation for definition: %s", d.query.Emulate)
	}
	if err != nil {
		return err
	}
	if d.query.JSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "\t")
		return enc.Encode(result)
	}
	switch d := result.(type) {
	case *Definition:
		fmt.Printf("%v: defined here as %s", d.Span, d.Description)
	case *guru.Definition:
		fmt.Printf("%s: defined here as %s", d.ObjPos, d.Desc)
	default:
		return fmt.Errorf("no printer for type %T", result)
	}
	return nil
}
<<<<<<< HEAD

func buildDefinition(ctx context.Context, view source.View, ident *source.IdentifierInfo) (*Definition, error) {
	content, err := ident.Hover(ctx, nil)
	if err != nil {
		return nil, err
	}
	spn, err := ident.Declaration.Range.Span()
	if err != nil {
		return nil, err
	}
	return &Definition{
		Span:        spn,
		Description: content,
	}, nil
}

func buildGuruDefinition(ctx context.Context, view source.View, ident *source.IdentifierInfo) (*guru.Definition, error) {
	spn, err := ident.Declaration.Range.Span()
	if err != nil {
		return nil, err
	}
	pkg := ident.File.GetPackage(ctx)
	// guru does not support ranges
	if !spn.IsPoint() {
		spn = span.New(spn.URI(), spn.Start(), spn.Start())
	}
	// Behavior that attempts to match the expected output for guru. For an example
	// of the format, see the associated definition tests.
	buf := &bytes.Buffer{}
	q := types.RelativeTo(pkg.GetTypes())
	qualifyName := ident.Declaration.Object.Pkg() != pkg.GetTypes()
	name := ident.Name
	var suffix interface{}
	switch obj := ident.Declaration.Object.(type) {
	case *types.TypeName:
		fmt.Fprint(buf, "type")
	case *types.Var:
		if obj.IsField() {
			qualifyName = false
			fmt.Fprint(buf, "field")
			suffix = obj.Type()
		} else {
			fmt.Fprint(buf, "var")
		}
	case *types.Func:
		fmt.Fprint(buf, "func")
		typ := obj.Type()
		if obj.Type() != nil {
			if sig, ok := typ.(*types.Signature); ok {
				buf := &bytes.Buffer{}
				if recv := sig.Recv(); recv != nil {
					if named, ok := recv.Type().(*types.Named); ok {
						fmt.Fprintf(buf, "(%s).%s", named.Obj().Name(), name)
					}
				}
				if buf.Len() == 0 {
					buf.WriteString(name)
				}
				types.WriteSignature(buf, sig, q)
				name = buf.String()
			}
		}
	default:
		fmt.Fprintf(buf, "unknown [%T]", obj)
	}
	fmt.Fprint(buf, " ")
	if qualifyName {
		fmt.Fprintf(buf, "%s.", ident.Declaration.Object.Pkg().Path())
	}
	fmt.Fprint(buf, name)
	if suffix != nil {
		fmt.Fprint(buf, " ")
		fmt.Fprint(buf, suffix)
	}
	return &guru.Definition{
		ObjPos: fmt.Sprint(spn),
		Desc:   buf.String(),
	}, nil
}
=======
>>>>>>> v0.0.4
