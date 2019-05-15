// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
<<<<<<< HEAD
	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
)

=======
	"context"

	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/span"
)

func (s *Server) documentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) ([]protocol.DocumentSymbol, error) {
	uri := span.NewURI(params.TextDocument.URI)
	view := s.findView(ctx, uri)
	f, m, err := newColumnMap(ctx, view, uri)
	if err != nil {
		return nil, err
	}
	symbols := source.DocumentSymbols(ctx, f)
	return toProtocolDocumentSymbols(m, symbols), nil
}

>>>>>>> v0.0.4
func toProtocolDocumentSymbols(m *protocol.ColumnMapper, symbols []source.Symbol) []protocol.DocumentSymbol {
	result := make([]protocol.DocumentSymbol, 0, len(symbols))
	for _, s := range symbols {
		ps := protocol.DocumentSymbol{
			Name:     s.Name,
			Kind:     toProtocolSymbolKind(s.Kind),
			Detail:   s.Detail,
			Children: toProtocolDocumentSymbols(m, s.Children),
		}
		if r, err := m.Range(s.Span); err == nil {
			ps.Range = r
		}
		if r, err := m.Range(s.SelectionSpan); err == nil {
			ps.SelectionRange = r
		}
		result = append(result, ps)
	}
	return result
}

func toProtocolSymbolKind(kind source.SymbolKind) protocol.SymbolKind {
	switch kind {
	case source.StructSymbol:
		return protocol.Struct
	case source.PackageSymbol:
		return protocol.Package
	case source.VariableSymbol:
		return protocol.Variable
	case source.ConstantSymbol:
		return protocol.Constant
	case source.FunctionSymbol:
		return protocol.Function
	case source.MethodSymbol:
		return protocol.Method
	case source.InterfaceSymbol:
		return protocol.Interface
<<<<<<< HEAD
=======
	case source.NumberSymbol:
		return protocol.Number
	case source.StringSymbol:
		return protocol.String
	case source.BooleanSymbol:
		return protocol.Boolean
	case source.FieldSymbol:
		return protocol.Field
>>>>>>> v0.0.4
	default:
		return 0
	}
}
