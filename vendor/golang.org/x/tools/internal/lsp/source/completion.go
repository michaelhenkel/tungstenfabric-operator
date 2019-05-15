<<<<<<< HEAD
package source

import (
	"bytes"
=======
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

import (
>>>>>>> v0.0.4
	"context"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
<<<<<<< HEAD
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

type CompletionItem struct {
	Label, Detail string
	Kind          CompletionItemKind
	Score         float64
=======

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/internal/lsp/snippet"
)

type CompletionItem struct {
	// Label is the primary text the user sees for this completion item.
	Label string

	// Detail is supplemental information to present to the user.
	// This often contains the type or return type of the completion item.
	Detail string

	// InsertText is the text to insert if this item is selected.
	// Any of the prefix that has already been typed is not trimmed.
	// The insert text does not contain snippets.
	InsertText string

	Kind CompletionItemKind

	// Score is the internal relevance score.
	// A higher score indicates that this completion item is more relevant.
	Score float64

	// Snippet is the LSP snippet for the completion item, without placeholders.
	// The LSP specification contains details about LSP snippets.
	// For example, a snippet for a function with the following signature:
	//
	//     func foo(a, b, c int)
	//
	// would be:
	//
	//     foo(${1:})
	//
	Snippet *snippet.Builder

	// PlaceholderSnippet is the LSP snippet for the completion ite, containing
	// placeholders. The LSP specification contains details about LSP snippets.
	// For example, a placeholder snippet for a function with the following signature:
	//
	//     func foo(a, b, c int)
	//
	// would be:
	//
	//     foo(${1:a int}, ${2: b int}, ${3: c int})
	//
	PlaceholderSnippet *snippet.Builder
>>>>>>> v0.0.4
}

type CompletionItemKind int

const (
	Unknown CompletionItemKind = iota
	InterfaceCompletionItem
	StructCompletionItem
	TypeCompletionItem
	ConstantCompletionItem
	FieldCompletionItem
	ParameterCompletionItem
	VariableCompletionItem
	FunctionCompletionItem
	MethodCompletionItem
	PackageCompletionItem
)

<<<<<<< HEAD
// stdScore is the base score value set for all completion items.
const stdScore float64 = 1.0

// finder is a function used to record a completion candidate item in a list of
// completion items.
type finder func(types.Object, float64, []CompletionItem) []CompletionItem

// Completion returns a list of possible candidates for completion, given a
// a file and a position. The prefix is computed based on the preceding
// identifier and can be used by the client to score the quality of the
// completion. For instance, some clients may tolerate imperfect matches as
// valid completion results, since users may make typos.
func Completion(ctx context.Context, f File, pos token.Pos) (items []CompletionItem, prefix string, err error) {
	file := f.GetAST(ctx)
	pkg := f.GetPackage(ctx)
	if pkg.IsIllTyped() {
		return nil, "", fmt.Errorf("package for %s is ill typed", f.URI())
	}
	path, _ := astutil.PathEnclosingInterval(file, pos, pos)
	if path == nil {
		return nil, "", fmt.Errorf("cannot find node enclosing position")
	}

	// If the position is not an identifier but immediately follows
	// an identifier or selector period (as is common when
	// requesting a completion), use the path to the preceding node.
	if _, ok := path[0].(*ast.Ident); !ok {
		if p, _ := astutil.PathEnclosingInterval(file, pos-1, pos-1); p != nil {
			switch p[0].(type) {
			case *ast.Ident, *ast.SelectorExpr:
				path = p // use preceding ident/selector
			}
		}
	}

	// Skip completion inside comment blocks or string literals.
	switch lit := path[0].(type) {
	case *ast.File, *ast.BlockStmt:
		if inComment(pos, file.Comments) {
			return items, prefix, nil
		}
	case *ast.BasicLit:
		if lit.Kind == token.STRING {
			return items, prefix, nil
		}
	}

	// Save certain facts about the query position, including the expected type
	// of the completion result, the signature of the function enclosing the
	// position.
	typ := expectedType(path, pos, pkg.GetTypesInfo())
	sig := enclosingFunction(path, pos, pkg.GetTypesInfo())
	pkgStringer := qualifier(file, pkg.GetTypes(), pkg.GetTypesInfo())
	preferTypeNames := wantTypeNames(pos, path)

	seen := make(map[types.Object]bool)
	// found adds a candidate completion.
	// Only the first candidate of a given name is considered.
	found := func(obj types.Object, weight float64, items []CompletionItem) []CompletionItem {
		if obj.Pkg() != nil && obj.Pkg() != pkg.GetTypes() && !obj.Exported() {
			return items // inaccessible
		}

		if !seen[obj] {
			seen[obj] = true
			if typ != nil && matchingTypes(typ, obj.Type()) {
				weight *= 10.0
			}
			if _, ok := obj.(*types.TypeName); !ok && preferTypeNames {
				weight *= 0.01
			}
			item := formatCompletion(obj, pkgStringer, weight, func(v *types.Var) bool {
				return isParameter(sig, v)
			})
			items = append(items, item)
		}
		return items
	}

	// The position is within a composite literal.
	if items, prefix, ok := complit(path, pos, pkg.GetTypes(), pkg.GetTypesInfo(), found); ok {
		return items, prefix, nil
	}
	switch n := path[0].(type) {
	case *ast.Ident:
		// Set the filter prefix.
		prefix = n.Name[:pos-n.Pos()]

		// Is this the Sel part of a selector?
		if sel, ok := path[1].(*ast.SelectorExpr); ok && sel.Sel == n {
			items, err = selector(sel, pos, pkg.GetTypesInfo(), found)
			return items, prefix, err
=======
// Scoring constants are used for weighting the relevance of different candidates.
const (
	// stdScore is the base score for all completion items.
	stdScore float64 = 1.0

	// highScore indicates a very relevant completion item.
	highScore float64 = 10.0

	// lowScore indicates an irrelevant or not useful completion item.
	lowScore float64 = 0.01
)

// completer contains the necessary information for a single completion request.
type completer struct {
	// Package-specific fields.
	types *types.Package
	info  *types.Info
	qf    types.Qualifier
	fset  *token.FileSet

	// pos is the position at which the request was triggered.
	pos token.Pos

	// path is the path of AST nodes enclosing the position.
	path []ast.Node

	// seen is the map that ensures we do not return duplicate results.
	seen map[types.Object]bool

	// items is the list of completion items returned.
	items []CompletionItem

	// prefix is the already-typed portion of the completion candidates.
	prefix string

	// expectedType is the type we expect the completion candidate to be.
	// It may not be set.
	expectedType types.Type

	// enclosingFunction is the function declaration enclosing the position.
	enclosingFunction *types.Signature

	// preferTypeNames is true if we are completing at a position that expects a type,
	// not a value.
	preferTypeNames bool

	// enclosingCompositeLiteral is the composite literal enclosing the position.
	enclosingCompositeLiteral *ast.CompositeLit

	// enclosingKeyValue is the key value expression enclosing the position.
	enclosingKeyValue *ast.KeyValueExpr

	// inCompositeLiteralField is true if we are completing a composite literal field.
	inCompositeLiteralField bool
}

// found adds a candidate completion.
//
// Only the first candidate of a given name is considered.
func (c *completer) found(obj types.Object, weight float64) {
	if obj.Pkg() != nil && obj.Pkg() != c.types && !obj.Exported() {
		return // inaccessible
	}
	if c.seen[obj] {
		return
	}
	c.seen[obj] = true
	if c.matchingType(obj.Type()) {
		weight *= highScore
	}
	if _, ok := obj.(*types.TypeName); !ok && c.preferTypeNames {
		weight *= lowScore
	}
	c.items = append(c.items, c.item(obj, weight))
}

// Completion returns a list of possible candidates for completion, given a
// a file and a position.
//
// The prefix is computed based on the preceding identifier and can be used by
// the client to score the quality of the completion. For instance, some clients
// may tolerate imperfect matches as valid completion results, since users may make typos.
func Completion(ctx context.Context, f File, pos token.Pos) ([]CompletionItem, string, error) {
	file := f.GetAST(ctx)
	pkg := f.GetPackage(ctx)
	if pkg == nil || pkg.IsIllTyped() {
		return nil, "", fmt.Errorf("package for %s is ill typed", f.URI())
	}

	// Completion is based on what precedes the cursor.
	// Find the path to the position before pos.
	path, _ := astutil.PathEnclosingInterval(file, pos-1, pos-1)
	if path == nil {
		return nil, "", fmt.Errorf("cannot find node enclosing position")
	}
	// Skip completion inside comments.
	for _, g := range file.Comments {
		if g.Pos() <= pos && pos <= g.End() {
			return nil, "", nil
		}
	}
	// Skip completion inside any kind of literal.
	if _, ok := path[0].(*ast.BasicLit); ok {
		return nil, "", nil
	}

	lit, kv, inCompositeLiteralField := enclosingCompositeLiteral(path, pos)
	c := &completer{
		types:                     pkg.GetTypes(),
		info:                      pkg.GetTypesInfo(),
		qf:                        qualifier(file, pkg.GetTypes(), pkg.GetTypesInfo()),
		fset:                      f.GetFileSet(ctx),
		path:                      path,
		pos:                       pos,
		seen:                      make(map[types.Object]bool),
		expectedType:              expectedType(path, pos, pkg.GetTypesInfo()),
		enclosingFunction:         enclosingFunction(path, pos, pkg.GetTypesInfo()),
		preferTypeNames:           preferTypeNames(path, pos),
		enclosingCompositeLiteral: lit,
		enclosingKeyValue:         kv,
		inCompositeLiteralField:   inCompositeLiteralField,
	}

	// Composite literals are handled entirely separately.
	if c.enclosingCompositeLiteral != nil {
		c.expectedType = c.expectedCompositeLiteralType(c.enclosingCompositeLiteral, c.enclosingKeyValue)

		if c.inCompositeLiteralField {
			if err := c.compositeLiteral(c.enclosingCompositeLiteral, c.enclosingKeyValue); err != nil {
				return nil, "", err
			}
			return c.items, c.prefix, nil
		}
	}

	switch n := path[0].(type) {
	case *ast.Ident:
		// Set the filter prefix.
		c.prefix = n.Name[:pos-n.Pos()]

		// Is this the Sel part of a selector?
		if sel, ok := path[1].(*ast.SelectorExpr); ok && sel.Sel == n {
			if err := c.selector(sel); err != nil {
				return nil, "", err
			}
			return c.items, c.prefix, nil
>>>>>>> v0.0.4
		}
		// reject defining identifiers
		if obj, ok := pkg.GetTypesInfo().Defs[n]; ok {
			if v, ok := obj.(*types.Var); ok && v.IsField() {
				// An anonymous field is also a reference to a type.
			} else {
				of := ""
				if obj != nil {
					qual := types.RelativeTo(pkg.GetTypes())
					of += ", of " + types.ObjectString(obj, qual)
				}
				return nil, "", fmt.Errorf("this is a definition%s", of)
			}
		}
<<<<<<< HEAD

		items = append(items, lexical(path, pos, pkg.GetTypes(), pkg.GetTypesInfo(), found)...)
=======
		if err := c.lexical(); err != nil {
			return nil, "", err
		}
>>>>>>> v0.0.4

	// The function name hasn't been typed yet, but the parens are there:
	//   recv.‸(arg)
	case *ast.TypeAssertExpr:
		// Create a fake selector expression.
<<<<<<< HEAD
		items, err = selector(&ast.SelectorExpr{X: n.X}, pos, pkg.GetTypesInfo(), found)
		return items, prefix, err

	case *ast.SelectorExpr:
		items, err = selector(n, pos, pkg.GetTypesInfo(), found)
		return items, prefix, err

	default:
		// fallback to lexical completions
		return lexical(path, pos, pkg.GetTypes(), pkg.GetTypesInfo(), found), "", nil
	}
	return items, prefix, nil
}

// selector finds completions for
// the specified selector expression.
// TODO(rstambler): Set the prefix filter correctly for selectors.
func selector(sel *ast.SelectorExpr, pos token.Pos, info *types.Info, found finder) (items []CompletionItem, err error) {
	// Is sel a qualified identifier?
	if id, ok := sel.X.(*ast.Ident); ok {
		if pkgname, ok := info.Uses[id].(*types.PkgName); ok {
			// Enumerate package members.
			// TODO(adonovan): can Imported() be nil?
			scope := pkgname.Imported().Scope()
			// TODO testcase: bad import
			for _, name := range scope.Names() {
				items = found(scope.Lookup(name), stdScore, items)
			}
			return items, nil
		}
	}

	// Inv: sel is a true selector.
	tv, ok := info.Types[sel.X]
	if !ok {
		return nil, fmt.Errorf("cannot resolve %s", sel.X)
	}

	// methods of T
	mset := types.NewMethodSet(tv.Type)
	for i := 0; i < mset.Len(); i++ {
		items = found(mset.At(i).Obj(), stdScore, items)
	}

	// methods of *T
	if tv.Addressable() && !types.IsInterface(tv.Type) && !isPointer(tv.Type) {
		mset := types.NewMethodSet(types.NewPointer(tv.Type))
		for i := 0; i < mset.Len(); i++ {
			items = found(mset.At(i).Obj(), stdScore, items)
		}
	}

	// fields of T
	for _, f := range fieldSelections(tv.Type) {
		items = found(f, stdScore, items)
	}

	return items, nil
}

// wantTypeNames checks if given token position is inside func receiver, type params
// or type results (e.g func (<>) foo(<>) (<>) {} ).
func wantTypeNames(pos token.Pos, path []ast.Node) bool {
	for _, p := range path {
		switch n := p.(type) {
		case *ast.FuncDecl:
			recv := n.Recv
			if recv != nil && recv.Pos() <= pos && pos <= recv.End() {
				return true
			}

			if n.Type != nil {
				params := n.Type.Params
				if params != nil && params.Pos() <= pos && pos <= params.End() {
					return true
				}

				results := n.Type.Results
				if results != nil && results.Pos() <= pos && pos <= results.End() {
					return true
				}
			}
			return false
		}
	}
	return false
}

// lexical finds completions in the lexical environment.
func lexical(path []ast.Node, pos token.Pos, pkg *types.Package, info *types.Info, found finder) (items []CompletionItem) {
	var scopes []*types.Scope // scopes[i], where i<len(path), is the possibly nil Scope of path[i].
	for _, n := range path {
=======
		if err := c.selector(&ast.SelectorExpr{X: n.X}); err != nil {
			return nil, "", err
		}

	case *ast.SelectorExpr:
		if err := c.selector(n); err != nil {
			return nil, "", err
		}

	default:
		// fallback to lexical completions
		if err := c.lexical(); err != nil {
			return nil, "", err
		}
	}
	return c.items, c.prefix, nil
}

// selector finds completions for the specified selector expression.
func (c *completer) selector(sel *ast.SelectorExpr) error {
	// Is sel a qualified identifier?
	if id, ok := sel.X.(*ast.Ident); ok {
		if pkgname, ok := c.info.Uses[id].(*types.PkgName); ok {
			// Enumerate package members.
			scope := pkgname.Imported().Scope()
			for _, name := range scope.Names() {
				c.found(scope.Lookup(name), stdScore)
			}
			return nil
		}
	}

	// Invariant: sel is a true selector.
	tv, ok := c.info.Types[sel.X]
	if !ok {
		return fmt.Errorf("cannot resolve %s", sel.X)
	}

	// Add methods of T.
	mset := types.NewMethodSet(tv.Type)
	for i := 0; i < mset.Len(); i++ {
		c.found(mset.At(i).Obj(), stdScore)
	}

	// Add methods of *T.
	if tv.Addressable() && !types.IsInterface(tv.Type) && !isPointer(tv.Type) {
		mset := types.NewMethodSet(types.NewPointer(tv.Type))
		for i := 0; i < mset.Len(); i++ {
			c.found(mset.At(i).Obj(), stdScore)
		}
	}

	// Add fields of T.
	for _, f := range fieldSelections(tv.Type) {
		c.found(f, stdScore)
	}
	return nil
}

// lexical finds completions in the lexical environment.
func (c *completer) lexical() error {
	var scopes []*types.Scope // scopes[i], where i<len(path), is the possibly nil Scope of path[i].
	for _, n := range c.path {
>>>>>>> v0.0.4
		switch node := n.(type) {
		case *ast.FuncDecl:
			n = node.Type
		case *ast.FuncLit:
			n = node.Type
		}
<<<<<<< HEAD
		scopes = append(scopes, info.Scopes[n])
	}
	scopes = append(scopes, pkg.Scope(), types.Universe)
=======
		scopes = append(scopes, c.info.Scopes[n])
	}
	scopes = append(scopes, c.types.Scope(), types.Universe)

	// Track seen variables to avoid showing completions for shadowed variables.
	// This works since we look at scopes from innermost to outermost.
	seen := make(map[string]struct{})
>>>>>>> v0.0.4

	// Process scopes innermost first.
	for i, scope := range scopes {
		if scope == nil {
			continue
		}
		for _, name := range scope.Names() {
<<<<<<< HEAD
			declScope, obj := scope.LookupParent(name, pos)
=======
			declScope, obj := scope.LookupParent(name, c.pos)
>>>>>>> v0.0.4
			if declScope != scope {
				continue // Name was declared in some enclosing scope, or not at all.
			}
			// If obj's type is invalid, find the AST node that defines the lexical block
			// containing the declaration of obj. Don't resolve types for packages.
			if _, ok := obj.(*types.PkgName); !ok && obj.Type() == types.Typ[types.Invalid] {
				// Match the scope to its ast.Node. If the scope is the package scope,
				// use the *ast.File as the starting node.
				var node ast.Node
<<<<<<< HEAD
				if i < len(path) {
					node = path[i]
				} else if i == len(path) { // use the *ast.File for package scope
					node = path[i-1]
				}
				if node != nil {
					if resolved := resolveInvalid(obj, node, info); resolved != nil {
=======
				if i < len(c.path) {
					node = c.path[i]
				} else if i == len(c.path) { // use the *ast.File for package scope
					node = c.path[i-1]
				}
				if node != nil {
					if resolved := resolveInvalid(obj, node, c.info); resolved != nil {
>>>>>>> v0.0.4
						obj = resolved
					}
				}
			}

			score := stdScore
			// Rank builtins significantly lower than other results.
			if scope == types.Universe {
				score *= 0.1
			}
<<<<<<< HEAD
			items = found(obj, score, items)
		}
	}
	return items
}

// inComment checks if given token position is inside ast.Comment node.
func inComment(pos token.Pos, commentGroups []*ast.CommentGroup) bool {
	for _, g := range commentGroups {
		for _, c := range g.List {
			if c.Pos() <= pos && pos <= c.End() {
				return true
			}
		}
	}
	return false
}

// complit finds completions for field names inside a composite literal.
// It reports whether the node was handled as part of a composite literal.
func complit(path []ast.Node, pos token.Pos, pkg *types.Package, info *types.Info, found finder) (items []CompletionItem, prefix string, ok bool) {
	var lit *ast.CompositeLit

	// First, determine if the pos is within a composite literal.
	switch n := path[0].(type) {
	case *ast.CompositeLit:
		// The enclosing node will be a composite literal if the user has just
		// opened the curly brace (e.g. &x{<>) or the completion request is triggered
		// from an already completed composite literal expression (e.g. &x{foo: 1, <>})
		//
		// If the cursor position is within a key-value expression inside the composite
		// literal, we try to determine if it is before or after the colon. If it is before
		// the colon, we return field completions. If the cursor does not belong to any
		// expression within the composite literal, we show composite literal completions.
		var expr ast.Expr
		for _, e := range n.Elts {
			if e.Pos() <= pos && pos < e.End() {
				expr = e
				break
			}
		}
		lit = n
		// If the position belongs to a key-value expression and is after the colon,
		// don't show composite literal completions.
		if kv, ok := expr.(*ast.KeyValueExpr); ok && pos > kv.Colon {
			lit = nil
		}
	case *ast.KeyValueExpr:
		// If the enclosing node is a key-value expression (e.g. &x{foo: <>}),
		// we show composite literal completions if the cursor position is before the colon.
		if len(path) > 1 && pos < n.Colon {
			if l, ok := path[1].(*ast.CompositeLit); ok {
				lit = l
			}
		}
	case *ast.Ident:
		prefix = n.Name[:pos-n.Pos()]

		// If the enclosing node is an identifier, it can either be an identifier that is
		// part of a composite literal (e.g. &x{fo<>}), or it can be an identifier that is
		// part of a key-value expression, which is part of a composite literal (e.g. &x{foo: ba<>).
		// We handle both of these cases, showing composite literal completions only if
		// the cursor position for the key-value expression is before the colon.
		if len(path) > 1 {
			if l, ok := path[1].(*ast.CompositeLit); ok {
				lit = l
			} else if len(path) > 2 {
				if l, ok := path[2].(*ast.CompositeLit); ok {
					// Confirm that cursor position is inside curly braces.
					if l.Lbrace <= pos && pos <= l.Rbrace {
						lit = l
						if kv, ok := path[1].(*ast.KeyValueExpr); ok {
							if pos > kv.Colon {
								lit = nil
							}
						}
					}
				}
			}
		}
	}
	// We are not in a composite literal.
	if lit == nil {
		return nil, prefix, false
	}
	// Mark fields of the composite literal that have already been set,
	// except for the current field.
	hasKeys := false // true if the composite literal already has key-value pairs
	addedFields := make(map[*types.Var]bool)
	for _, el := range lit.Elts {
		if kv, ok := el.(*ast.KeyValueExpr); ok {
			hasKeys = true
			if kv.Pos() <= pos && pos <= kv.End() {
				continue
			}
			if key, ok := kv.Key.(*ast.Ident); ok {
				if used, ok := info.Uses[key]; ok {
=======
			// If we haven't already added a candidate for an object with this name.
			if _, ok := seen[obj.Name()]; !ok {
				seen[obj.Name()] = struct{}{}
				c.found(obj, score)
			}
		}
	}
	return nil
}

// compositeLiteral finds completions for field names inside a composite literal.
func (c *completer) compositeLiteral(lit *ast.CompositeLit, kv *ast.KeyValueExpr) error {
	switch n := c.path[0].(type) {
	case *ast.Ident:
		c.prefix = n.Name[:c.pos-n.Pos()]
	}
	// Mark fields of the composite literal that have already been set,
	// except for the current field.
	hasKeys := kv != nil // true if the composite literal already has key-value pairs
	addedFields := make(map[*types.Var]bool)
	for _, el := range lit.Elts {
		if kvExpr, ok := el.(*ast.KeyValueExpr); ok {
			if kv == kvExpr {
				continue
			}

			hasKeys = true
			if key, ok := kvExpr.Key.(*ast.Ident); ok {
				if used, ok := c.info.Uses[key]; ok {
>>>>>>> v0.0.4
					if usedVar, ok := used.(*types.Var); ok {
						addedFields[usedVar] = true
					}
				}
			}
		}
	}
	// If the underlying type of the composite literal is a struct,
	// collect completions for the fields of this struct.
<<<<<<< HEAD
	if tv, ok := info.Types[lit]; ok {
		var structPkg *types.Package // package containing the struct type declaration
		if s, ok := tv.Type.Underlying().(*types.Struct); ok {
			for i := 0; i < s.NumFields(); i++ {
				field := s.Field(i)
=======
	if tv, ok := c.info.Types[lit]; ok {
		switch t := tv.Type.Underlying().(type) {
		case *types.Struct:
			var structPkg *types.Package // package that struct is declared in
			for i := 0; i < t.NumFields(); i++ {
				field := t.Field(i)
>>>>>>> v0.0.4
				if i == 0 {
					structPkg = field.Pkg()
				}
				if !addedFields[field] {
<<<<<<< HEAD
					items = found(field, 10.0, items)
=======
					c.found(field, highScore)
>>>>>>> v0.0.4
				}
			}
			// Add lexical completions if the user hasn't typed a key value expression
			// and if the struct fields are defined in the same package as the user is in.
<<<<<<< HEAD
			if !hasKeys && structPkg == pkg {
				items = append(items, lexical(path, pos, pkg, info, found)...)
			}
			return items, prefix, true
		}
	}
	return items, prefix, false
}

// formatCompletion creates a completion item for a given types.Object.
func formatCompletion(obj types.Object, qualifier types.Qualifier, score float64, isParam func(*types.Var) bool) CompletionItem {
	label := obj.Name()
	detail := types.TypeString(obj.Type(), qualifier)
	var kind CompletionItemKind

	switch o := obj.(type) {
	case *types.TypeName:
		detail, kind = formatType(o.Type(), qualifier)
		if obj.Parent() == types.Universe {
			detail = ""
		}
	case *types.Const:
		if obj.Parent() == types.Universe {
			detail = ""
		} else {
			val := o.Val().ExactString()
			if !strings.Contains(val, "\\n") { // skip any multiline constants
				label += " = " + o.Val().ExactString()
			}
		}
		kind = ConstantCompletionItem
	case *types.Var:
		if _, ok := o.Type().(*types.Struct); ok {
			detail = "struct{...}" // for anonymous structs
		}
		if o.IsField() {
			kind = FieldCompletionItem
		} else if isParam(o) {
			kind = ParameterCompletionItem
		} else {
			kind = VariableCompletionItem
		}
	case *types.Func:
		if sig, ok := o.Type().(*types.Signature); ok {
			label += formatParams(sig.Params(), sig.Variadic(), qualifier)
			detail = strings.Trim(types.TypeString(sig.Results(), qualifier), "()")
			kind = FunctionCompletionItem
			if sig.Recv() != nil {
				kind = MethodCompletionItem
			}
		}
	case *types.Builtin:
		item, ok := builtinDetails[obj.Name()]
		if !ok {
			break
		}
		label, detail = item.label, item.detail
		kind = FunctionCompletionItem
	case *types.PkgName:
		kind = PackageCompletionItem
		detail = fmt.Sprintf("\"%s\"", o.Imported().Path())
	case *types.Nil:
		kind = VariableCompletionItem
		detail = ""
	}
	detail = strings.TrimPrefix(detail, "untyped ")

	return CompletionItem{
		Label:  label,
		Detail: detail,
		Kind:   kind,
		Score:  score,
	}
}

// formatType returns the detail and kind for an object of type *types.TypeName.
func formatType(typ types.Type, qualifier types.Qualifier) (detail string, kind CompletionItemKind) {
	if types.IsInterface(typ) {
		detail = "interface{...}"
		kind = InterfaceCompletionItem
	} else if _, ok := typ.(*types.Struct); ok {
		detail = "struct{...}"
		kind = StructCompletionItem
	} else if typ != typ.Underlying() {
		detail, kind = formatType(typ.Underlying(), qualifier)
	} else {
		detail = types.TypeString(typ, qualifier)
		kind = TypeCompletionItem
	}
	return detail, kind
}

// formatParams correctly format the parameters of a function.
func formatParams(t *types.Tuple, variadic bool, qualifier types.Qualifier) string {
	var b bytes.Buffer
	b.WriteByte('(')
	for i := 0; i < t.Len(); i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		el := t.At(i)
		typ := types.TypeString(el.Type(), qualifier)
		// Handle a variadic parameter (can only be the final parameter).
		if variadic && i == t.Len()-1 {
			typ = strings.Replace(typ, "[]", "...", 1)
		}
		fmt.Fprintf(&b, "%v %v", el.Name(), typ)
	}
	b.WriteByte(')')
	return b.String()
}

// isParameter returns true if the given *types.Var is a parameter to the given
// *types.Signature.
func isParameter(sig *types.Signature, v *types.Var) bool {
	if sig == nil {
		return false
	}
	for i := 0; i < sig.Params().Len(); i++ {
		if sig.Params().At(i) == v {
			return true
		}
	}
	return false
}

// qualifier returns a function that appropriately formats a types.PkgName
// appearing in a *ast.File.
func qualifier(f *ast.File, pkg *types.Package, info *types.Info) types.Qualifier {
	// Construct mapping of import paths to their defined or implicit names.
	imports := make(map[*types.Package]string)
	for _, imp := range f.Imports {
		var obj types.Object
		if imp.Name != nil {
			obj = info.Defs[imp.Name]
		} else {
			obj = info.Implicits[imp]
		}
		if pkgname, ok := obj.(*types.PkgName); ok {
			imports[pkgname.Imported()] = pkgname.Name()
		}
	}
	// Define qualifier to replace full package paths with names of the imports.
	return func(p *types.Package) string {
		if p == pkg {
			return ""
		}
		if name, ok := imports[p]; ok {
			return name
		}
		return p.Name()
	}
}

// enclosingFunction returns the signature of the function enclosing the given
// position.
=======
			if !hasKeys && structPkg == c.types {
				return c.lexical()
			}
		default:
			return c.lexical()
		}
	}
	return nil
}

func enclosingCompositeLiteral(path []ast.Node, pos token.Pos) (lit *ast.CompositeLit, kv *ast.KeyValueExpr, ok bool) {
	for _, n := range path {
		switch n := n.(type) {
		case *ast.CompositeLit:
			// The enclosing node will be a composite literal if the user has just
			// opened the curly brace (e.g. &x{<>) or the completion request is triggered
			// from an already completed composite literal expression (e.g. &x{foo: 1, <>})
			//
			// The position is not part of the composite literal unless it falls within the
			// curly braces (e.g. "foo.Foo<>Struct{}").
			if n.Lbrace <= pos && pos <= n.Rbrace {
				lit = n

				// If the cursor position is within a key-value expression inside the composite
				// literal, we try to determine if it is before or after the colon. If it is before
				// the colon, we return field completions. If the cursor does not belong to any
				// expression within the composite literal, we show composite literal completions.
				if expr, isKeyValue := exprAtPos(pos, n.Elts).(*ast.KeyValueExpr); kv == nil && isKeyValue {
					kv = expr

					// If the position belongs to a key-value expression and is after the colon,
					// don't show composite literal completions.
					ok = pos <= kv.Colon
				} else if kv == nil {
					ok = true
				}
			}
			return lit, kv, ok
		case *ast.KeyValueExpr:
			if kv == nil {
				kv = n

				// If the position belongs to a key-value expression and is after the colon,
				// don't show composite literal completions.
				ok = pos <= kv.Colon
			}
		case *ast.FuncType, *ast.CallExpr, *ast.TypeAssertExpr:
			// These node types break the type link between the leaf node and
			// the composite literal. The type of the leaf node becomes unrelated
			// to the type of the composite literal, so we return nil to avoid
			// inappropriate completions. For example, "Foo{Bar: x.Baz(<>)}"
			// should complete as a function argument to Baz, not part of the Foo
			// composite literal.
			return nil, nil, false
		}
	}
	return lit, kv, ok
}

// enclosingFunction returns the signature of the function enclosing the given position.
>>>>>>> v0.0.4
func enclosingFunction(path []ast.Node, pos token.Pos, info *types.Info) *types.Signature {
	for _, node := range path {
		switch t := node.(type) {
		case *ast.FuncDecl:
			if obj, ok := info.Defs[t.Name]; ok {
				return obj.Type().(*types.Signature)
			}
		case *ast.FuncLit:
			if typ, ok := info.Types[t]; ok {
				return typ.Type.(*types.Signature)
			}
		}
	}
	return nil
}

<<<<<<< HEAD
=======
func (c *completer) expectedCompositeLiteralType(lit *ast.CompositeLit, kv *ast.KeyValueExpr) types.Type {
	litType, ok := c.info.Types[lit]
	if !ok {
		return nil
	}
	switch t := litType.Type.Underlying().(type) {
	case *types.Slice:
		return t.Elem()
	case *types.Array:
		return t.Elem()
	case *types.Map:
		if kv == nil || c.pos <= kv.Colon {
			return t.Key()
		}
		return t.Elem()
	case *types.Struct:
		//  If we are in a key-value expression.
		if kv != nil {
			// There is no expected type for a struct field name.
			if c.pos <= kv.Colon {
				return nil
			}
			// Find the type of the struct field whose name matches the key.
			if key, ok := kv.Key.(*ast.Ident); ok {
				for i := 0; i < t.NumFields(); i++ {
					if field := t.Field(i); field.Name() == key.Name {
						return field.Type()
					}
				}
			}
			return nil
		}
		// We are in a struct literal, but not a specific key-value pair.
		// If the struct literal doesn't have explicit field names,
		// we may still be able to suggest an expected type.
		for _, el := range lit.Elts {
			if _, ok := el.(*ast.KeyValueExpr); ok {
				return nil
			}
		}
		// The order of the literal fields must match the order in the struct definition.
		// Find the element that the position belongs to and suggest that field's type.
		if i := indexExprAtPos(c.pos, lit.Elts); i < t.NumFields() {
			return t.Field(i).Type()
		}
	}
	return nil
}

>>>>>>> v0.0.4
// expectedType returns the expected type for an expression at the query position.
func expectedType(path []ast.Node, pos token.Pos, info *types.Info) types.Type {
	for i, node := range path {
		if i == 2 {
			break
		}
		switch expr := node.(type) {
		case *ast.BinaryExpr:
			// Determine if query position comes from left or right of op.
			e := expr.X
			if pos < expr.OpPos {
				e = expr.Y
			}
			if tv, ok := info.Types[e]; ok {
				return tv.Type
			}
		case *ast.AssignStmt:
			// Only rank completions if you are on the right side of the token.
			if pos <= expr.TokPos {
				break
			}
<<<<<<< HEAD
			i := exprAtPos(pos, expr.Rhs)
=======
			i := indexExprAtPos(pos, expr.Rhs)
>>>>>>> v0.0.4
			if i >= len(expr.Lhs) {
				i = len(expr.Lhs) - 1
			}
			if tv, ok := info.Types[expr.Lhs[i]]; ok {
				return tv.Type
			}
		case *ast.CallExpr:
			if tv, ok := info.Types[expr.Fun]; ok {
				if sig, ok := tv.Type.(*types.Signature); ok {
					if sig.Params().Len() == 0 {
						return nil
					}
<<<<<<< HEAD
					i := exprAtPos(pos, expr.Args)
=======
					i := indexExprAtPos(pos, expr.Args)
>>>>>>> v0.0.4
					// Make sure not to run past the end of expected parameters.
					if i >= sig.Params().Len() {
						i = sig.Params().Len() - 1
					}
					return sig.Params().At(i).Type()
				}
			}
		}
	}
	return nil
}

<<<<<<< HEAD
// matchingTypes reports whether actual is a good candidate type
// for a completion in a context of the expected type.
func matchingTypes(expected, actual types.Type) bool {
	// Use a function's return type as its type.
	if sig, ok := actual.(*types.Signature); ok {
		if sig.Results().Len() == 1 {
			actual = sig.Results().At(0).Type()
		}
	}
	return types.Identical(types.Default(expected), types.Default(actual))
}

// exprAtPos returns the index of the expression containing pos.
func exprAtPos(pos token.Pos, args []ast.Expr) int {
	for i, expr := range args {
		if expr.Pos() <= pos && pos <= expr.End() {
			return i
		}
	}
	return len(args)
}

// fieldSelections returns the set of fields that can
// be selected from a value of type T.
func fieldSelections(T types.Type) (fields []*types.Var) {
	// TODO(adonovan): this algorithm doesn't exclude ambiguous
	// selections that match more than one field/method.
	// types.NewSelectionSet should do that for us.

	seen := make(map[types.Type]bool) // for termination on recursive types
	var visit func(T types.Type)
	visit = func(T types.Type) {
		if !seen[T] {
			seen[T] = true
			if T, ok := deref(T).Underlying().(*types.Struct); ok {
				for i := 0; i < T.NumFields(); i++ {
					f := T.Field(i)
					fields = append(fields, f)
					if f.Anonymous() {
						visit(f.Type())
					}
				}
			}
		}
	}
	visit(T)

	return fields
}

func isPointer(T types.Type) bool {
	_, ok := T.(*types.Pointer)
	return ok
}

// deref returns a pointer's element type; otherwise it returns typ.
func deref(typ types.Type) types.Type {
	if p, ok := typ.Underlying().(*types.Pointer); ok {
		return p.Elem()
	}
	return typ
}

// resolveInvalid traverses the node of the AST that defines the scope
// containing the declaration of obj, and attempts to find a user-friendly
// name for its invalid type. The resulting Object and its Type are fake.
func resolveInvalid(obj types.Object, node ast.Node, info *types.Info) types.Object {
	// Construct a fake type for the object and return a fake object with this type.
	formatResult := func(expr ast.Expr) types.Object {
		var typename string
		switch t := expr.(type) {
		case *ast.SelectorExpr:
			typename = fmt.Sprintf("%s.%s", t.X, t.Sel)
		case *ast.Ident:
			typename = t.String()
		default:
			return nil
		}
		typ := types.NewNamed(types.NewTypeName(token.NoPos, obj.Pkg(), typename, nil), nil, nil)
		return types.NewVar(obj.Pos(), obj.Pkg(), obj.Name(), typ)
	}
	var resultExpr ast.Expr
	ast.Inspect(node, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.ValueSpec:
			for _, name := range n.Names {
				if info.Defs[name] == obj {
					resultExpr = n.Type
				}
			}
			return false
		case *ast.Field: // This case handles parameters and results of a FuncDecl or FuncLit.
			for _, name := range n.Names {
				if info.Defs[name] == obj {
					resultExpr = n.Type
				}
			}
			return false
		// TODO(rstambler): Handle range statements.
		default:
			return true
		}
	})
	return formatResult(resultExpr)
}

type itemDetails struct {
	label, detail string
}

var builtinDetails = map[string]itemDetails{
	"append": { // append(slice []T, elems ...T)
		label:  "append(slice []T, elems ...T)",
		detail: "[]T",
	},
	"cap": { // cap(v []T) int
		label:  "cap(v []T)",
		detail: "int",
	},
	"close": { // close(c chan<- T)
		label: "close(c chan<- T)",
	},
	"complex": { // complex(r, i float64) complex128
		label:  "complex(real float64, imag float64)",
		detail: "complex128",
	},
	"copy": { // copy(dst, src []T) int
		label:  "copy(dst []T, src []T)",
		detail: "int",
	},
	"delete": { // delete(m map[T]T1, key T)
		label: "delete(m map[K]V, key K)",
	},
	"imag": { // imag(c complex128) float64
		label:  "imag(complex128)",
		detail: "float64",
	},
	"len": { // len(v T) int
		label:  "len(T)",
		detail: "int",
	},
	"make": { // make(t T, size ...int) T
		label:  "make(t T, size ...int)",
		detail: "T",
	},
	"new": { // new(T) *T
		label:  "new(T)",
		detail: "*T",
	},
	"panic": { // panic(v interface{})
		label: "panic(interface{})",
	},
	"print": { // print(args ...T)
		label: "print(args ...T)",
	},
	"println": { // println(args ...T)
		label: "println(args ...T)",
	},
	"real": { // real(c complex128) float64
		label:  "real(complex128)",
		detail: "float64",
	},
	"recover": { // recover() interface{}
		label:  "recover()",
		detail: "interface{}",
	},
=======
// preferTypeNames checks if given token position is inside func receiver,
// type params, or type results. For example:
//
// func (<>) foo(<>) (<>) {}
//
func preferTypeNames(path []ast.Node, pos token.Pos) bool {
	for _, p := range path {
		switch n := p.(type) {
		case *ast.FuncDecl:
			if r := n.Recv; r != nil && r.Pos() <= pos && pos <= r.End() {
				return true
			}
			if t := n.Type; t != nil {
				if p := t.Params; p != nil && p.Pos() <= pos && pos <= p.End() {
					return true
				}
				if r := t.Results; r != nil && r.Pos() <= pos && pos <= r.End() {
					return true
				}
			}
			return false
		}
	}
	return false
}

// matchingTypes reports whether actual is a good candidate type
// for a completion in a context of the expected type.
func (c *completer) matchingType(actual types.Type) bool {
	if c.expectedType == nil {
		return false
	}
	// Use a function's return type as its type.
	if sig, ok := actual.(*types.Signature); ok {
		if sig.Results().Len() == 1 {
			actual = sig.Results().At(0).Type()
		}
	}
	return types.Identical(types.Default(c.expectedType), types.Default(actual))
>>>>>>> v0.0.4
}
