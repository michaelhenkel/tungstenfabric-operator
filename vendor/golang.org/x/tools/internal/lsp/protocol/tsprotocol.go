// Package protocol contains data types for LSP jsonrpcs
// generated automatically from vscode-languageserver-node
<<<<<<< HEAD
//  version of Mon Feb 25 2019 09:01:22 GMT-0500 (Eastern Standard Time)
package protocol

// ImplementationClientCapabilities is:
type ImplementationClientCapabilities struct {

	/** TextDocument defined:
	 * The text document client capabilities
	 */
	TextDocument *struct {

		/** Implementation defined:
		 * Capabilities specific to the `textDocument/implementation`
		 */
		Implementation *struct {

			/** DynamicRegistration defined:
			 * Whether implementation supports dynamic registration. If this is set to `true`
			 * the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/** LinkSupport defined:
			 * The client supports additional metadata in the form of definition links.
			 */
			LinkSupport bool `json:"linkSupport,omitempty"`
		} `json:"implementation,omitempty"`
	} `json:"textDocument,omitempty"`
}

// ImplementationServerCapabilities is:
type ImplementationServerCapabilities struct {

	/** ImplementationProvider defined:
	 * The server provides Goto Implementation support.
	 */
	ImplementationProvider bool `json:"implementationProvider,omitempty"` // boolean | (TextDocumentRegistrationOptions & StaticRegistrationOptions)
}

// TypeDefinitionClientCapabilities is:
type TypeDefinitionClientCapabilities struct {

	/** TextDocument defined:
	 * The text document client capabilities
	 */
	TextDocument *struct {

		/** TypeDefinition defined:
		 * Capabilities specific to the `textDocument/typeDefinition`
		 */
		TypeDefinition *struct {

			/** DynamicRegistration defined:
			 * Whether implementation supports dynamic registration. If this is set to `true`
			 * the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/** LinkSupport defined:
			 * The client supports additional metadata in the form of definition links.
			 */
			LinkSupport bool `json:"linkSupport,omitempty"`
		} `json:"typeDefinition,omitempty"`
	} `json:"textDocument,omitempty"`
}

// TypeDefinitionServerCapabilities is:
type TypeDefinitionServerCapabilities struct {

	/** TypeDefinitionProvider defined:
	 * The server provides Goto Type Definition support.
	 */
	TypeDefinitionProvider bool `json:"typeDefinitionProvider,omitempty"` // boolean | (TextDocumentRegistrationOptions & StaticRegistrationOptions)
}

// WorkspaceFoldersInitializeParams is:
type WorkspaceFoldersInitializeParams struct {

	/** WorkspaceFolders defined:
	 * The actual configured workspace folders.
	 */
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders"`
}

// WorkspaceFoldersClientCapabilities is:
type WorkspaceFoldersClientCapabilities struct {

	/** Workspace defined:
	 * The workspace client capabilities
	 */
	Workspace *struct {

		/** WorkspaceFolders defined:
		 * The client has support for workspace folders
		 */
		WorkspaceFolders bool `json:"workspaceFolders,omitempty"`
	} `json:"workspace,omitempty"`
}

// WorkspaceFoldersServerCapabilities is:
type WorkspaceFoldersServerCapabilities struct {

	/** Workspace defined:
	 * The workspace server capabilities
	 */
	Workspace *struct {

		// WorkspaceFolders is
		WorkspaceFolders *struct {

			/** Supported defined:
			 * The Server has support for workspace folders
			 */
			Supported bool `json:"supported,omitempty"`

			/** ChangeNotifications defined:
			 * Whether the server wants to receive workspace folder
			 * change notifications.
			 *
			 * If a strings is provided the string is treated as a ID
			 * under which the notification is registed on the client
			 * side. The ID can be used to unregister for these events
			 * using the `client/unregisterCapability` request.
			 */
			ChangeNotifications string `json:"changeNotifications,omitempty"` // string | boolean
		} `json:"workspaceFolders,omitempty"`
	} `json:"workspace,omitempty"`
}

// WorkspaceFolder is:
type WorkspaceFolder struct {

	/** URI defined:
=======
//  version of Fri Apr 05 2019 10:16:07 GMT-0400 (Eastern Daylight Time)
package protocol

// WorkspaceFolder is
type WorkspaceFolder struct {

	/*URI defined:
>>>>>>> v0.0.4
	 * The associated URI for this workspace folder.
	 */
	URI string `json:"uri"`

<<<<<<< HEAD
	/** Name defined:
=======
	/*Name defined:
>>>>>>> v0.0.4
	 * The name of the workspace folder. Used to refer to this
	 * workspace folder in thge user interface.
	 */
	Name string `json:"name"`
}

<<<<<<< HEAD
// DidChangeWorkspaceFoldersParams is:
/**
=======
/*DidChangeWorkspaceFoldersParams defined:
>>>>>>> v0.0.4
 * The parameters of a `workspace/didChangeWorkspaceFolders` notification.
 */
type DidChangeWorkspaceFoldersParams struct {

<<<<<<< HEAD
	/** Event defined:
=======
	/*Event defined:
>>>>>>> v0.0.4
	 * The actual workspace folder change event.
	 */
	Event WorkspaceFoldersChangeEvent `json:"event"`
}

<<<<<<< HEAD
// WorkspaceFoldersChangeEvent is:
/**
=======
/*WorkspaceFoldersChangeEvent defined:
>>>>>>> v0.0.4
 * The workspace folder change event.
 */
type WorkspaceFoldersChangeEvent struct {

<<<<<<< HEAD
	/** Added defined:
=======
	/*Added defined:
>>>>>>> v0.0.4
	 * The array of added workspace folders
	 */
	Added []WorkspaceFolder `json:"added"`

<<<<<<< HEAD
	/** Removed defined:
=======
	/*Removed defined:
>>>>>>> v0.0.4
	 * The array of the removed workspace folders
	 */
	Removed []WorkspaceFolder `json:"removed"`
}

<<<<<<< HEAD
// ConfigurationClientCapabilities is:
type ConfigurationClientCapabilities struct {

	/** Workspace defined:
	 * The workspace client capabilities
	 */
	Workspace *struct {

		/** Configuration defined:
		* The client supports `workspace/configuration` requests.
		 */
		Configuration bool `json:"configuration,omitempty"`
	} `json:"workspace,omitempty"`
}

// ConfigurationItem is:
type ConfigurationItem struct {

	/** ScopeURI defined:
=======
// ConfigurationItem is
type ConfigurationItem struct {

	/*ScopeURI defined:
>>>>>>> v0.0.4
	 * The scope to get the configuration section for.
	 */
	ScopeURI string `json:"scopeUri,omitempty"`

<<<<<<< HEAD
	/** Section defined:
=======
	/*Section defined:
>>>>>>> v0.0.4
	 * The configuration section asked for.
	 */
	Section string `json:"section,omitempty"`
}

<<<<<<< HEAD
// ConfigurationParams is:
/**
=======
/*ConfigurationParams defined:
>>>>>>> v0.0.4
 * The parameters of a configuration request.
 */
type ConfigurationParams struct {

	// Items is
	Items []ConfigurationItem `json:"items"`
}

<<<<<<< HEAD
// ColorClientCapabilities is:
type ColorClientCapabilities struct {

	/** TextDocument defined:
	 * The text document client capabilities
	 */
	TextDocument *struct {

		/** ColorProvider defined:
		 * Capabilities specific to the colorProvider
		 */
		ColorProvider *struct {

			/** DynamicRegistration defined:
			 * Whether implementation supports dynamic registration. If this is set to `true`
			 * the client supports the new `(ColorProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"colorProvider,omitempty"`
	} `json:"textDocument,omitempty"`
}

// ColorProviderOptions is:
type ColorProviderOptions struct {
}

// ColorServerCapabilities is:
type ColorServerCapabilities struct {

	/** ColorProvider defined:
	 * The server provides color provider support.
	 */
	ColorProvider bool `json:"colorProvider,omitempty"` // boolean | ColorProviderOptions | (ColorProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)
}

// DocumentColorParams is:
/**
=======
// ColorProviderOptions is
type ColorProviderOptions struct {
}

/*DocumentColorParams defined:
>>>>>>> v0.0.4
 * Parameters for a [DocumentColorRequest](#DocumentColorRequest).
 */
type DocumentColorParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

<<<<<<< HEAD
// ColorPresentationParams is:
/**
=======
/*ColorPresentationParams defined:
>>>>>>> v0.0.4
 * Parameters for a [ColorPresentationRequest](#ColorPresentationRequest).
 */
type ColorPresentationParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Color defined:
=======
	/*Color defined:
>>>>>>> v0.0.4
	 * The color to request presentations for.
	 */
	Color Color `json:"color"`

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range where the color would be inserted. Serves as a context.
	 */
	Range Range `json:"range"`
}

<<<<<<< HEAD
// FoldingRangeClientCapabilities is:
type FoldingRangeClientCapabilities struct {

	/** TextDocument defined:
	 * The text document client capabilities
	 */
	TextDocument *struct {

		/** FoldingRange defined:
		 * Capabilities specific to `textDocument/foldingRange` requests
		 */
		FoldingRange *struct {

			/** DynamicRegistration defined:
			 * Whether implementation supports dynamic registration for folding range providers. If this is set to `true`
			 * the client supports the new `(FoldingRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/** RangeLimit defined:
			 * The maximum number of folding ranges that the client prefers to receive per document. The value serves as a
			 * hint, servers are free to follow the limit.
			 */
			RangeLimit float64 `json:"rangeLimit,omitempty"`

			/** LineFoldingOnly defined:
			 * If set, the client signals that it only supports folding complete lines. If set, client will
			 * ignore specified `startCharacter` and `endCharacter` properties in a FoldingRange.
			 */
			LineFoldingOnly bool `json:"lineFoldingOnly,omitempty"`
		} `json:"foldingRange,omitempty"`
	} `json:"textDocument,omitempty"`
}

// FoldingRangeProviderOptions is:
type FoldingRangeProviderOptions struct {
}

// FoldingRangeServerCapabilities is:
type FoldingRangeServerCapabilities struct {

	/** FoldingRangeProvider defined:
	 * The server provides folding provider support.
	 */
	FoldingRangeProvider bool `json:"foldingRangeProvider,omitempty"` // boolean | FoldingRangeProviderOptions | (FoldingRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)
}

// FoldingRange is:
/**
=======
// FoldingRangeProviderOptions is
type FoldingRangeProviderOptions struct {
}

/*FoldingRange defined:
>>>>>>> v0.0.4
 * Represents a folding range.
 */
type FoldingRange struct {

<<<<<<< HEAD
	/** StartLine defined:
=======
	/*StartLine defined:
>>>>>>> v0.0.4
	 * The zero-based line number from where the folded range starts.
	 */
	StartLine float64 `json:"startLine"`

<<<<<<< HEAD
	/** StartCharacter defined:
=======
	/*StartCharacter defined:
>>>>>>> v0.0.4
	 * The zero-based character offset from where the folded range starts. If not defined, defaults to the length of the start line.
	 */
	StartCharacter float64 `json:"startCharacter,omitempty"`

<<<<<<< HEAD
	/** EndLine defined:
=======
	/*EndLine defined:
>>>>>>> v0.0.4
	 * The zero-based line number where the folded range ends.
	 */
	EndLine float64 `json:"endLine"`

<<<<<<< HEAD
	/** EndCharacter defined:
=======
	/*EndCharacter defined:
>>>>>>> v0.0.4
	 * The zero-based character offset before the folded range ends. If not defined, defaults to the length of the end line.
	 */
	EndCharacter float64 `json:"endCharacter,omitempty"`

<<<<<<< HEAD
	/** Kind defined:
=======
	/*Kind defined:
>>>>>>> v0.0.4
	 * Describes the kind of the folding range such as `comment' or 'region'. The kind
	 * is used to categorize folding ranges and used by commands like 'Fold all comments'. See
	 * [FoldingRangeKind](#FoldingRangeKind) for an enumeration of standardized kinds.
	 */
	Kind string `json:"kind,omitempty"`
}

<<<<<<< HEAD
// FoldingRangeParams is:
/**
=======
/*FoldingRangeParams defined:
>>>>>>> v0.0.4
 * Parameters for a [FoldingRangeRequest](#FoldingRangeRequest).
 */
type FoldingRangeParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

<<<<<<< HEAD
// DeclarationClientCapabilities is:
type DeclarationClientCapabilities struct {

	/** TextDocument defined:
	 * The text document client capabilities
	 */
	TextDocument *struct {

		/** Declaration defined:
		 * Capabilities specific to the `textDocument/declaration`
		 */
		Declaration *struct {

			/** DynamicRegistration defined:
			 * Whether declaration supports dynamic registration. If this is set to `true`
			 * the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/** LinkSupport defined:
			 * The client supports additional metadata in the form of declaration links.
			 */
			LinkSupport bool `json:"linkSupport,omitempty"`
		} `json:"declaration,omitempty"`
	} `json:"textDocument,omitempty"`
}

// DeclarationServerCapabilities is:
type DeclarationServerCapabilities struct {

	/** DeclarationProvider defined:
	 * The server provides Goto Type Definition support.
	 */
	DeclarationProvider bool `json:"declarationProvider,omitempty"` // boolean | (TextDocumentRegistrationOptions & StaticRegistrationOptions)
}

// SelectionRangeClientCapabilities is:
type SelectionRangeClientCapabilities struct {

	/** TextDocument defined:
	 * The text document client capabilities
	 */
	TextDocument *struct {

		/** SelectionRange defined:
		 * Capabilities specific to `textDocument/selectionRange` requests
		 */
		SelectionRange *struct {

			/** DynamicRegistration defined:
			 * Whether implementation supports dynamic registration for selection range providers. If this is set to `true`
			 * the client supports the new `(SelectionRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"selectionRange,omitempty"`
	} `json:"textDocument,omitempty"`
}

// SelectionRangeProviderOptions is:
type SelectionRangeProviderOptions struct {
}

// SelectionRangeServerCapabilities is:
type SelectionRangeServerCapabilities struct {

	/** SelectionRangeProvider defined:
	 * The server provides selection range support.
	 */
	SelectionRangeProvider bool `json:"selectionRangeProvider,omitempty"` // boolean | SelectionRangeProviderOptions | (SelectionRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)
}

// SelectionRange is:
/**
=======
// SelectionRangeProviderOptions is
type SelectionRangeProviderOptions struct {
}

/*SelectionRange defined:
>>>>>>> v0.0.4
 * Represents a selection range
 */
type SelectionRange struct {

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * Range of the selection.
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** Kind defined:
=======
	/*Kind defined:
>>>>>>> v0.0.4
	 * Describes the kind of the selection range such as `statemet' or 'declaration'. See
	 * [SelectionRangeKind](#SelectionRangeKind) for an enumeration of standardized kinds.
	 */
	Kind string `json:"kind"`
}

<<<<<<< HEAD
// Registration is:
/**
=======
/*SelectionRangeParams defined:
 * A parameter literal used in selection range requests.
 */
type SelectionRangeParams struct {

	/*TextDocument defined:
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/*Positions defined:
	 * The positions inside the text document.
	 */
	Positions []Position `json:"positions"`
}

/*Registration defined:
>>>>>>> v0.0.4
 * General parameters to to register for an notification or to register a provider.
 */
type Registration struct {

<<<<<<< HEAD
	/** ID defined:
=======
	/*ID defined:
>>>>>>> v0.0.4
	 * The id used to register the request. The id can be used to deregister
	 * the request again.
	 */
	ID string `json:"id"`

<<<<<<< HEAD
	/** Method defined:
=======
	/*Method defined:
>>>>>>> v0.0.4
	 * The method to register for.
	 */
	Method string `json:"method"`

<<<<<<< HEAD
	/** RegisterOptions defined:
=======
	/*RegisterOptions defined:
>>>>>>> v0.0.4
	 * Options necessary for the registration.
	 */
	RegisterOptions interface{} `json:"registerOptions,omitempty"`
}

<<<<<<< HEAD
// RegistrationParams is:
=======
// RegistrationParams is
>>>>>>> v0.0.4
type RegistrationParams struct {

	// Registrations is
	Registrations []Registration `json:"registrations"`
}

<<<<<<< HEAD
// Unregistration is:
/**
=======
/*Unregistration defined:
>>>>>>> v0.0.4
 * General parameters to unregister a request or notification.
 */
type Unregistration struct {

<<<<<<< HEAD
	/** ID defined:
=======
	/*ID defined:
>>>>>>> v0.0.4
	 * The id used to unregister the request or notification. Usually an id
	 * provided during the register request.
	 */
	ID string `json:"id"`

<<<<<<< HEAD
	/** Method defined:
=======
	/*Method defined:
>>>>>>> v0.0.4
	 * The method to unregister for.
	 */
	Method string `json:"method"`
}

<<<<<<< HEAD
// UnregistrationParams is:
=======
// UnregistrationParams is
>>>>>>> v0.0.4
type UnregistrationParams struct {

	// Unregisterations is
	Unregisterations []Unregistration `json:"unregisterations"`
}

<<<<<<< HEAD
// TextDocumentPositionParams is:
/**
=======
/*TextDocumentPositionParams defined:
>>>>>>> v0.0.4
 * A parameter literal used in requests to pass a text document and a position inside that
 * document.
 */
type TextDocumentPositionParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Position defined:
=======
	/*Position defined:
>>>>>>> v0.0.4
	 * The position inside the text document.
	 */
	Position Position `json:"position"`
}

<<<<<<< HEAD
// WorkspaceClientCapabilities is:
/**
=======
/*WorkspaceClientCapabilities defined:
>>>>>>> v0.0.4
 * Workspace specific client capabilities.
 */
type WorkspaceClientCapabilities struct {

<<<<<<< HEAD
	/** ApplyEdit defined:
=======
	/*ApplyEdit defined:
>>>>>>> v0.0.4
	 * The client supports applying batch edits
	 * to the workspace by supporting the request
	 * 'workspace/applyEdit'
	 */
	ApplyEdit bool `json:"applyEdit,omitempty"`

<<<<<<< HEAD
	/** WorkspaceEdit defined:
=======
	/*WorkspaceEdit defined:
>>>>>>> v0.0.4
	 * Capabilities specific to `WorkspaceEdit`s
	 */
	WorkspaceEdit *struct {

<<<<<<< HEAD
		/** DocumentChanges defined:
=======
		/*DocumentChanges defined:
>>>>>>> v0.0.4
		 * The client supports versioned document changes in `WorkspaceEdit`s
		 */
		DocumentChanges bool `json:"documentChanges,omitempty"`

<<<<<<< HEAD
		/** ResourceOperations defined:
=======
		/*ResourceOperations defined:
>>>>>>> v0.0.4
		 * The resource operations the client supports. Clients should at least
		 * support 'create', 'rename' and 'delete' files and folders.
		 */
		ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitempty"`

<<<<<<< HEAD
		/** FailureHandling defined:
		 * The failure handling strategy of a client if applying the workspace edit
		 * failes.
		 */
		FailureHandling *FailureHandlingKind `json:"failureHandling,omitempty"`
	} `json:"workspaceEdit,omitempty"`

	/** DidChangeConfiguration defined:
=======
		/*FailureHandling defined:
		 * The failure handling strategy of a client if applying the workspace edit
		 * failes.
		 */
		FailureHandling FailureHandlingKind `json:"failureHandling,omitempty"`
	} `json:"workspaceEdit,omitempty"`

	/*DidChangeConfiguration defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `workspace/didChangeConfiguration` notification.
	 */
	DidChangeConfiguration *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Did change configuration notification supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"didChangeConfiguration,omitempty"`

<<<<<<< HEAD
	/** DidChangeWatchedFiles defined:
=======
	/*DidChangeWatchedFiles defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `workspace/didChangeWatchedFiles` notification.
	 */
	DidChangeWatchedFiles *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Did change watched files notification supports dynamic registration. Please note
		 * that the current protocol doesn't support static configuration for file changes
		 * from the server side.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"didChangeWatchedFiles,omitempty"`

<<<<<<< HEAD
	/** Symbol defined:
=======
	/*Symbol defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `workspace/symbol` request.
	 */
	Symbol *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Symbol request supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

<<<<<<< HEAD
		/** SymbolKind defined:
=======
		/*SymbolKind defined:
>>>>>>> v0.0.4
		 * Specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
		 */
		SymbolKind *struct {

<<<<<<< HEAD
			/** ValueSet defined:
=======
			/*ValueSet defined:
>>>>>>> v0.0.4
			 * The symbol kind values the client supports. When this
			 * property exists the client also guarantees that it will
			 * handle values outside its set gracefully and falls back
			 * to a default value when unknown.
			 *
			 * If this property is not present the client only supports
			 * the symbol kinds from `File` to `Array` as defined in
			 * the initial version of the protocol.
			 */
			ValueSet []SymbolKind `json:"valueSet,omitempty"`
		} `json:"symbolKind,omitempty"`
	} `json:"symbol,omitempty"`

<<<<<<< HEAD
	/** ExecuteCommand defined:
=======
	/*ExecuteCommand defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `workspace/executeCommand` request.
	 */
	ExecuteCommand *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Execute command supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"executeCommand,omitempty"`
}

<<<<<<< HEAD
// TextDocumentClientCapabilities is:
/**
=======
/*TextDocumentClientCapabilities defined:
>>>>>>> v0.0.4
 * Text document specific client capabilities.
 */
type TextDocumentClientCapabilities struct {

<<<<<<< HEAD
	/** Synchronization defined:
=======
	/*Synchronization defined:
>>>>>>> v0.0.4
	 * Defines which synchronization capabilities the client supports.
	 */
	Synchronization *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether text document synchronization supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

<<<<<<< HEAD
		/** WillSave defined:
=======
		/*WillSave defined:
>>>>>>> v0.0.4
		 * The client supports sending will save notifications.
		 */
		WillSave bool `json:"willSave,omitempty"`

<<<<<<< HEAD
		/** WillSaveWaitUntil defined:
=======
		/*WillSaveWaitUntil defined:
>>>>>>> v0.0.4
		 * The client supports sending a will save request and
		 * waits for a response providing text edits which will
		 * be applied to the document before it is saved.
		 */
		WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

<<<<<<< HEAD
		/** DidSave defined:
=======
		/*DidSave defined:
>>>>>>> v0.0.4
		 * The client supports did save notifications.
		 */
		DidSave bool `json:"didSave,omitempty"`
	} `json:"synchronization,omitempty"`

<<<<<<< HEAD
	/** Completion defined:
=======
	/*Completion defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/completion`
	 */
	Completion *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether completion supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

<<<<<<< HEAD
		/** CompletionItem defined:
=======
		/*CompletionItem defined:
>>>>>>> v0.0.4
		 * The client supports the following `CompletionItem` specific
		 * capabilities.
		 */
		CompletionItem *struct {

<<<<<<< HEAD
			/** SnippetSupport defined:
=======
			/*SnippetSupport defined:
>>>>>>> v0.0.4
			 * Client supports snippets as insert text.
			 *
			 * A snippet can define tab stops and placeholders with `$1`, `$2`
			 * and `${3:foo}`. `$0` defines the final tab stop, it defaults to
			 * the end of the snippet. Placeholders with equal identifiers are linked,
			 * that is typing in one will update others too.
			 */
			SnippetSupport bool `json:"snippetSupport,omitempty"`

<<<<<<< HEAD
			/** CommitCharactersSupport defined:
=======
			/*CommitCharactersSupport defined:
>>>>>>> v0.0.4
			 * Client supports commit characters on a completion item.
			 */
			CommitCharactersSupport bool `json:"commitCharactersSupport,omitempty"`

<<<<<<< HEAD
			/** DocumentationFormat defined:
=======
			/*DocumentationFormat defined:
>>>>>>> v0.0.4
			 * Client supports the follow content formats for the documentation
			 * property. The order describes the preferred format of the client.
			 */
			DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

<<<<<<< HEAD
			/** DeprecatedSupport defined:
=======
			/*DeprecatedSupport defined:
>>>>>>> v0.0.4
			 * Client supports the deprecated property on a completion item.
			 */
			DeprecatedSupport bool `json:"deprecatedSupport,omitempty"`

<<<<<<< HEAD
			/** PreselectSupport defined:
=======
			/*PreselectSupport defined:
>>>>>>> v0.0.4
			 * Client supports the preselect property on a completion item.
			 */
			PreselectSupport bool `json:"preselectSupport,omitempty"`
		} `json:"completionItem,omitempty"`

		// CompletionItemKind is
		CompletionItemKind *struct {

<<<<<<< HEAD
			/** ValueSet defined:
=======
			/*ValueSet defined:
>>>>>>> v0.0.4
			 * The completion item kind values the client supports. When this
			 * property exists the client also guarantees that it will
			 * handle values outside its set gracefully and falls back
			 * to a default value when unknown.
			 *
			 * If this property is not present the client only supports
			 * the completion items kinds from `Text` to `Reference` as defined in
			 * the initial version of the protocol.
			 */
			ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
		} `json:"completionItemKind,omitempty"`

<<<<<<< HEAD
		/** ContextSupport defined:
=======
		/*ContextSupport defined:
>>>>>>> v0.0.4
		 * The client supports to send additional context information for a
		 * `textDocument/completion` requestion.
		 */
		ContextSupport bool `json:"contextSupport,omitempty"`
	} `json:"completion,omitempty"`

<<<<<<< HEAD
	/** Hover defined:
=======
	/*Hover defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/hover`
	 */
	Hover *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether hover supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

<<<<<<< HEAD
		/** ContentFormat defined:
=======
		/*ContentFormat defined:
>>>>>>> v0.0.4
		 * Client supports the follow content formats for the content
		 * property. The order describes the preferred format of the client.
		 */
		ContentFormat []MarkupKind `json:"contentFormat,omitempty"`
	} `json:"hover,omitempty"`

<<<<<<< HEAD
	/** SignatureHelp defined:
=======
	/*SignatureHelp defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/signatureHelp`
	 */
	SignatureHelp *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether signature help supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

<<<<<<< HEAD
		/** SignatureInformation defined:
=======
		/*SignatureInformation defined:
>>>>>>> v0.0.4
		 * The client supports the following `SignatureInformation`
		 * specific properties.
		 */
		SignatureInformation *struct {

<<<<<<< HEAD
			/** DocumentationFormat defined:
=======
			/*DocumentationFormat defined:
>>>>>>> v0.0.4
			 * Client supports the follow content formats for the documentation
			 * property. The order describes the preferred format of the client.
			 */
			DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

<<<<<<< HEAD
			/** ParameterInformation defined:
=======
			/*ParameterInformation defined:
>>>>>>> v0.0.4
			 * Client capabilities specific to parameter information.
			 */
			ParameterInformation *struct {

<<<<<<< HEAD
				/** LabelOffsetSupport defined:
=======
				/*LabelOffsetSupport defined:
>>>>>>> v0.0.4
				 * The client supports processing label offsets instead of a
				 * simple label string.
				 */
				LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
			} `json:"parameterInformation,omitempty"`
		} `json:"signatureInformation,omitempty"`
	} `json:"signatureHelp,omitempty"`

<<<<<<< HEAD
	/** References defined:
=======
	/*References defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/references`
	 */
	References *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether references supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"references,omitempty"`

<<<<<<< HEAD
	/** DocumentHighlight defined:
=======
	/*DocumentHighlight defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/documentHighlight`
	 */
	DocumentHighlight *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether document highlight supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"documentHighlight,omitempty"`

<<<<<<< HEAD
	/** DocumentSymbol defined:
=======
	/*DocumentSymbol defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/documentSymbol`
	 */
	DocumentSymbol *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether document symbol supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

<<<<<<< HEAD
		/** SymbolKind defined:
=======
		/*SymbolKind defined:
>>>>>>> v0.0.4
		 * Specific capabilities for the `SymbolKind`.
		 */
		SymbolKind *struct {

<<<<<<< HEAD
			/** ValueSet defined:
=======
			/*ValueSet defined:
>>>>>>> v0.0.4
			 * The symbol kind values the client supports. When this
			 * property exists the client also guarantees that it will
			 * handle values outside its set gracefully and falls back
			 * to a default value when unknown.
			 *
			 * If this property is not present the client only supports
			 * the symbol kinds from `File` to `Array` as defined in
			 * the initial version of the protocol.
			 */
			ValueSet []SymbolKind `json:"valueSet,omitempty"`
		} `json:"symbolKind,omitempty"`

<<<<<<< HEAD
		/** HierarchicalDocumentSymbolSupport defined:
=======
		/*HierarchicalDocumentSymbolSupport defined:
>>>>>>> v0.0.4
		 * The client support hierarchical document symbols.
		 */
		HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`
	} `json:"documentSymbol,omitempty"`

<<<<<<< HEAD
	/** Formatting defined:
=======
	/*Formatting defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/formatting`
	 */
	Formatting *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether formatting supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"formatting,omitempty"`

<<<<<<< HEAD
	/** RangeFormatting defined:
=======
	/*RangeFormatting defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/rangeFormatting`
	 */
	RangeFormatting *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether range formatting supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"rangeFormatting,omitempty"`

<<<<<<< HEAD
	/** OnTypeFormatting defined:
=======
	/*OnTypeFormatting defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/onTypeFormatting`
	 */
	OnTypeFormatting *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether on type formatting supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"onTypeFormatting,omitempty"`

<<<<<<< HEAD
	/** Definition defined:
=======
	/*Definition defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/definition`
	 */
	Definition *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether definition supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

<<<<<<< HEAD
		/** LinkSupport defined:
=======
		/*LinkSupport defined:
>>>>>>> v0.0.4
		 * The client supports additional metadata in the form of definition links.
		 */
		LinkSupport bool `json:"linkSupport,omitempty"`
	} `json:"definition,omitempty"`

<<<<<<< HEAD
	/** CodeAction defined:
=======
	/*CodeAction defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/codeAction`
	 */
	CodeAction *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
=======
		/*DynamicRegistration defined:
>>>>>>> v0.0.4
		 * Whether code action supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

<<<<<<< HEAD
		/** CodeActionLiteralSupport defined:
=======
		/*CodeActionLiteralSupport defined:
>>>>>>> v0.0.4
		 * The client support code action literals as a valid
		 * response of the `textDocument/codeAction` request.
		 */
		CodeActionLiteralSupport *struct {

<<<<<<< HEAD
			/** CodeActionKind defined:
=======
			/*CodeActionKind defined:
>>>>>>> v0.0.4
			 * The code action kind is support with the following value
			 * set.
			 */
			CodeActionKind struct {

<<<<<<< HEAD
				/** ValueSet defined:
=======
				/*ValueSet defined:
>>>>>>> v0.0.4
				 * The code action kind values the client supports. When this
				 * property exists the client also guarantees that it will
				 * handle values outside its set gracefully and falls back
				 * to a default value when unknown.
				 */
				ValueSet []CodeActionKind `json:"valueSet"`
			} `json:"codeActionKind"`
		} `json:"codeActionLiteralSupport,omitempty"`
	} `json:"codeAction,omitempty"`

<<<<<<< HEAD
	/** CodeLens defined:
=======
	/*CodeLens defined:
>>>>>>> v0.0.4
	 * Capabilities specific to the `textDocument/codeLens`
	 */
	CodeLens *struct {

<<<<<<< HEAD
		/** DynamicRegistration defined:
		 * Whether code lens supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"codeLens,omitempty"`

	/** DocumentLink defined:
	 * Capabilities specific to the `textDocument/documentLink`
	 */
	DocumentLink *struct {

		/** DynamicRegistration defined:
		 * Whether document link supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"documentLink,omitempty"`

	/** Rename defined:
	 * Capabilities specific to the `textDocument/rename`
	 */
	Rename *struct {

		/** DynamicRegistration defined:
		 * Whether rename supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		/** PrepareSupport defined:
		 * Client supports testing for validity of rename operations
		 * before execution.
		 */
		PrepareSupport bool `json:"prepareSupport,omitempty"`
	} `json:"rename,omitempty"`

	/** PublishDiagnostics defined:
	 * Capabilities specific to `textDocument/publishDiagnostics`.
	 */
	PublishDiagnostics *struct {

		/** RelatedInformation defined:
		 * Whether the clients accepts diagnostics with related information.
		 */
		RelatedInformation bool `json:"relatedInformation,omitempty"`

		/** TagSupport defined:
		 * Client supports the tag property to provide meta data about a diagnostic.
		 */
		TagSupport bool `json:"tagSupport,omitempty"`
	} `json:"publishDiagnostics,omitempty"`
}

// InnerClientCapabilities is:
/**
 * Defines the capabilities provided by the client.
 */
type InnerClientCapabilities struct {

	/** Workspace defined:
	 * Workspace specific client capabilities.
	 */
	Workspace *WorkspaceClientCapabilities `json:"workspace,omitempty"`

	/** TextDocument defined:
	 * Text document specific client capabilities.
	 */
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`

	/** Experimental defined:
=======
		/*DynamicRegistration defined:
		 * Whether code lens supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"codeLens,omitempty"`

	/*DocumentLink defined:
	 * Capabilities specific to the `textDocument/documentLink`
	 */
	DocumentLink *struct {

		/*DynamicRegistration defined:
		 * Whether document link supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"documentLink,omitempty"`

	/*Rename defined:
	 * Capabilities specific to the `textDocument/rename`
	 */
	Rename *struct {

		/*DynamicRegistration defined:
		 * Whether rename supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		/*PrepareSupport defined:
		 * Client supports testing for validity of rename operations
		 * before execution.
		 */
		PrepareSupport bool `json:"prepareSupport,omitempty"`
	} `json:"rename,omitempty"`

	/*PublishDiagnostics defined:
	 * Capabilities specific to `textDocument/publishDiagnostics`.
	 */
	PublishDiagnostics *struct {

		/*RelatedInformation defined:
		 * Whether the clients accepts diagnostics with related information.
		 */
		RelatedInformation bool `json:"relatedInformation,omitempty"`

		/*TagSupport defined:
		 * Client supports the tag property to provide meta data about a diagnostic.
		 */
		TagSupport bool `json:"tagSupport,omitempty"`
	} `json:"publishDiagnostics,omitempty"`
}

// ClientCapabilities is
type ClientCapabilities struct {

	/*Workspace defined:
	 * Workspace specific client capabilities.
	 */
	Workspace struct {

		/*ApplyEdit defined:
		 * The client supports applying batch edits
		 * to the workspace by supporting the request
		 * 'workspace/applyEdit'
		 */
		ApplyEdit bool `json:"applyEdit,omitempty"`

		/*WorkspaceEdit defined:
		 * Capabilities specific to `WorkspaceEdit`s
		 */
		WorkspaceEdit struct {

			/*DocumentChanges defined:
			 * The client supports versioned document changes in `WorkspaceEdit`s
			 */
			DocumentChanges bool `json:"documentChanges,omitempty"`

			/*ResourceOperations defined:
			 * The resource operations the client supports. Clients should at least
			 * support 'create', 'rename' and 'delete' files and folders.
			 */
			ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitempty"`

			/*FailureHandling defined:
			 * The failure handling strategy of a client if applying the workspace edit
			 * failes.
			 */
			FailureHandling FailureHandlingKind `json:"failureHandling,omitempty"`
		} `json:"workspaceEdit,omitempty"`

		/*DidChangeConfiguration defined:
		 * Capabilities specific to the `workspace/didChangeConfiguration` notification.
		 */
		DidChangeConfiguration struct {

			/*DynamicRegistration defined:
			 * Did change configuration notification supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"didChangeConfiguration,omitempty"`

		/*DidChangeWatchedFiles defined:
		 * Capabilities specific to the `workspace/didChangeWatchedFiles` notification.
		 */
		DidChangeWatchedFiles struct {

			/*DynamicRegistration defined:
			 * Did change watched files notification supports dynamic registration. Please note
			 * that the current protocol doesn't support static configuration for file changes
			 * from the server side.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"didChangeWatchedFiles,omitempty"`

		/*Symbol defined:
		 * Capabilities specific to the `workspace/symbol` request.
		 */
		Symbol struct {

			/*DynamicRegistration defined:
			 * Symbol request supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*SymbolKind defined:
			 * Specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
			 */
			SymbolKind struct {

				/*ValueSet defined:
				 * The symbol kind values the client supports. When this
				 * property exists the client also guarantees that it will
				 * handle values outside its set gracefully and falls back
				 * to a default value when unknown.
				 *
				 * If this property is not present the client only supports
				 * the symbol kinds from `File` to `Array` as defined in
				 * the initial version of the protocol.
				 */
				ValueSet []SymbolKind `json:"valueSet,omitempty"`
			} `json:"symbolKind,omitempty"`
		} `json:"symbol,omitempty"`

		/*ExecuteCommand defined:
		 * Capabilities specific to the `workspace/executeCommand` request.
		 */
		ExecuteCommand struct {

			/*DynamicRegistration defined:
			 * Execute command supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"executeCommand,omitempty"`

		/*WorkspaceFolders defined:
		 * The client has support for workspace folders
		 */
		WorkspaceFolders bool `json:"workspaceFolders,omitempty"`

		/*Configuration defined:
		* The client supports `workspace/configuration` requests.
		 */
		Configuration bool `json:"configuration,omitempty"`
	} `json:"workspace,omitempty"`

	/*TextDocument defined:
	 * Text document specific client capabilities.
	 */
	TextDocument struct {

		/*Synchronization defined:
		 * Defines which synchronization capabilities the client supports.
		 */
		Synchronization struct {

			/*DynamicRegistration defined:
			 * Whether text document synchronization supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*WillSave defined:
			 * The client supports sending will save notifications.
			 */
			WillSave bool `json:"willSave,omitempty"`

			/*WillSaveWaitUntil defined:
			 * The client supports sending a will save request and
			 * waits for a response providing text edits which will
			 * be applied to the document before it is saved.
			 */
			WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

			/*DidSave defined:
			 * The client supports did save notifications.
			 */
			DidSave bool `json:"didSave,omitempty"`
		} `json:"synchronization,omitempty"`

		/*Completion defined:
		 * Capabilities specific to the `textDocument/completion`
		 */
		Completion struct {

			/*DynamicRegistration defined:
			 * Whether completion supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*CompletionItem defined:
			 * The client supports the following `CompletionItem` specific
			 * capabilities.
			 */
			CompletionItem struct {

				/*SnippetSupport defined:
				 * Client supports snippets as insert text.
				 *
				 * A snippet can define tab stops and placeholders with `$1`, `$2`
				 * and `${3:foo}`. `$0` defines the final tab stop, it defaults to
				 * the end of the snippet. Placeholders with equal identifiers are linked,
				 * that is typing in one will update others too.
				 */
				SnippetSupport bool `json:"snippetSupport,omitempty"`

				/*CommitCharactersSupport defined:
				 * Client supports commit characters on a completion item.
				 */
				CommitCharactersSupport bool `json:"commitCharactersSupport,omitempty"`

				/*DocumentationFormat defined:
				 * Client supports the follow content formats for the documentation
				 * property. The order describes the preferred format of the client.
				 */
				DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

				/*DeprecatedSupport defined:
				 * Client supports the deprecated property on a completion item.
				 */
				DeprecatedSupport bool `json:"deprecatedSupport,omitempty"`

				/*PreselectSupport defined:
				 * Client supports the preselect property on a completion item.
				 */
				PreselectSupport bool `json:"preselectSupport,omitempty"`
			} `json:"completionItem,omitempty"`

			// CompletionItemKind is
			CompletionItemKind struct {

				/*ValueSet defined:
				 * The completion item kind values the client supports. When this
				 * property exists the client also guarantees that it will
				 * handle values outside its set gracefully and falls back
				 * to a default value when unknown.
				 *
				 * If this property is not present the client only supports
				 * the completion items kinds from `Text` to `Reference` as defined in
				 * the initial version of the protocol.
				 */
				ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
			} `json:"completionItemKind,omitempty"`

			/*ContextSupport defined:
			 * The client supports to send additional context information for a
			 * `textDocument/completion` requestion.
			 */
			ContextSupport bool `json:"contextSupport,omitempty"`
		} `json:"completion,omitempty"`

		/*Hover defined:
		 * Capabilities specific to the `textDocument/hover`
		 */
		Hover struct {

			/*DynamicRegistration defined:
			 * Whether hover supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*ContentFormat defined:
			 * Client supports the follow content formats for the content
			 * property. The order describes the preferred format of the client.
			 */
			ContentFormat []MarkupKind `json:"contentFormat,omitempty"`
		} `json:"hover,omitempty"`

		/*SignatureHelp defined:
		 * Capabilities specific to the `textDocument/signatureHelp`
		 */
		SignatureHelp struct {

			/*DynamicRegistration defined:
			 * Whether signature help supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*SignatureInformation defined:
			 * The client supports the following `SignatureInformation`
			 * specific properties.
			 */
			SignatureInformation struct {

				/*DocumentationFormat defined:
				 * Client supports the follow content formats for the documentation
				 * property. The order describes the preferred format of the client.
				 */
				DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

				/*ParameterInformation defined:
				 * Client capabilities specific to parameter information.
				 */
				ParameterInformation struct {

					/*LabelOffsetSupport defined:
					 * The client supports processing label offsets instead of a
					 * simple label string.
					 */
					LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
				} `json:"parameterInformation,omitempty"`
			} `json:"signatureInformation,omitempty"`
		} `json:"signatureHelp,omitempty"`

		/*References defined:
		 * Capabilities specific to the `textDocument/references`
		 */
		References struct {

			/*DynamicRegistration defined:
			 * Whether references supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"references,omitempty"`

		/*DocumentHighlight defined:
		 * Capabilities specific to the `textDocument/documentHighlight`
		 */
		DocumentHighlight struct {

			/*DynamicRegistration defined:
			 * Whether document highlight supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"documentHighlight,omitempty"`

		/*DocumentSymbol defined:
		 * Capabilities specific to the `textDocument/documentSymbol`
		 */
		DocumentSymbol struct {

			/*DynamicRegistration defined:
			 * Whether document symbol supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*SymbolKind defined:
			 * Specific capabilities for the `SymbolKind`.
			 */
			SymbolKind struct {

				/*ValueSet defined:
				 * The symbol kind values the client supports. When this
				 * property exists the client also guarantees that it will
				 * handle values outside its set gracefully and falls back
				 * to a default value when unknown.
				 *
				 * If this property is not present the client only supports
				 * the symbol kinds from `File` to `Array` as defined in
				 * the initial version of the protocol.
				 */
				ValueSet []SymbolKind `json:"valueSet,omitempty"`
			} `json:"symbolKind,omitempty"`

			/*HierarchicalDocumentSymbolSupport defined:
			 * The client support hierarchical document symbols.
			 */
			HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`
		} `json:"documentSymbol,omitempty"`

		/*Formatting defined:
		 * Capabilities specific to the `textDocument/formatting`
		 */
		Formatting struct {

			/*DynamicRegistration defined:
			 * Whether formatting supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"formatting,omitempty"`

		/*RangeFormatting defined:
		 * Capabilities specific to the `textDocument/rangeFormatting`
		 */
		RangeFormatting struct {

			/*DynamicRegistration defined:
			 * Whether range formatting supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"rangeFormatting,omitempty"`

		/*OnTypeFormatting defined:
		 * Capabilities specific to the `textDocument/onTypeFormatting`
		 */
		OnTypeFormatting struct {

			/*DynamicRegistration defined:
			 * Whether on type formatting supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"onTypeFormatting,omitempty"`

		/*Definition defined:
		 * Capabilities specific to the `textDocument/definition`
		 */
		Definition struct {

			/*DynamicRegistration defined:
			 * Whether definition supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*LinkSupport defined:
			 * The client supports additional metadata in the form of definition links.
			 */
			LinkSupport bool `json:"linkSupport,omitempty"`
		} `json:"definition,omitempty"`

		/*CodeAction defined:
		 * Capabilities specific to the `textDocument/codeAction`
		 */
		CodeAction struct {

			/*DynamicRegistration defined:
			 * Whether code action supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*CodeActionLiteralSupport defined:
			 * The client support code action literals as a valid
			 * response of the `textDocument/codeAction` request.
			 */
			CodeActionLiteralSupport struct {

				/*CodeActionKind defined:
				 * The code action kind is support with the following value
				 * set.
				 */
				CodeActionKind struct {

					/*ValueSet defined:
					 * The code action kind values the client supports. When this
					 * property exists the client also guarantees that it will
					 * handle values outside its set gracefully and falls back
					 * to a default value when unknown.
					 */
					ValueSet []CodeActionKind `json:"valueSet"`
				} `json:"codeActionKind"`
			} `json:"codeActionLiteralSupport,omitempty"`
		} `json:"codeAction,omitempty"`

		/*CodeLens defined:
		 * Capabilities specific to the `textDocument/codeLens`
		 */
		CodeLens struct {

			/*DynamicRegistration defined:
			 * Whether code lens supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"codeLens,omitempty"`

		/*DocumentLink defined:
		 * Capabilities specific to the `textDocument/documentLink`
		 */
		DocumentLink struct {

			/*DynamicRegistration defined:
			 * Whether document link supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"documentLink,omitempty"`

		/*Rename defined:
		 * Capabilities specific to the `textDocument/rename`
		 */
		Rename struct {

			/*DynamicRegistration defined:
			 * Whether rename supports dynamic registration.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*PrepareSupport defined:
			 * Client supports testing for validity of rename operations
			 * before execution.
			 */
			PrepareSupport bool `json:"prepareSupport,omitempty"`
		} `json:"rename,omitempty"`

		/*PublishDiagnostics defined:
		 * Capabilities specific to `textDocument/publishDiagnostics`.
		 */
		PublishDiagnostics struct {

			/*RelatedInformation defined:
			 * Whether the clients accepts diagnostics with related information.
			 */
			RelatedInformation bool `json:"relatedInformation,omitempty"`

			/*TagSupport defined:
			 * Client supports the tag property to provide meta data about a diagnostic.
			 */
			TagSupport bool `json:"tagSupport,omitempty"`
		} `json:"publishDiagnostics,omitempty"`

		/*Implementation defined:
		 * Capabilities specific to the `textDocument/implementation`
		 */
		Implementation struct {

			/*DynamicRegistration defined:
			 * Whether implementation supports dynamic registration. If this is set to `true`
			 * the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*LinkSupport defined:
			 * The client supports additional metadata in the form of definition links.
			 */
			LinkSupport bool `json:"linkSupport,omitempty"`
		} `json:"implementation,omitempty"`

		/*TypeDefinition defined:
		 * Capabilities specific to the `textDocument/typeDefinition`
		 */
		TypeDefinition struct {

			/*DynamicRegistration defined:
			 * Whether implementation supports dynamic registration. If this is set to `true`
			 * the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*LinkSupport defined:
			 * The client supports additional metadata in the form of definition links.
			 */
			LinkSupport bool `json:"linkSupport,omitempty"`
		} `json:"typeDefinition,omitempty"`

		/*ColorProvider defined:
		 * Capabilities specific to the colorProvider
		 */
		ColorProvider struct {

			/*DynamicRegistration defined:
			 * Whether implementation supports dynamic registration. If this is set to `true`
			 * the client supports the new `(ColorProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"colorProvider,omitempty"`

		/*FoldingRange defined:
		 * Capabilities specific to `textDocument/foldingRange` requests
		 */
		FoldingRange struct {

			/*DynamicRegistration defined:
			 * Whether implementation supports dynamic registration for folding range providers. If this is set to `true`
			 * the client supports the new `(FoldingRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*RangeLimit defined:
			 * The maximum number of folding ranges that the client prefers to receive per document. The value serves as a
			 * hint, servers are free to follow the limit.
			 */
			RangeLimit float64 `json:"rangeLimit,omitempty"`

			/*LineFoldingOnly defined:
			 * If set, the client signals that it only supports folding complete lines. If set, client will
			 * ignore specified `startCharacter` and `endCharacter` properties in a FoldingRange.
			 */
			LineFoldingOnly bool `json:"lineFoldingOnly,omitempty"`
		} `json:"foldingRange,omitempty"`

		/*Declaration defined:
		 * Capabilities specific to the `textDocument/declaration`
		 */
		Declaration struct {

			/*DynamicRegistration defined:
			 * Whether declaration supports dynamic registration. If this is set to `true`
			 * the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			/*LinkSupport defined:
			 * The client supports additional metadata in the form of declaration links.
			 */
			LinkSupport bool `json:"linkSupport,omitempty"`
		} `json:"declaration,omitempty"`

		/*SelectionRange defined:
		 * Capabilities specific to `textDocument/selectionRange` requests
		 */
		SelectionRange struct {

			/*DynamicRegistration defined:
			 * Whether implementation supports dynamic registration for selection range providers. If this is set to `true`
			 * the client supports the new `(SelectionRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)`
			 * return value for the corresponding server capability as well.
			 */
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		} `json:"selectionRange,omitempty"`
	} `json:"textDocument,omitempty"`

	/*Experimental defined:
>>>>>>> v0.0.4
	 * Experimental client capabilities.
	 */
	Experimental interface{} `json:"experimental,omitempty"`
}

<<<<<<< HEAD
// TODO(rstambler): Remove this when golang.org/issue/31090 is resolved.
type ClientCapabilities map[string]interface{}

// clientCapabilities is:
type clientCapabilities struct {
	InnerClientCapabilities
	ImplementationClientCapabilities
	TypeDefinitionClientCapabilities
	WorkspaceFoldersClientCapabilities
	ConfigurationClientCapabilities
	ColorClientCapabilities
	FoldingRangeClientCapabilities
	DeclarationClientCapabilities
	SelectionRangeClientCapabilities
}

// StaticRegistrationOptions is:
/**
=======
/*StaticRegistrationOptions defined:
>>>>>>> v0.0.4
 * Static registration options to be returned in the initialize
 * request.
 */
type StaticRegistrationOptions struct {

<<<<<<< HEAD
	/** ID defined:
=======
	/*ID defined:
>>>>>>> v0.0.4
	 * The id used to register the request. The id can be used to deregister
	 * the request again. See also Registration#id.
	 */
	ID string `json:"id,omitempty"`
}

<<<<<<< HEAD
// TextDocumentRegistrationOptions is:
/**
=======
/*TextDocumentRegistrationOptions defined:
>>>>>>> v0.0.4
 * General text document registration options.
 */
type TextDocumentRegistrationOptions struct {

<<<<<<< HEAD
	/** DocumentSelector defined:
=======
	/*DocumentSelector defined:
>>>>>>> v0.0.4
	 * A document selector to identify the scope of the registration. If set to null
	 * the document selector provided on the client side will be used.
	 */
	DocumentSelector DocumentSelector `json:"documentSelector"`
}

<<<<<<< HEAD
// CompletionOptions is:
/**
=======
/*CompletionOptions defined:
>>>>>>> v0.0.4
 * Completion options.
 */
type CompletionOptions struct {

<<<<<<< HEAD
	/** TriggerCharacters defined:
=======
	/*TriggerCharacters defined:
>>>>>>> v0.0.4
	 * Most tools trigger completion request automatically without explicitly requesting
	 * it using a keyboard shortcut (e.g. Ctrl+Space). Typically they do so when the user
	 * starts to type an identifier. For example if the user types `c` in a JavaScript file
	 * code complete will automatically pop up present `console` besides others as a
	 * completion item. Characters that make up identifiers don't need to be listed here.
	 *
	 * If code complete should automatically be trigger on characters not being valid inside
	 * an identifier (for example `.` in JavaScript) list them in `triggerCharacters`.
	 */
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

<<<<<<< HEAD
	/** AllCommitCharacters defined:
=======
	/*AllCommitCharacters defined:
>>>>>>> v0.0.4
	 * The list of all possible characters that commit a completion. This field can be used
	 * if clients don't support individual commmit characters per completion item. See
	 * `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport`
	 */
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`

<<<<<<< HEAD
	/** ResolveProvider defined:
=======
	/*ResolveProvider defined:
>>>>>>> v0.0.4
	 * The server provides support to resolve additional
	 * information for a completion item.
	 */
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

<<<<<<< HEAD
// SignatureHelpOptions is:
/**
=======
/*SignatureHelpOptions defined:
>>>>>>> v0.0.4
 * Signature help options.
 */
type SignatureHelpOptions struct {

<<<<<<< HEAD
	/** TriggerCharacters defined:
=======
	/*TriggerCharacters defined:
>>>>>>> v0.0.4
	 * The characters that trigger signature help
	 * automatically.
	 */
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
}

<<<<<<< HEAD
// CodeActionOptions is:
/**
=======
/*CodeActionOptions defined:
>>>>>>> v0.0.4
 * Code Action options.
 */
type CodeActionOptions struct {

<<<<<<< HEAD
	/** CodeActionKinds defined:
=======
	/*CodeActionKinds defined:
>>>>>>> v0.0.4
	 * CodeActionKinds that this server may return.
	 *
	 * The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server
	 * may list out every specific kind they provide.
	 */
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`
}

<<<<<<< HEAD
// CodeLensOptions is:
/**
=======
/*CodeLensOptions defined:
>>>>>>> v0.0.4
 * Code Lens options.
 */
type CodeLensOptions struct {

<<<<<<< HEAD
	/** ResolveProvider defined:
=======
	/*ResolveProvider defined:
>>>>>>> v0.0.4
	 * Code lens has a resolve provider as well.
	 */
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

<<<<<<< HEAD
// DocumentOnTypeFormattingOptions is:
/**
=======
/*DocumentOnTypeFormattingOptions defined:
>>>>>>> v0.0.4
 * Format document on type options
 */
type DocumentOnTypeFormattingOptions struct {

<<<<<<< HEAD
	/** FirstTriggerCharacter defined:
=======
	/*FirstTriggerCharacter defined:
>>>>>>> v0.0.4
	 * A character on which formatting should be triggered, like `}`.
	 */
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

<<<<<<< HEAD
	/** MoreTriggerCharacter defined:
=======
	/*MoreTriggerCharacter defined:
>>>>>>> v0.0.4
	 * More trigger characters.
	 */
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
}

<<<<<<< HEAD
// RenameOptions is:
/**
=======
/*RenameOptions defined:
>>>>>>> v0.0.4
 * Rename options
 */
type RenameOptions struct {

<<<<<<< HEAD
	/** PrepareProvider defined:
=======
	/*PrepareProvider defined:
>>>>>>> v0.0.4
	 * Renames should be checked and tested before being executed.
	 */
	PrepareProvider bool `json:"prepareProvider,omitempty"`
}

<<<<<<< HEAD
// DocumentLinkOptions is:
/**
=======
/*DocumentLinkOptions defined:
>>>>>>> v0.0.4
 * Document link options
 */
type DocumentLinkOptions struct {

<<<<<<< HEAD
	/** ResolveProvider defined:
=======
	/*ResolveProvider defined:
>>>>>>> v0.0.4
	 * Document links have a resolve provider as well.
	 */
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

<<<<<<< HEAD
// ExecuteCommandOptions is:
/**
=======
/*ExecuteCommandOptions defined:
>>>>>>> v0.0.4
 * Execute command options.
 */
type ExecuteCommandOptions struct {

<<<<<<< HEAD
	/** Commands defined:
=======
	/*Commands defined:
>>>>>>> v0.0.4
	 * The commands to be executed on the server
	 */
	Commands []string `json:"commands"`
}

<<<<<<< HEAD
// SaveOptions is:
/**
=======
/*SaveOptions defined:
>>>>>>> v0.0.4
 * Save options.
 */
type SaveOptions struct {

<<<<<<< HEAD
	/** IncludeText defined:
=======
	/*IncludeText defined:
>>>>>>> v0.0.4
	 * The client is supposed to include the content on save.
	 */
	IncludeText bool `json:"includeText,omitempty"`
}

<<<<<<< HEAD
// TextDocumentSyncOptions is:
type TextDocumentSyncOptions struct {

	/** OpenClose defined:
=======
// TextDocumentSyncOptions is
type TextDocumentSyncOptions struct {

	/*OpenClose defined:
>>>>>>> v0.0.4
	 * Open and close notifications are sent to the server.
	 */
	OpenClose bool `json:"openClose,omitempty"`

<<<<<<< HEAD
	/** Change defined:
=======
	/*Change defined:
>>>>>>> v0.0.4
	 * Change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full
	 * and TextDocumentSyncKind.Incremental.
	 */
	Change TextDocumentSyncKind `json:"change,omitempty"`

<<<<<<< HEAD
	/** WillSave defined:
=======
	/*WillSave defined:
>>>>>>> v0.0.4
	 * Will save notifications are sent to the server.
	 */
	WillSave bool `json:"willSave,omitempty"`

<<<<<<< HEAD
	/** WillSaveWaitUntil defined:
=======
	/*WillSaveWaitUntil defined:
>>>>>>> v0.0.4
	 * Will save wait until requests are sent to the server.
	 */
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

<<<<<<< HEAD
	/** Save defined:
=======
	/*Save defined:
>>>>>>> v0.0.4
	 * Save notifications are sent to the server.
	 */
	Save *SaveOptions `json:"save,omitempty"`
}

<<<<<<< HEAD
// InnerServerCapabilities is:
/**
 * Defines the capabilities provided by a language
 * server.
 */
type InnerServerCapabilities struct {

	/** TextDocumentSync defined:
=======
// ServerCapabilities is
type ServerCapabilities struct {

	/*TextDocumentSync defined:
>>>>>>> v0.0.4
	 * Defines how text documents are synced. Is either a detailed structure defining each notification or
	 * for backwards compatibility the TextDocumentSyncKind number.
	 */
	TextDocumentSync interface{} `json:"textDocumentSync,omitempty"` // TextDocumentSyncOptions | TextDocumentSyncKind

<<<<<<< HEAD
	/** HoverProvider defined:
=======
	/*HoverProvider defined:
>>>>>>> v0.0.4
	 * The server provides hover support.
	 */
	HoverProvider bool `json:"hoverProvider,omitempty"`

<<<<<<< HEAD
	/** CompletionProvider defined:
=======
	/*CompletionProvider defined:
>>>>>>> v0.0.4
	 * The server provides completion support.
	 */
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

<<<<<<< HEAD
	/** SignatureHelpProvider defined:
=======
	/*SignatureHelpProvider defined:
>>>>>>> v0.0.4
	 * The server provides signature help support.
	 */
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

<<<<<<< HEAD
	/** DefinitionProvider defined:
=======
	/*DefinitionProvider defined:
>>>>>>> v0.0.4
	 * The server provides goto definition support.
	 */
	DefinitionProvider bool `json:"definitionProvider,omitempty"`

<<<<<<< HEAD
	/** ReferencesProvider defined:
=======
	/*ReferencesProvider defined:
>>>>>>> v0.0.4
	 * The server provides find references support.
	 */
	ReferencesProvider bool `json:"referencesProvider,omitempty"`

<<<<<<< HEAD
	/** DocumentHighlightProvider defined:
=======
	/*DocumentHighlightProvider defined:
>>>>>>> v0.0.4
	 * The server provides document highlight support.
	 */
	DocumentHighlightProvider bool `json:"documentHighlightProvider,omitempty"`

<<<<<<< HEAD
	/** DocumentSymbolProvider defined:
=======
	/*DocumentSymbolProvider defined:
>>>>>>> v0.0.4
	 * The server provides document symbol support.
	 */
	DocumentSymbolProvider bool `json:"documentSymbolProvider,omitempty"`

<<<<<<< HEAD
	/** WorkspaceSymbolProvider defined:
=======
	/*WorkspaceSymbolProvider defined:
>>>>>>> v0.0.4
	 * The server provides workspace symbol support.
	 */
	WorkspaceSymbolProvider bool `json:"workspaceSymbolProvider,omitempty"`

<<<<<<< HEAD
	/** CodeActionProvider defined:
=======
	/*CodeActionProvider defined:
>>>>>>> v0.0.4
	 * The server provides code actions. CodeActionOptions may only be
	 * specified if the client states that it supports
	 * `codeActionLiteralSupport` in its initial `initialize` request.
	 */
	CodeActionProvider bool `json:"codeActionProvider,omitempty"` // boolean | CodeActionOptions

<<<<<<< HEAD
	/** CodeLensProvider defined:
=======
	/*CodeLensProvider defined:
>>>>>>> v0.0.4
	 * The server provides code lens.
	 */
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider,omitempty"`

<<<<<<< HEAD
	/** DocumentFormattingProvider defined:
=======
	/*DocumentFormattingProvider defined:
>>>>>>> v0.0.4
	 * The server provides document formatting.
	 */
	DocumentFormattingProvider bool `json:"documentFormattingProvider,omitempty"`

<<<<<<< HEAD
	/** DocumentRangeFormattingProvider defined:
=======
	/*DocumentRangeFormattingProvider defined:
>>>>>>> v0.0.4
	 * The server provides document range formatting.
	 */
	DocumentRangeFormattingProvider bool `json:"documentRangeFormattingProvider,omitempty"`

<<<<<<< HEAD
	/** DocumentOnTypeFormattingProvider defined:
=======
	/*DocumentOnTypeFormattingProvider defined:
>>>>>>> v0.0.4
	 * The server provides document formatting on typing.
	 */
	DocumentOnTypeFormattingProvider *struct {

<<<<<<< HEAD
		/** FirstTriggerCharacter defined:
=======
		/*FirstTriggerCharacter defined:
>>>>>>> v0.0.4
		 * A character on which formatting should be triggered, like `}`.
		 */
		FirstTriggerCharacter string `json:"firstTriggerCharacter"`

<<<<<<< HEAD
		/** MoreTriggerCharacter defined:
=======
		/*MoreTriggerCharacter defined:
>>>>>>> v0.0.4
		 * More trigger characters.
		 */
		MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
	} `json:"documentOnTypeFormattingProvider,omitempty"`

<<<<<<< HEAD
	/** RenameProvider defined:
=======
	/*RenameProvider defined:
>>>>>>> v0.0.4
	 * The server provides rename support. RenameOptions may only be
	 * specified if the client states that it supports
	 * `prepareSupport` in its initial `initialize` request.
	 */
<<<<<<< HEAD
	RenameProvider bool `json:"renameProvider,omitempty"` // boolean | RenameOptions

	/** DocumentLinkProvider defined:
=======
	RenameProvider *RenameOptions `json:"renameProvider,omitempty"` // boolean | RenameOptions

	/*DocumentLinkProvider defined:
>>>>>>> v0.0.4
	 * The server provides document link support.
	 */
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider,omitempty"`

<<<<<<< HEAD
	/** ExecuteCommandProvider defined:
=======
	/*ExecuteCommandProvider defined:
>>>>>>> v0.0.4
	 * The server provides execute command support.
	 */
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider,omitempty"`

<<<<<<< HEAD
	/** Experimental defined:
	 * Experimental server capabilities.
	 */
	Experimental interface{} `json:"experimental,omitempty"`
}

// ServerCapabilities is:
type ServerCapabilities struct {
	InnerServerCapabilities
	ImplementationServerCapabilities
	TypeDefinitionServerCapabilities
	WorkspaceFoldersServerCapabilities
	ColorServerCapabilities
	FoldingRangeServerCapabilities
	DeclarationServerCapabilities
	SelectionRangeServerCapabilities
}

// InnerInitializeParams is:
/**
 * The initialize parameters
 */
type InnerInitializeParams struct {

	/** ProcessID defined:
=======
	/*Experimental defined:
	 * Experimental server capabilities.
	 */
	Experimental interface{} `json:"experimental,omitempty"`

	/*ImplementationProvider defined:
	 * The server provides Goto Implementation support.
	 */
	ImplementationProvider bool `json:"implementationProvider,omitempty"` // boolean | (TextDocumentRegistrationOptions & StaticRegistrationOptions)

	/*TypeDefinitionProvider defined:
	 * The server provides Goto Type Definition support.
	 */
	TypeDefinitionProvider bool `json:"typeDefinitionProvider,omitempty"` // boolean | (TextDocumentRegistrationOptions & StaticRegistrationOptions)

	/*Workspace defined:
	 * The workspace server capabilities
	 */
	Workspace *struct {

		// WorkspaceFolders is
		WorkspaceFolders *struct {

			/*Supported defined:
			 * The Server has support for workspace folders
			 */
			Supported bool `json:"supported,omitempty"`

			/*ChangeNotifications defined:
			 * Whether the server wants to receive workspace folder
			 * change notifications.
			 *
			 * If a strings is provided the string is treated as a ID
			 * under which the notification is registed on the client
			 * side. The ID can be used to unregister for these events
			 * using the `client/unregisterCapability` request.
			 */
			ChangeNotifications string `json:"changeNotifications,omitempty"` // string | boolean
		} `json:"workspaceFolders,omitempty"`
	} `json:"workspace,omitempty"`

	/*ColorProvider defined:
	 * The server provides color provider support.
	 */
	ColorProvider bool `json:"colorProvider,omitempty"` // boolean | ColorProviderOptions | (ColorProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)

	/*FoldingRangeProvider defined:
	 * The server provides folding provider support.
	 */
	FoldingRangeProvider bool `json:"foldingRangeProvider,omitempty"` // boolean | FoldingRangeProviderOptions | (FoldingRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)

	/*DeclarationProvider defined:
	 * The server provides Goto Type Definition support.
	 */
	DeclarationProvider bool `json:"declarationProvider,omitempty"` // boolean | (TextDocumentRegistrationOptions & StaticRegistrationOptions)

	/*SelectionRangeProvider defined:
	 * The server provides selection range support.
	 */
	SelectionRangeProvider bool `json:"selectionRangeProvider,omitempty"` // boolean | SelectionRangeProviderOptions | (SelectionRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)
}

// InitializeParams is
type InitializeParams struct {

	/*ProcessID defined:
>>>>>>> v0.0.4
	 * The process Id of the parent process that started
	 * the server.
	 */
	ProcessID float64 `json:"processId"`

<<<<<<< HEAD
	/** RootPath defined:
=======
	/*RootPath defined:
>>>>>>> v0.0.4
	 * The rootPath of the workspace. Is null
	 * if no folder is open.
	 *
	 * @deprecated in favour of rootUri.
	 */
	RootPath string `json:"rootPath,omitempty"`

<<<<<<< HEAD
	/** RootURI defined:
=======
	/*RootURI defined:
>>>>>>> v0.0.4
	 * The rootUri of the workspace. Is null if no
	 * folder is open. If both `rootPath` and `rootUri` are set
	 * `rootUri` wins.
	 *
	 * @deprecated in favour of workspaceFolders.
	 */
	RootURI string `json:"rootUri"`

<<<<<<< HEAD
	/** Capabilities defined:
=======
	/*Capabilities defined:
>>>>>>> v0.0.4
	 * The capabilities provided by the client (editor or tool)
	 */
	Capabilities ClientCapabilities `json:"capabilities"`

<<<<<<< HEAD
	/** InitializationOptions defined:
=======
	/*InitializationOptions defined:
>>>>>>> v0.0.4
	 * User provided initialization options.
	 */
	InitializationOptions interface{} `json:"initializationOptions,omitempty"`

<<<<<<< HEAD
	/** Trace defined:
	 * The initial trace setting. If omitted trace is disabled ('off').
	 */
	Trace string `json:"trace,omitempty"` // 'off' | 'messages' | 'verbose'
}

// InitializeParams is:
type InitializeParams struct {
	InnerInitializeParams
	WorkspaceFoldersInitializeParams
}

// InitializeResult is:
/**
=======
	/*Trace defined:
	 * The initial trace setting. If omitted trace is disabled ('off').
	 */
	Trace string `json:"trace,omitempty"` // 'off' | 'messages' | 'verbose'

	/*WorkspaceFolders defined:
	 * The actual configured workspace folders.
	 */
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders"`
}

/*InitializeResult defined:
>>>>>>> v0.0.4
 * The result returned from an initialize request.
 */
type InitializeResult struct {

<<<<<<< HEAD
	/** Capabilities defined:
=======
	/*Capabilities defined:
>>>>>>> v0.0.4
	 * The capabilities the language server provides.
	 */
	Capabilities ServerCapabilities `json:"capabilities"`

<<<<<<< HEAD
	/** Custom defined:
=======
	/*Custom defined:
>>>>>>> v0.0.4
	 * Custom initialization results.
	 */
	Custom map[string]interface{} `json:"custom"` // [custom: string]: any;
}

<<<<<<< HEAD
// InitializedParams is:
type InitializedParams struct {
}

// DidChangeConfigurationRegistrationOptions is:
=======
// InitializedParams is
type InitializedParams struct {
}

// DidChangeConfigurationRegistrationOptions is
>>>>>>> v0.0.4
type DidChangeConfigurationRegistrationOptions struct {

	// Section is
	Section string `json:"section,omitempty"` // string | string[]
}

<<<<<<< HEAD
// DidChangeConfigurationParams is:
/**
=======
/*DidChangeConfigurationParams defined:
>>>>>>> v0.0.4
 * The parameters of a change configuration notification.
 */
type DidChangeConfigurationParams struct {

<<<<<<< HEAD
	/** Settings defined:
=======
	/*Settings defined:
>>>>>>> v0.0.4
	 * The actual changed settings
	 */
	Settings interface{} `json:"settings"`
}

<<<<<<< HEAD
// ShowMessageParams is:
/**
=======
/*ShowMessageParams defined:
>>>>>>> v0.0.4
 * The parameters of a notification message.
 */
type ShowMessageParams struct {

<<<<<<< HEAD
	/** Type defined:
=======
	/*Type defined:
>>>>>>> v0.0.4
	 * The message type. See {@link MessageType}
	 */
	Type MessageType `json:"type"`

<<<<<<< HEAD
	/** Message defined:
=======
	/*Message defined:
>>>>>>> v0.0.4
	 * The actual message
	 */
	Message string `json:"message"`
}

<<<<<<< HEAD
// MessageActionItem is:
type MessageActionItem struct {

	/** Title defined:
=======
// MessageActionItem is
type MessageActionItem struct {

	/*Title defined:
>>>>>>> v0.0.4
	 * A short title like 'Retry', 'Open Log' etc.
	 */
	Title string `json:"title"`
}

<<<<<<< HEAD
// ShowMessageRequestParams is:
type ShowMessageRequestParams struct {

	/** Type defined:
=======
// ShowMessageRequestParams is
type ShowMessageRequestParams struct {

	/*Type defined:
>>>>>>> v0.0.4
	 * The message type. See {@link MessageType}
	 */
	Type MessageType `json:"type"`

<<<<<<< HEAD
	/** Message defined:
=======
	/*Message defined:
>>>>>>> v0.0.4
	 * The actual message
	 */
	Message string `json:"message"`

<<<<<<< HEAD
	/** Actions defined:
=======
	/*Actions defined:
>>>>>>> v0.0.4
	 * The message action items to present.
	 */
	Actions []MessageActionItem `json:"actions,omitempty"`
}

<<<<<<< HEAD
// LogMessageParams is:
/**
=======
/*LogMessageParams defined:
>>>>>>> v0.0.4
 * The log message parameters.
 */
type LogMessageParams struct {

<<<<<<< HEAD
	/** Type defined:
=======
	/*Type defined:
>>>>>>> v0.0.4
	 * The message type. See {@link MessageType}
	 */
	Type MessageType `json:"type"`

<<<<<<< HEAD
	/** Message defined:
=======
	/*Message defined:
>>>>>>> v0.0.4
	 * The actual message
	 */
	Message string `json:"message"`
}

<<<<<<< HEAD
// DidOpenTextDocumentParams is:
/**
=======
/*DidOpenTextDocumentParams defined:
>>>>>>> v0.0.4
 * The parameters send in a open text document notification
 */
type DidOpenTextDocumentParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document that was opened.
	 */
	TextDocument TextDocumentItem `json:"textDocument"`
}

<<<<<<< HEAD
// DidChangeTextDocumentParams is:
/**
=======
/*DidChangeTextDocumentParams defined:
>>>>>>> v0.0.4
 * The change text document notification's parameters.
 */
type DidChangeTextDocumentParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document that did change. The version number points
	 * to the version after all provided content changes have
	 * been applied.
	 */
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** ContentChanges defined:
=======
	/*ContentChanges defined:
>>>>>>> v0.0.4
	 * The actual content changes. The content changes describe single state changes
	 * to the document. So if there are two content changes c1 and c2 for a document
	 * in state S then c1 move the document to S' and c2 to S''.
	 */
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

<<<<<<< HEAD
// TextDocumentChangeRegistrationOptions is:
/**
 * Describe options to be used when registered for text document change events.
 */
type TextDocumentChangeRegistrationOptions struct {
	TextDocumentRegistrationOptions

	/** SyncKind defined:
	 * How documents are synced to the server.
	 */
	SyncKind TextDocumentSyncKind `json:"syncKind"`
}

// DidCloseTextDocumentParams is:
/**
=======
/*TextDocumentChangeRegistrationOptions defined:
 * Describe options to be used when registered for text document change events.
 */
type TextDocumentChangeRegistrationOptions struct {

	/*SyncKind defined:
	 * How documents are synced to the server.
	 */
	SyncKind TextDocumentSyncKind `json:"syncKind"`
	TextDocumentRegistrationOptions
}

/*DidCloseTextDocumentParams defined:
>>>>>>> v0.0.4
 * The parameters send in a close text document notification
 */
type DidCloseTextDocumentParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document that was closed.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

<<<<<<< HEAD
// DidSaveTextDocumentParams is:
/**
=======
/*DidSaveTextDocumentParams defined:
>>>>>>> v0.0.4
 * The parameters send in a save text document notification
 */
type DidSaveTextDocumentParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document that was closed.
	 */
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Text defined:
=======
	/*Text defined:
>>>>>>> v0.0.4
	 * Optional the content when saved. Depends on the includeText value
	 * when the save notification was requested.
	 */
	Text string `json:"text,omitempty"`
}

<<<<<<< HEAD
// TextDocumentSaveRegistrationOptions is:
/**
=======
/*TextDocumentSaveRegistrationOptions defined:
>>>>>>> v0.0.4
 * Save registration options.
 */
type TextDocumentSaveRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SaveOptions
}

<<<<<<< HEAD
// WillSaveTextDocumentParams is:
/**
=======
/*WillSaveTextDocumentParams defined:
>>>>>>> v0.0.4
 * The parameters send in a will save text document notification.
 */
type WillSaveTextDocumentParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document that will be saved.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Reason defined:
=======
	/*Reason defined:
>>>>>>> v0.0.4
	 * The 'TextDocumentSaveReason'.
	 */
	Reason TextDocumentSaveReason `json:"reason"`
}

<<<<<<< HEAD
// DidChangeWatchedFilesParams is:
/**
=======
/*DidChangeWatchedFilesParams defined:
>>>>>>> v0.0.4
 * The watched files change notification's parameters.
 */
type DidChangeWatchedFilesParams struct {

<<<<<<< HEAD
	/** Changes defined:
=======
	/*Changes defined:
>>>>>>> v0.0.4
	 * The actual file events.
	 */
	Changes []FileEvent `json:"changes"`
}

<<<<<<< HEAD
// FileEvent is:
/**
=======
/*FileEvent defined:
>>>>>>> v0.0.4
 * An event describing a file change.
 */
type FileEvent struct {

<<<<<<< HEAD
	/** URI defined:
=======
	/*URI defined:
>>>>>>> v0.0.4
	 * The file's uri.
	 */
	URI string `json:"uri"`

<<<<<<< HEAD
	/** Type defined:
=======
	/*Type defined:
>>>>>>> v0.0.4
	 * The change type.
	 */
	Type FileChangeType `json:"type"`
}

<<<<<<< HEAD
// DidChangeWatchedFilesRegistrationOptions is:
/**
=======
/*DidChangeWatchedFilesRegistrationOptions defined:
>>>>>>> v0.0.4
 * Describe options to be used when registered for text document change events.
 */
type DidChangeWatchedFilesRegistrationOptions struct {

<<<<<<< HEAD
	/** Watchers defined:
=======
	/*Watchers defined:
>>>>>>> v0.0.4
	 * The watchers to register.
	 */
	Watchers []FileSystemWatcher `json:"watchers"`
}

<<<<<<< HEAD
// FileSystemWatcher is:
type FileSystemWatcher struct {

	/** GlobPattern defined:
=======
// FileSystemWatcher is
type FileSystemWatcher struct {

	/*GlobPattern defined:
>>>>>>> v0.0.4
	 * The  glob pattern to watch. Glob patterns can have the following syntax:
	 * - `*` to match one or more characters in a path segment
	 * - `?` to match on one character in a path segment
	 * - `**` to match any number of path segments, including none
	 * - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
	 * - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	 * - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
	 */
	GlobPattern string `json:"globPattern"`

<<<<<<< HEAD
	/** Kind defined:
=======
	/*Kind defined:
>>>>>>> v0.0.4
	 * The kind of events of interest. If omitted it defaults
	 * to WatchKind.Create | WatchKind.Change | WatchKind.Delete
	 * which is 7.
	 */
	Kind float64 `json:"kind,omitempty"`
}

<<<<<<< HEAD
// PublishDiagnosticsParams is:
/**
=======
/*PublishDiagnosticsParams defined:
>>>>>>> v0.0.4
 * The publish diagnostic notification's parameters.
 */
type PublishDiagnosticsParams struct {

<<<<<<< HEAD
	/** URI defined:
=======
	/*URI defined:
>>>>>>> v0.0.4
	 * The URI for which diagnostic information is reported.
	 */
	URI string `json:"uri"`

<<<<<<< HEAD
	/** Version defined:
=======
	/*Version defined:
>>>>>>> v0.0.4
	 * Optional the version number of the document the diagnostics are published for.
	 */
	Version float64 `json:"version,omitempty"`

<<<<<<< HEAD
	/** Diagnostics defined:
=======
	/*Diagnostics defined:
>>>>>>> v0.0.4
	 * An array of diagnostic information items.
	 */
	Diagnostics []Diagnostic `json:"diagnostics"`
}

<<<<<<< HEAD
// CompletionRegistrationOptions is:
/**
=======
/*CompletionRegistrationOptions defined:
>>>>>>> v0.0.4
 * Completion registration options.
 */
type CompletionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CompletionOptions
}

<<<<<<< HEAD
// CompletionContext is:
/**
=======
/*CompletionContext defined:
>>>>>>> v0.0.4
 * Contains additional information about the context in which a completion request is triggered.
 */
type CompletionContext struct {

<<<<<<< HEAD
	/** TriggerKind defined:
=======
	/*TriggerKind defined:
>>>>>>> v0.0.4
	 * How the completion was triggered.
	 */
	TriggerKind CompletionTriggerKind `json:"triggerKind"`

<<<<<<< HEAD
	/** TriggerCharacter defined:
=======
	/*TriggerCharacter defined:
>>>>>>> v0.0.4
	 * The trigger character (a single character) that has trigger code complete.
	 * Is undefined if `triggerKind !== CompletionTriggerKind.TriggerCharacter`
	 */
	TriggerCharacter string `json:"triggerCharacter,omitempty"`
}

<<<<<<< HEAD
// CompletionParams is:
/**
 * Completion parameters
 */
type CompletionParams struct {
	TextDocumentPositionParams

	/** Context defined:
=======
/*CompletionParams defined:
 * Completion parameters
 */
type CompletionParams struct {

	/*Context defined:
>>>>>>> v0.0.4
	 * The completion context. This is only available it the client specifies
	 * to send this using `ClientCapabilities.textDocument.completion.contextSupport === true`
	 */
	Context *CompletionContext `json:"context,omitempty"`
<<<<<<< HEAD
}

// SignatureHelpRegistrationOptions is:
/**
=======
	TextDocumentPositionParams
}

/*SignatureHelpRegistrationOptions defined:
>>>>>>> v0.0.4
 * Signature help registration options.
 */
type SignatureHelpRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SignatureHelpOptions
}

<<<<<<< HEAD
// ReferenceParams is:
/**
 * Parameters for a [ReferencesRequest](#ReferencesRequest).
 */
type ReferenceParams struct {
	TextDocumentPositionParams

	// Context is
	Context ReferenceContext `json:"context"`
}

// CodeActionParams is:
/**
=======
/*ReferenceParams defined:
 * Parameters for a [ReferencesRequest](#ReferencesRequest).
 */
type ReferenceParams struct {

	// Context is
	Context ReferenceContext `json:"context"`
	TextDocumentPositionParams
}

/*CodeActionParams defined:
>>>>>>> v0.0.4
 * Params for the CodeActionRequest
 */
type CodeActionParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document in which the command was invoked.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range for which the command was invoked.
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** Context defined:
=======
	/*Context defined:
>>>>>>> v0.0.4
	 * Context carrying additional information.
	 */
	Context CodeActionContext `json:"context"`
}

<<<<<<< HEAD
// CodeActionRegistrationOptions is:
=======
// CodeActionRegistrationOptions is
>>>>>>> v0.0.4
type CodeActionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeActionOptions
}

<<<<<<< HEAD
// CodeLensParams is:
/**
=======
/*CodeLensParams defined:
>>>>>>> v0.0.4
 * Params for the Code Lens request.
 */
type CodeLensParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document to request code lens for.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

<<<<<<< HEAD
// CodeLensRegistrationOptions is:
/**
=======
/*CodeLensRegistrationOptions defined:
>>>>>>> v0.0.4
 * Code Lens registration options.
 */
type CodeLensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeLensOptions
}

<<<<<<< HEAD
// DocumentFormattingParams is:
type DocumentFormattingParams struct {

	/** TextDocument defined:
=======
// DocumentFormattingParams is
type DocumentFormattingParams struct {

	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document to format.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Options defined:
=======
	/*Options defined:
>>>>>>> v0.0.4
	 * The format options
	 */
	Options FormattingOptions `json:"options"`
}

<<<<<<< HEAD
// DocumentRangeFormattingParams is:
type DocumentRangeFormattingParams struct {

	/** TextDocument defined:
=======
// DocumentRangeFormattingParams is
type DocumentRangeFormattingParams struct {

	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document to format.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range to format
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** Options defined:
=======
	/*Options defined:
>>>>>>> v0.0.4
	 * The format options
	 */
	Options FormattingOptions `json:"options"`
}

<<<<<<< HEAD
// DocumentOnTypeFormattingParams is:
type DocumentOnTypeFormattingParams struct {

	/** TextDocument defined:
=======
// DocumentOnTypeFormattingParams is
type DocumentOnTypeFormattingParams struct {

	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document to format.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Position defined:
=======
	/*Position defined:
>>>>>>> v0.0.4
	 * The position at which this request was send.
	 */
	Position Position `json:"position"`

<<<<<<< HEAD
	/** Ch defined:
=======
	/*Ch defined:
>>>>>>> v0.0.4
	 * The character that has been typed.
	 */
	Ch string `json:"ch"`

<<<<<<< HEAD
	/** Options defined:
=======
	/*Options defined:
>>>>>>> v0.0.4
	 * The format options.
	 */
	Options FormattingOptions `json:"options"`
}

<<<<<<< HEAD
// DocumentOnTypeFormattingRegistrationOptions is:
/**
=======
/*DocumentOnTypeFormattingRegistrationOptions defined:
>>>>>>> v0.0.4
 * Format document on type options
 */
type DocumentOnTypeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentOnTypeFormattingOptions
}

<<<<<<< HEAD
// RenameParams is:
type RenameParams struct {

	/** TextDocument defined:
=======
// RenameParams is
type RenameParams struct {

	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document to rename.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Position defined:
=======
	/*Position defined:
>>>>>>> v0.0.4
	 * The position at which this request was sent.
	 */
	Position Position `json:"position"`

<<<<<<< HEAD
	/** NewName defined:
=======
	/*NewName defined:
>>>>>>> v0.0.4
	 * The new name of the symbol. If the given name is not valid the
	 * request must return a [ResponseError](#ResponseError) with an
	 * appropriate message set.
	 */
	NewName string `json:"newName"`
}

<<<<<<< HEAD
// RenameRegistrationOptions is:
/**
=======
/*RenameRegistrationOptions defined:
>>>>>>> v0.0.4
 * Rename registration options.
 */
type RenameRegistrationOptions struct {
	TextDocumentRegistrationOptions
	RenameOptions
}

<<<<<<< HEAD
// DocumentLinkParams is:
type DocumentLinkParams struct {

	/** TextDocument defined:
=======
// DocumentLinkParams is
type DocumentLinkParams struct {

	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The document to provide document links for.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

<<<<<<< HEAD
// DocumentLinkRegistrationOptions is:
/**
=======
/*DocumentLinkRegistrationOptions defined:
>>>>>>> v0.0.4
 * Document link registration options
 */
type DocumentLinkRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentLinkOptions
}

<<<<<<< HEAD
// ExecuteCommandParams is:
type ExecuteCommandParams struct {

	/** Command defined:
=======
// ExecuteCommandParams is
type ExecuteCommandParams struct {

	/*Command defined:
>>>>>>> v0.0.4
	 * The identifier of the actual command handler.
	 */
	Command string `json:"command"`

<<<<<<< HEAD
	/** Arguments defined:
=======
	/*Arguments defined:
>>>>>>> v0.0.4
	 * Arguments that the command should be invoked with.
	 */
	Arguments []interface{} `json:"arguments,omitempty"`
}

<<<<<<< HEAD
// ExecuteCommandRegistrationOptions is:
/**
=======
/*ExecuteCommandRegistrationOptions defined:
>>>>>>> v0.0.4
 * Execute command registration options.
 */
type ExecuteCommandRegistrationOptions struct {
	ExecuteCommandOptions
}

<<<<<<< HEAD
// ApplyWorkspaceEditParams is:
/**
=======
/*ApplyWorkspaceEditParams defined:
>>>>>>> v0.0.4
 * The parameters passed via a apply workspace edit request.
 */
type ApplyWorkspaceEditParams struct {

<<<<<<< HEAD
	/** Label defined:
=======
	/*Label defined:
>>>>>>> v0.0.4
	 * An optional label of the workspace edit. This label is
	 * presented in the user interface for example on an undo
	 * stack to undo the workspace edit.
	 */
	Label string `json:"label,omitempty"`

<<<<<<< HEAD
	/** Edit defined:
=======
	/*Edit defined:
>>>>>>> v0.0.4
	 * The edits to apply.
	 */
	Edit WorkspaceEdit `json:"edit"`
}

<<<<<<< HEAD
// ApplyWorkspaceEditResponse is:
/**
=======
/*ApplyWorkspaceEditResponse defined:
>>>>>>> v0.0.4
 * A response returned from the apply workspace edit request.
 */
type ApplyWorkspaceEditResponse struct {

<<<<<<< HEAD
	/** Applied defined:
=======
	/*Applied defined:
>>>>>>> v0.0.4
	 * Indicates whether the edit was applied or not.
	 */
	Applied bool `json:"applied"`

<<<<<<< HEAD
	/** FailureReason defined:
=======
	/*FailureReason defined:
>>>>>>> v0.0.4
	 * An optional textual description for why the edit was not applied.
	 * This may be used by the server for diagnostic logging or to provide
	 * a suitable error for a request that triggered the edit.
	 */
	FailureReason string `json:"failureReason,omitempty"`

<<<<<<< HEAD
	/** FailedChange defined:
=======
	/*FailedChange defined:
>>>>>>> v0.0.4
	 * Depending on the client's failure handling strategy `failedChange` might
	 * contain the index of the change that failed. This property is only available
	 * if the client signals a `failureHandlingStrategy` in its client capabilities.
	 */
	FailedChange float64 `json:"failedChange,omitempty"`
}

<<<<<<< HEAD
// Position is:
/**
=======
/*Position defined:
>>>>>>> v0.0.4
 * Position in a text document expressed as zero-based line and character offset.
 * The offsets are based on a UTF-16 string representation. So a string of the form
 * `a𐐀b` the character offset of the character `a` is 0, the character offset of `𐐀`
 * is 1 and the character offset of b is 3 since `𐐀` is represented using two code
 * units in UTF-16.
 *
 * Positions are line end character agnostic. So you can not specify a position that
 * denotes `\r|\n` or `\n|` where `|` represents the character offset.
 */
type Position struct {

<<<<<<< HEAD
	/** Line defined:
=======
	/*Line defined:
>>>>>>> v0.0.4
	 * Line position in a document (zero-based).
	 * If a line number is greater than the number of lines in a document, it defaults back to the number of lines in the document.
	 * If a line number is negative, it defaults to 0.
	 */
	Line float64 `json:"line"`

<<<<<<< HEAD
	/** Character defined:
=======
	/*Character defined:
>>>>>>> v0.0.4
	 * Character offset on a line in a document (zero-based). Assuming that the line is
	 * represented as a string, the `character` value represents the gap between the
	 * `character` and `character + 1`.
	 *
	 * If the character value is greater than the line length it defaults back to the
	 * line length.
	 * If a line number is negative, it defaults to 0.
	 */
	Character float64 `json:"character"`
}

<<<<<<< HEAD
// Range is:
/**
=======
/*Range defined:
>>>>>>> v0.0.4
 * A range in a text document expressed as (zero-based) start and end positions.
 *
 * If you want to specify a range that contains a line including the line ending
 * character(s) then use an end position denoting the start of the next line.
 * For example:
 * ```ts
 * {
 *     start: { line: 5, character: 23 }
 *     end : { line 6, character : 0 }
 * }
 * ```
 */
type Range struct {

<<<<<<< HEAD
	/** Start defined:
=======
	/*Start defined:
>>>>>>> v0.0.4
	 * The range's start position
	 */
	Start Position `json:"start"`

<<<<<<< HEAD
	/** End defined:
=======
	/*End defined:
>>>>>>> v0.0.4
	 * The range's end position.
	 */
	End Position `json:"end"`
}

<<<<<<< HEAD
// Location is:
/**
=======
/*Location defined:
>>>>>>> v0.0.4
 * Represents a location inside a resource, such as a line
 * inside a text file.
 */
type Location struct {

	// URI is
	URI string `json:"uri"`

	// Range is
	Range Range `json:"range"`
}

<<<<<<< HEAD
// LocationLink is:
/**
=======
/*LocationLink defined:
>>>>>>> v0.0.4
 * Represents the connection of two locations. Provides additional metadata over normal [locations](#Location),
 * including an origin range.
 */
type LocationLink struct {

<<<<<<< HEAD
	/** OriginSelectionRange defined:
=======
	/*OriginSelectionRange defined:
>>>>>>> v0.0.4
	 * Span of the origin of this link.
	 *
	 * Used as the underlined span for mouse definition hover. Defaults to the word range at
	 * the definition position.
	 */
	OriginSelectionRange *Range `json:"originSelectionRange,omitempty"`

<<<<<<< HEAD
	/** TargetURI defined:
=======
	/*TargetURI defined:
>>>>>>> v0.0.4
	 * The target resource identifier of this link.
	 */
	TargetURI string `json:"targetUri"`

<<<<<<< HEAD
	/** TargetRange defined:
=======
	/*TargetRange defined:
>>>>>>> v0.0.4
	 * The full target range of this link. If the target for example is a symbol then target range is the
	 * range enclosing this symbol not including leading/trailing whitespace but everything else
	 * like comments. This information is typically used to highlight the range in the editor.
	 */
	TargetRange Range `json:"targetRange"`

<<<<<<< HEAD
	/** TargetSelectionRange defined:
=======
	/*TargetSelectionRange defined:
>>>>>>> v0.0.4
	 * The range that should be selected and revealed when this link is being followed, e.g the name of a function.
	 * Must be contained by the the `targetRange`. See also `DocumentSymbol#range`
	 */
	TargetSelectionRange Range `json:"targetSelectionRange"`
}

<<<<<<< HEAD
// Color is:
/**
=======
/*Color defined:
>>>>>>> v0.0.4
 * Represents a color in RGBA space.
 */
type Color struct {

<<<<<<< HEAD
	/** Red defined:
=======
	/*Red defined:
>>>>>>> v0.0.4
	 * The red component of this color in the range [0-1].
	 */
	Red float64 `json:"red"`

<<<<<<< HEAD
	/** Green defined:
=======
	/*Green defined:
>>>>>>> v0.0.4
	 * The green component of this color in the range [0-1].
	 */
	Green float64 `json:"green"`

<<<<<<< HEAD
	/** Blue defined:
=======
	/*Blue defined:
>>>>>>> v0.0.4
	 * The blue component of this color in the range [0-1].
	 */
	Blue float64 `json:"blue"`

<<<<<<< HEAD
	/** Alpha defined:
=======
	/*Alpha defined:
>>>>>>> v0.0.4
	 * The alpha component of this color in the range [0-1].
	 */
	Alpha float64 `json:"alpha"`
}

<<<<<<< HEAD
// ColorInformation is:
/**
=======
/*ColorInformation defined:
>>>>>>> v0.0.4
 * Represents a color range from a document.
 */
type ColorInformation struct {

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range in the document where this color appers.
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** Color defined:
=======
	/*Color defined:
>>>>>>> v0.0.4
	 * The actual color value for this color range.
	 */
	Color Color `json:"color"`
}

<<<<<<< HEAD
// ColorPresentation is:
type ColorPresentation struct {

	/** Label defined:
=======
// ColorPresentation is
type ColorPresentation struct {

	/*Label defined:
>>>>>>> v0.0.4
	 * The label of this color presentation. It will be shown on the color
	 * picker header. By default this is also the text that is inserted when selecting
	 * this color presentation.
	 */
	Label string `json:"label"`

<<<<<<< HEAD
	/** TextEdit defined:
=======
	/*TextEdit defined:
>>>>>>> v0.0.4
	 * An [edit](#TextEdit) which is applied to a document when selecting
	 * this presentation for the color.  When `falsy` the [label](#ColorPresentation.label)
	 * is used.
	 */
	TextEdit *TextEdit `json:"textEdit,omitempty"`

<<<<<<< HEAD
	/** AdditionalTextEdits defined:
=======
	/*AdditionalTextEdits defined:
>>>>>>> v0.0.4
	 * An optional array of additional [text edits](#TextEdit) that are applied when
	 * selecting this color presentation. Edits must not overlap with the main [edit](#ColorPresentation.textEdit) nor with themselves.
	 */
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
}

<<<<<<< HEAD
// DiagnosticRelatedInformation is:
/**
=======
/*DiagnosticRelatedInformation defined:
>>>>>>> v0.0.4
 * Represents a related message and source code location for a diagnostic. This should be
 * used to point to code locations that cause or related to a diagnostics, e.g when duplicating
 * a symbol in a scope.
 */
type DiagnosticRelatedInformation struct {

<<<<<<< HEAD
	/** Location defined:
=======
	/*Location defined:
>>>>>>> v0.0.4
	 * The location of this related diagnostic information.
	 */
	Location Location `json:"location"`

<<<<<<< HEAD
	/** Message defined:
=======
	/*Message defined:
>>>>>>> v0.0.4
	 * The message of this related diagnostic information.
	 */
	Message string `json:"message"`
}

<<<<<<< HEAD
// Diagnostic is:
/**
=======
/*Diagnostic defined:
>>>>>>> v0.0.4
 * Represents a diagnostic, such as a compiler error or warning. Diagnostic objects
 * are only valid in the scope of a resource.
 */
type Diagnostic struct {

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range at which the message applies
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** Severity defined:
=======
	/*Severity defined:
>>>>>>> v0.0.4
	 * The diagnostic's severity. Can be omitted. If omitted it is up to the
	 * client to interpret diagnostics as error, warning, info or hint.
	 */
	Severity DiagnosticSeverity `json:"severity,omitempty"`

<<<<<<< HEAD
	/** Code defined:
=======
	/*Code defined:
>>>>>>> v0.0.4
	 * The diagnostic's code, which usually appear in the user interface.
	 */
	Code interface{} `json:"code,omitempty"` // number | string

<<<<<<< HEAD
	/** Source defined:
=======
	/*Source defined:
>>>>>>> v0.0.4
	 * A human-readable string describing the source of this
	 * diagnostic, e.g. 'typescript' or 'super lint'. It usually
	 * appears in the user interface.
	 */
	Source string `json:"source,omitempty"`

<<<<<<< HEAD
	/** Message defined:
=======
	/*Message defined:
>>>>>>> v0.0.4
	 * The diagnostic's message. It usually appears in the user interface
	 */
	Message string `json:"message"`

<<<<<<< HEAD
	/** Tags defined:
=======
	/*Tags defined:
>>>>>>> v0.0.4
	 * Additional metadata about the diagnostic.
	 */
	Tags []DiagnosticTag `json:"tags,omitempty"`

<<<<<<< HEAD
	/** RelatedInformation defined:
=======
	/*RelatedInformation defined:
>>>>>>> v0.0.4
	 * An array of related diagnostic information, e.g. when symbol-names within
	 * a scope collide all definitions can be marked via this property.
	 */
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`
}

<<<<<<< HEAD
// Command is:
/**
=======
/*Command defined:
>>>>>>> v0.0.4
 * Represents a reference to a command. Provides a title which
 * will be used to represent a command in the UI and, optionally,
 * an array of arguments which will be passed to the command handler
 * function when invoked.
 */
type Command struct {

<<<<<<< HEAD
	/** Title defined:
=======
	/*Title defined:
>>>>>>> v0.0.4
	 * Title of the command, like `save`.
	 */
	Title string `json:"title"`

<<<<<<< HEAD
	/** Command defined:
=======
	/*Command defined:
>>>>>>> v0.0.4
	 * The identifier of the actual command handler.
	 */
	Command string `json:"command"`

<<<<<<< HEAD
	/** Arguments defined:
=======
	/*Arguments defined:
>>>>>>> v0.0.4
	 * Arguments that the command handler should be
	 * invoked with.
	 */
	Arguments []interface{} `json:"arguments,omitempty"`
}

<<<<<<< HEAD
// TextEdit is:
/**
=======
/*TextEdit defined:
>>>>>>> v0.0.4
 * A text edit applicable to a text document.
 */
type TextEdit struct {

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range of the text document to be manipulated. To insert
	 * text into a document create a range where start === end.
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** NewText defined:
=======
	/*NewText defined:
>>>>>>> v0.0.4
	 * The string to be inserted. For delete operations use an
	 * empty string.
	 */
	NewText string `json:"newText"`
}

<<<<<<< HEAD
// TextDocumentEdit is:
/**
=======
/*TextDocumentEdit defined:
>>>>>>> v0.0.4
 * Describes textual changes on a text document.
 */
type TextDocumentEdit struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The text document to change.
	 */
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

<<<<<<< HEAD
	/** Edits defined:
=======
	/*Edits defined:
>>>>>>> v0.0.4
	 * The edits to be applied.
	 */
	Edits []TextEdit `json:"edits"`
}

<<<<<<< HEAD
// ResourceOperation is:
=======
// ResourceOperation is
>>>>>>> v0.0.4
type ResourceOperation struct {

	// Kind is
	Kind string `json:"kind"`
}

<<<<<<< HEAD
// CreateFileOptions is:
/**
=======
/*CreateFileOptions defined:
>>>>>>> v0.0.4
 * Options to create a file.
 */
type CreateFileOptions struct {

<<<<<<< HEAD
	/** Overwrite defined:
=======
	/*Overwrite defined:
>>>>>>> v0.0.4
	 * Overwrite existing file. Overwrite wins over `ignoreIfExists`
	 */
	Overwrite bool `json:"overwrite,omitempty"`

<<<<<<< HEAD
	/** IgnoreIfExists defined:
=======
	/*IgnoreIfExists defined:
>>>>>>> v0.0.4
	 * Ignore if exists.
	 */
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

<<<<<<< HEAD
// CreateFile is:
/**
 * Create file operation.
 */
type CreateFile struct {
	ResourceOperation

	/** Kind defined:
=======
/*CreateFile defined:
 * Create file operation.
 */
type CreateFile struct {

	/*Kind defined:
>>>>>>> v0.0.4
	 * A create
	 */
	Kind string `json:"kind"` // 'create'

<<<<<<< HEAD
	/** URI defined:
=======
	/*URI defined:
>>>>>>> v0.0.4
	 * The resource to create.
	 */
	URI string `json:"uri"`

<<<<<<< HEAD
	/** Options defined:
=======
	/*Options defined:
>>>>>>> v0.0.4
	 * Additional options
	 */
	Options *CreateFileOptions `json:"options,omitempty"`
}

<<<<<<< HEAD
// RenameFileOptions is:
/**
=======
/*RenameFileOptions defined:
>>>>>>> v0.0.4
 * Rename file options
 */
type RenameFileOptions struct {

<<<<<<< HEAD
	/** Overwrite defined:
=======
	/*Overwrite defined:
>>>>>>> v0.0.4
	 * Overwrite target if existing. Overwrite wins over `ignoreIfExists`
	 */
	Overwrite bool `json:"overwrite,omitempty"`

<<<<<<< HEAD
	/** IgnoreIfExists defined:
=======
	/*IgnoreIfExists defined:
>>>>>>> v0.0.4
	 * Ignores if target exists.
	 */
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

<<<<<<< HEAD
// RenameFile is:
/**
 * Rename file operation
 */
type RenameFile struct {
	ResourceOperation

	/** Kind defined:
=======
/*RenameFile defined:
 * Rename file operation
 */
type RenameFile struct {

	/*Kind defined:
>>>>>>> v0.0.4
	 * A rename
	 */
	Kind string `json:"kind"` // 'rename'

<<<<<<< HEAD
	/** OldURI defined:
=======
	/*OldURI defined:
>>>>>>> v0.0.4
	 * The old (existing) location.
	 */
	OldURI string `json:"oldUri"`

<<<<<<< HEAD
	/** NewURI defined:
=======
	/*NewURI defined:
>>>>>>> v0.0.4
	 * The new location.
	 */
	NewURI string `json:"newUri"`

<<<<<<< HEAD
	/** Options defined:
=======
	/*Options defined:
>>>>>>> v0.0.4
	 * Rename options.
	 */
	Options *RenameFileOptions `json:"options,omitempty"`
}

<<<<<<< HEAD
// DeleteFileOptions is:
/**
=======
/*DeleteFileOptions defined:
>>>>>>> v0.0.4
 * Delete file options
 */
type DeleteFileOptions struct {

<<<<<<< HEAD
	/** Recursive defined:
=======
	/*Recursive defined:
>>>>>>> v0.0.4
	 * Delete the content recursively if a folder is denoted.
	 */
	Recursive bool `json:"recursive,omitempty"`

<<<<<<< HEAD
	/** IgnoreIfNotExists defined:
=======
	/*IgnoreIfNotExists defined:
>>>>>>> v0.0.4
	 * Ignore the operation if the file doesn't exist.
	 */
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
}

<<<<<<< HEAD
// DeleteFile is:
/**
 * Delete file operation
 */
type DeleteFile struct {
	ResourceOperation

	/** Kind defined:
=======
/*DeleteFile defined:
 * Delete file operation
 */
type DeleteFile struct {

	/*Kind defined:
>>>>>>> v0.0.4
	 * A delete
	 */
	Kind string `json:"kind"` // 'delete'

<<<<<<< HEAD
	/** URI defined:
=======
	/*URI defined:
>>>>>>> v0.0.4
	 * The file to delete.
	 */
	URI string `json:"uri"`

<<<<<<< HEAD
	/** Options defined:
=======
	/*Options defined:
>>>>>>> v0.0.4
	 * Delete options.
	 */
	Options *DeleteFileOptions `json:"options,omitempty"`
}

<<<<<<< HEAD
// WorkspaceEdit is:
/**
=======
/*WorkspaceEdit defined:
>>>>>>> v0.0.4
 * A workspace edit represents changes to many resources managed in the workspace. The edit
 * should either provide `changes` or `documentChanges`. If documentChanges are present
 * they are preferred over `changes` if the client can handle versioned document edits.
 */
type WorkspaceEdit struct {

<<<<<<< HEAD
	/** Changes defined:
=======
	/*Changes defined:
>>>>>>> v0.0.4
	 * Holds changes to existing resources.
	 */
	Changes *map[string][]TextEdit `json:"changes,omitempty"` // [uri: string]: TextEdit[];

<<<<<<< HEAD
	/** DocumentChanges defined:
=======
	/*DocumentChanges defined:
>>>>>>> v0.0.4
	 * Depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes
	 * are either an array of `TextDocumentEdit`s to express changes to n different text documents
	 * where each text document edit addresses a specific version of a text document. Or it can contain
	 * above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations.
	 *
	 * Whether a client supports versioned document edits is expressed via
	 * `workspace.workspaceEdit.documentChanges` client capability.
	 *
	 * If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then
	 * only plain `TextEdit`s using the `changes` property are supported.
	 */
	DocumentChanges []TextDocumentEdit `json:"documentChanges,omitempty"` // (TextDocumentEdit | CreateFile | RenameFile | DeleteFile)
}

<<<<<<< HEAD
// TextEditChange is:
/**
=======
/*TextEditChange defined:
>>>>>>> v0.0.4
 * A change to capture text edits for existing resources.
 */
type TextEditChange struct {
}

<<<<<<< HEAD
// TextDocumentIdentifier is:
/**
=======
/*TextDocumentIdentifier defined:
>>>>>>> v0.0.4
 * A literal to identify a text document in the client.
 */
type TextDocumentIdentifier struct {

<<<<<<< HEAD
	/** URI defined:
=======
	/*URI defined:
>>>>>>> v0.0.4
	 * The text document's uri.
	 */
	URI string `json:"uri"`
}

<<<<<<< HEAD
// VersionedTextDocumentIdentifier is:
/**
 * An identifier to denote a specific version of a text document.
 */
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	/** Version defined:
=======
/*VersionedTextDocumentIdentifier defined:
 * An identifier to denote a specific version of a text document.
 */
type VersionedTextDocumentIdentifier struct {

	/*Version defined:
>>>>>>> v0.0.4
	 * The version number of this document. If a versioned text document identifier
	 * is sent from the server to the client and the file is not open in the editor
	 * (the server has not received an open notification before) the server can send
	 * `null` to indicate that the version is unknown and the content on disk is the
	 * truth (as speced with document content ownership).
	 */
	Version float64 `json:"version"`
<<<<<<< HEAD
}

// TextDocumentItem is:
/**
=======
	TextDocumentIdentifier
}

/*TextDocumentItem defined:
>>>>>>> v0.0.4
 * An item to transfer a text document from the client to the
 * server.
 */
type TextDocumentItem struct {

<<<<<<< HEAD
	/** URI defined:
=======
	/*URI defined:
>>>>>>> v0.0.4
	 * The text document's uri.
	 */
	URI string `json:"uri"`

<<<<<<< HEAD
	/** LanguageID defined:
=======
	/*LanguageID defined:
>>>>>>> v0.0.4
	 * The text document's language identifier
	 */
	LanguageID string `json:"languageId"`

<<<<<<< HEAD
	/** Version defined:
=======
	/*Version defined:
>>>>>>> v0.0.4
	 * The version number of this document (it will increase after each
	 * change, including undo/redo).
	 */
	Version float64 `json:"version"`

<<<<<<< HEAD
	/** Text defined:
=======
	/*Text defined:
>>>>>>> v0.0.4
	 * The content of the opened text document.
	 */
	Text string `json:"text"`
}

<<<<<<< HEAD
// MarkupContent is:
/**
=======
/*MarkupContent defined:
>>>>>>> v0.0.4
 * A `MarkupContent` literal represents a string value which content is interpreted base on its
 * kind flag. Currently the protocol supports `plaintext` and `markdown` as markup kinds.
 *
 * If the kind is `markdown` then the value can contain fenced code blocks like in GitHub issues.
 * See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting
 *
 * Here is an example how such a string can be constructed using JavaScript / TypeScript:
 * ```ts
 * let markdown: MarkdownContent = {
 *  kind: MarkupKind.Markdown,
 *	value: [
 *		'# Header',
 *		'Some text',
 *		'```typescript',
 *		'someCode();',
 *		'```'
 *	].join('\n')
 * };
 * ```
 *
 * *Please Note* that clients might sanitize the return markdown. A client could decide to
 * remove HTML from the markdown to avoid script execution.
 */
type MarkupContent struct {

<<<<<<< HEAD
	/** Kind defined:
=======
	/*Kind defined:
>>>>>>> v0.0.4
	 * The type of the Markup
	 */
	Kind MarkupKind `json:"kind"`

<<<<<<< HEAD
	/** Value defined:
=======
	/*Value defined:
>>>>>>> v0.0.4
	 * The content itself
	 */
	Value string `json:"value"`
}

<<<<<<< HEAD
// CompletionItem is:
/**
=======
/*CompletionItem defined:
>>>>>>> v0.0.4
 * A completion item represents a text snippet that is
 * proposed to complete text that is being typed.
 */
type CompletionItem struct {

<<<<<<< HEAD
	/** Label defined:
=======
	/*Label defined:
>>>>>>> v0.0.4
	 * The label of this completion item. By default
	 * also the text that is inserted when selecting
	 * this completion.
	 */
	Label string `json:"label"`

<<<<<<< HEAD
	/** Kind defined:
=======
	/*Kind defined:
>>>>>>> v0.0.4
	 * The kind of this completion item. Based of the kind
	 * an icon is chosen by the editor.
	 */
	Kind CompletionItemKind `json:"kind,omitempty"`

<<<<<<< HEAD
	/** Detail defined:
=======
	/*Detail defined:
>>>>>>> v0.0.4
	 * A human-readable string with additional information
	 * about this item, like type or symbol information.
	 */
	Detail string `json:"detail,omitempty"`

<<<<<<< HEAD
	/** Documentation defined:
=======
	/*Documentation defined:
>>>>>>> v0.0.4
	 * A human-readable string that represents a doc-comment.
	 */
	Documentation string `json:"documentation,omitempty"` // string | MarkupContent

<<<<<<< HEAD
	/** Deprecated defined:
=======
	/*Deprecated defined:
>>>>>>> v0.0.4
	 * Indicates if this item is deprecated.
	 */
	Deprecated bool `json:"deprecated,omitempty"`

<<<<<<< HEAD
	/** Preselect defined:
=======
	/*Preselect defined:
>>>>>>> v0.0.4
	 * Select this item when showing.
	 *
	 * *Note* that only one completion item can be selected and that the
	 * tool / client decides which item that is. The rule is that the *first*
	 * item of those that match best is selected.
	 */
	Preselect bool `json:"preselect,omitempty"`

<<<<<<< HEAD
	/** SortText defined:
=======
	/*SortText defined:
>>>>>>> v0.0.4
	 * A string that should be used when comparing this item
	 * with other items. When `falsy` the [label](#CompletionItem.label)
	 * is used.
	 */
	SortText string `json:"sortText,omitempty"`

<<<<<<< HEAD
	/** FilterText defined:
=======
	/*FilterText defined:
>>>>>>> v0.0.4
	 * A string that should be used when filtering a set of
	 * completion items. When `falsy` the [label](#CompletionItem.label)
	 * is used.
	 */
	FilterText string `json:"filterText,omitempty"`

<<<<<<< HEAD
	/** InsertText defined:
=======
	/*InsertText defined:
>>>>>>> v0.0.4
	 * A string that should be inserted into a document when selecting
	 * this completion. When `falsy` the [label](#CompletionItem.label)
	 * is used.
	 *
	 * The `insertText` is subject to interpretation by the client side.
	 * Some tools might not take the string literally. For example
	 * VS Code when code complete is requested in this example `con<cursor position>`
	 * and a completion item with an `insertText` of `console` is provided it
	 * will only insert `sole`. Therefore it is recommended to use `textEdit` instead
	 * since it avoids additional client side interpretation.
	 *
	 * @deprecated Use textEdit instead.
	 */
	InsertText string `json:"insertText,omitempty"`

<<<<<<< HEAD
	/** InsertTextFormat defined:
=======
	/*InsertTextFormat defined:
>>>>>>> v0.0.4
	 * The format of the insert text. The format applies to both the `insertText` property
	 * and the `newText` property of a provided `textEdit`.
	 */
	InsertTextFormat InsertTextFormat `json:"insertTextFormat,omitempty"`

<<<<<<< HEAD
	/** TextEdit defined:
=======
	/*TextEdit defined:
>>>>>>> v0.0.4
	 * An [edit](#TextEdit) which is applied to a document when selecting
	 * this completion. When an edit is provided the value of
	 * [insertText](#CompletionItem.insertText) is ignored.
	 *
	 * *Note:* The text edit's range must be a [single line] and it must contain the position
	 * at which completion has been requested.
	 */
	TextEdit *TextEdit `json:"textEdit,omitempty"`

<<<<<<< HEAD
	/** AdditionalTextEdits defined:
=======
	/*AdditionalTextEdits defined:
>>>>>>> v0.0.4
	 * An optional array of additional [text edits](#TextEdit) that are applied when
	 * selecting this completion. Edits must not overlap (including the same insert position)
	 * with the main [edit](#CompletionItem.textEdit) nor with themselves.
	 *
	 * Additional text edits should be used to change text unrelated to the current cursor position
	 * (for example adding an import statement at the top of the file if the completion item will
	 * insert an unqualified type).
	 */
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`

<<<<<<< HEAD
	/** CommitCharacters defined:
=======
	/*CommitCharacters defined:
>>>>>>> v0.0.4
	 * An optional set of characters that when pressed while this completion is active will accept it first and
	 * then type that character. *Note* that all commit characters should have `length=1` and that superfluous
	 * characters will be ignored.
	 */
	CommitCharacters []string `json:"commitCharacters,omitempty"`

<<<<<<< HEAD
	/** Command defined:
=======
	/*Command defined:
>>>>>>> v0.0.4
	 * An optional [command](#Command) that is executed *after* inserting this completion. *Note* that
	 * additional modifications to the current document should be described with the
	 * [additionalTextEdits](#CompletionItem.additionalTextEdits)-property.
	 */
	Command *Command `json:"command,omitempty"`

<<<<<<< HEAD
	/** Data defined:
=======
	/*Data defined:
>>>>>>> v0.0.4
	 * An data entry field that is preserved on a completion item between
	 * a [CompletionRequest](#CompletionRequest) and a [CompletionResolveRequest]
	 * (#CompletionResolveRequest)
	 */
	Data interface{} `json:"data,omitempty"`
}

<<<<<<< HEAD
// CompletionList is:
/**
=======
/*CompletionList defined:
>>>>>>> v0.0.4
 * Represents a collection of [completion items](#CompletionItem) to be presented
 * in the editor.
 */
type CompletionList struct {

<<<<<<< HEAD
	/** IsIncomplete defined:
=======
	/*IsIncomplete defined:
>>>>>>> v0.0.4
	 * This list it not complete. Further typing results in recomputing this list.
	 */
	IsIncomplete bool `json:"isIncomplete"`

<<<<<<< HEAD
	/** Items defined:
=======
	/*Items defined:
>>>>>>> v0.0.4
	 * The completion items.
	 */
	Items []CompletionItem `json:"items"`
}

<<<<<<< HEAD
// Hover is:
/**
=======
/*Hover defined:
>>>>>>> v0.0.4
 * The result of a hover request.
 */
type Hover struct {

<<<<<<< HEAD
	/** Contents defined:
=======
	/*Contents defined:
>>>>>>> v0.0.4
	 * The hover's content
	 */
	Contents MarkupContent `json:"contents"` // MarkupContent | MarkedString | MarkedString[]

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * An optional range
	 */
	Range *Range `json:"range,omitempty"`
}

<<<<<<< HEAD
// ParameterInformation is:
/**
=======
/*ParameterInformation defined:
>>>>>>> v0.0.4
 * Represents a parameter of a callable-signature. A parameter can
 * have a label and a doc-comment.
 */
type ParameterInformation struct {

<<<<<<< HEAD
	/** Label defined:
=======
	/*Label defined:
>>>>>>> v0.0.4
	 * The label of this parameter information.
	 *
	 * Either a string or an inclusive start and exclusive end offsets within its containing
	 * signature label. (see SignatureInformation.label). The offsets are based on a UTF-16
	 * string representation as `Position` and `Range` does.
	 *
	 * *Note*: a label of type string should be a substring of its containing signature label.
	 * Its intended use case is to highlight the parameter label part in the `SignatureInformation.label`.
	 */
	Label string `json:"label"` // string | [number, number]

<<<<<<< HEAD
	/** Documentation defined:
=======
	/*Documentation defined:
>>>>>>> v0.0.4
	 * The human-readable doc-comment of this signature. Will be shown
	 * in the UI but can be omitted.
	 */
	Documentation string `json:"documentation,omitempty"` // string | MarkupContent
}

<<<<<<< HEAD
// SignatureInformation is:
/**
=======
/*SignatureInformation defined:
>>>>>>> v0.0.4
 * Represents the signature of something callable. A signature
 * can have a label, like a function-name, a doc-comment, and
 * a set of parameters.
 */
type SignatureInformation struct {

<<<<<<< HEAD
	/** Label defined:
=======
	/*Label defined:
>>>>>>> v0.0.4
	 * The label of this signature. Will be shown in
	 * the UI.
	 */
	Label string `json:"label"`

<<<<<<< HEAD
	/** Documentation defined:
=======
	/*Documentation defined:
>>>>>>> v0.0.4
	 * The human-readable doc-comment of this signature. Will be shown
	 * in the UI but can be omitted.
	 */
	Documentation string `json:"documentation,omitempty"` // string | MarkupContent

<<<<<<< HEAD
	/** Parameters defined:
=======
	/*Parameters defined:
>>>>>>> v0.0.4
	 * The parameters of this signature.
	 */
	Parameters []ParameterInformation `json:"parameters,omitempty"`
}

<<<<<<< HEAD
// SignatureHelp is:
/**
=======
/*SignatureHelp defined:
>>>>>>> v0.0.4
 * Signature help represents the signature of something
 * callable. There can be multiple signature but only one
 * active and only one active parameter.
 */
type SignatureHelp struct {

<<<<<<< HEAD
	/** Signatures defined:
=======
	/*Signatures defined:
>>>>>>> v0.0.4
	 * One or more signatures.
	 */
	Signatures []SignatureInformation `json:"signatures"`

<<<<<<< HEAD
	/** ActiveSignature defined:
=======
	/*ActiveSignature defined:
>>>>>>> v0.0.4
	 * The active signature. Set to `null` if no
	 * signatures exist.
	 */
	ActiveSignature float64 `json:"activeSignature"`

<<<<<<< HEAD
	/** ActiveParameter defined:
=======
	/*ActiveParameter defined:
>>>>>>> v0.0.4
	 * The active parameter of the active signature. Set to `null`
	 * if the active signature has no parameters.
	 */
	ActiveParameter float64 `json:"activeParameter"`
}

<<<<<<< HEAD
// ReferenceContext is:
/**
=======
/*ReferenceContext defined:
>>>>>>> v0.0.4
 * Value-object that contains additional information when
 * requesting references.
 */
type ReferenceContext struct {

<<<<<<< HEAD
	/** IncludeDeclaration defined:
=======
	/*IncludeDeclaration defined:
>>>>>>> v0.0.4
	 * Include the declaration of the current symbol.
	 */
	IncludeDeclaration bool `json:"includeDeclaration"`
}

<<<<<<< HEAD
// DocumentHighlight is:
/**
=======
/*DocumentHighlight defined:
>>>>>>> v0.0.4
 * A document highlight is a range inside a text document which deserves
 * special attention. Usually a document highlight is visualized by changing
 * the background color of its range.
 */
type DocumentHighlight struct {

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range this highlight applies to.
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** Kind defined:
=======
	/*Kind defined:
>>>>>>> v0.0.4
	 * The highlight kind, default is [text](#DocumentHighlightKind.Text).
	 */
	Kind *DocumentHighlightKind `json:"kind,omitempty"`
}

<<<<<<< HEAD
// SymbolInformation is:
/**
=======
/*SymbolInformation defined:
>>>>>>> v0.0.4
 * Represents information about programming constructs like variables, classes,
 * interfaces etc.
 */
type SymbolInformation struct {

<<<<<<< HEAD
	/** Name defined:
=======
	/*Name defined:
>>>>>>> v0.0.4
	 * The name of this symbol.
	 */
	Name string `json:"name"`

<<<<<<< HEAD
	/** Kind defined:
=======
	/*Kind defined:
>>>>>>> v0.0.4
	 * The kind of this symbol.
	 */
	Kind SymbolKind `json:"kind"`

<<<<<<< HEAD
	/** Deprecated defined:
=======
	/*Deprecated defined:
>>>>>>> v0.0.4
	 * Indicates if this symbol is deprecated.
	 */
	Deprecated bool `json:"deprecated,omitempty"`

<<<<<<< HEAD
	/** Location defined:
=======
	/*Location defined:
>>>>>>> v0.0.4
	 * The location of this symbol. The location's range is used by a tool
	 * to reveal the location in the editor. If the symbol is selected in the
	 * tool the range's start information is used to position the cursor. So
	 * the range usually spans more than the actual symbol's name and does
	 * normally include thinks like visibility modifiers.
	 *
	 * The range doesn't have to denote a node range in the sense of a abstract
	 * syntax tree. It can therefore not be used to re-construct a hierarchy of
	 * the symbols.
	 */
	Location Location `json:"location"`

<<<<<<< HEAD
	/** ContainerName defined:
=======
	/*ContainerName defined:
>>>>>>> v0.0.4
	 * The name of the symbol containing this symbol. This information is for
	 * user interface purposes (e.g. to render a qualifier in the user interface
	 * if necessary). It can't be used to re-infer a hierarchy for the document
	 * symbols.
	 */
	ContainerName string `json:"containerName,omitempty"`
}

<<<<<<< HEAD
// DocumentSymbol is:
/**
=======
/*DocumentSymbol defined:
>>>>>>> v0.0.4
 * Represents programming constructs like variables, classes, interfaces etc.
 * that appear in a document. Document symbols can be hierarchical and they
 * have two ranges: one that encloses its definition and one that points to
 * its most interesting range, e.g. the range of an identifier.
 */
type DocumentSymbol struct {

<<<<<<< HEAD
	/** Name defined:
=======
	/*Name defined:
>>>>>>> v0.0.4
	 * The name of this symbol. Will be displayed in the user interface and therefore must not be
	 * an empty string or a string only consisting of white spaces.
	 */
	Name string `json:"name"`

<<<<<<< HEAD
	/** Detail defined:
=======
	/*Detail defined:
>>>>>>> v0.0.4
	 * More detail for this symbol, e.g the signature of a function.
	 */
	Detail string `json:"detail,omitempty"`

<<<<<<< HEAD
	/** Kind defined:
=======
	/*Kind defined:
>>>>>>> v0.0.4
	 * The kind of this symbol.
	 */
	Kind SymbolKind `json:"kind"`

<<<<<<< HEAD
	/** Deprecated defined:
=======
	/*Deprecated defined:
>>>>>>> v0.0.4
	 * Indicates if this symbol is deprecated.
	 */
	Deprecated bool `json:"deprecated,omitempty"`

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range enclosing this symbol not including leading/trailing whitespace but everything else
	 * like comments. This information is typically used to determine if the the clients cursor is
	 * inside the symbol to reveal in the symbol in the UI.
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** SelectionRange defined:
=======
	/*SelectionRange defined:
>>>>>>> v0.0.4
	 * The range that should be selected and revealed when this symbol is being picked, e.g the name of a function.
	 * Must be contained by the the `range`.
	 */
	SelectionRange Range `json:"selectionRange"`

<<<<<<< HEAD
	/** Children defined:
=======
	/*Children defined:
>>>>>>> v0.0.4
	 * Children of this symbol, e.g. properties of a class.
	 */
	Children []DocumentSymbol `json:"children,omitempty"`
}

<<<<<<< HEAD
// DocumentSymbolParams is:
/**
=======
/*DocumentSymbolParams defined:
>>>>>>> v0.0.4
 * Parameters for a [DocumentSymbolRequest](#DocumentSymbolRequest).
 */
type DocumentSymbolParams struct {

<<<<<<< HEAD
	/** TextDocument defined:
=======
	/*TextDocument defined:
>>>>>>> v0.0.4
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

<<<<<<< HEAD
// WorkspaceSymbolParams is:
/**
=======
/*WorkspaceSymbolParams defined:
>>>>>>> v0.0.4
 * The parameters of a [WorkspaceSymbolRequest](#WorkspaceSymbolRequest).
 */
type WorkspaceSymbolParams struct {

<<<<<<< HEAD
	/** Query defined:
=======
	/*Query defined:
>>>>>>> v0.0.4
	 * A non-empty query string
	 */
	Query string `json:"query"`
}

<<<<<<< HEAD
// CodeActionContext is:
/**
=======
/*CodeActionContext defined:
>>>>>>> v0.0.4
 * Contains additional diagnostic information about the context in which
 * a [code action](#CodeActionProvider.provideCodeActions) is run.
 */
type CodeActionContext struct {

<<<<<<< HEAD
	/** Diagnostics defined:
=======
	/*Diagnostics defined:
>>>>>>> v0.0.4
	 * An array of diagnostics known on the client side overlapping the range provided to the
	 * `textDocument/codeAction` request. They are provied so that the server knows which
	 * errors are currently presented to the user for the given range. There is no guarantee
	 * that these accurately reflect the error state of the resource. The primary parameter
	 * to compute code actions is the provided range.
	 */
	Diagnostics []Diagnostic `json:"diagnostics"`

<<<<<<< HEAD
	/** Only defined:
=======
	/*Only defined:
>>>>>>> v0.0.4
	 * Requested kind of actions to return.
	 *
	 * Actions not of this kind are filtered out by the client before being shown. So servers
	 * can omit computing them.
	 */
	Only []CodeActionKind `json:"only,omitempty"`
}

<<<<<<< HEAD
// CodeAction is:
/**
=======
/*CodeAction defined:
>>>>>>> v0.0.4
 * A code action represents a change that can be performed in code, e.g. to fix a problem or
 * to refactor code.
 *
 * A CodeAction must set either `edit` and/or a `command`. If both are supplied, the `edit` is applied first, then the `command` is executed.
 */
type CodeAction struct {

<<<<<<< HEAD
	/** Title defined:
=======
	/*Title defined:
>>>>>>> v0.0.4
	 * A short, human-readable, title for this code action.
	 */
	Title string `json:"title"`

<<<<<<< HEAD
	/** Kind defined:
=======
	/*Kind defined:
>>>>>>> v0.0.4
	 * The kind of the code action.
	 *
	 * Used to filter code actions.
	 */
	Kind CodeActionKind `json:"kind,omitempty"`

<<<<<<< HEAD
	/** Diagnostics defined:
=======
	/*Diagnostics defined:
>>>>>>> v0.0.4
	 * The diagnostics that this code action resolves.
	 */
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`

<<<<<<< HEAD
	/** Edit defined:
=======
	/*Edit defined:
>>>>>>> v0.0.4
	 * The workspace edit this code action performs.
	 */
	Edit *WorkspaceEdit `json:"edit,omitempty"`

<<<<<<< HEAD
	/** Command defined:
=======
	/*Command defined:
>>>>>>> v0.0.4
	 * A command this code action executes. If a code action
	 * provides a edit and a command, first the edit is
	 * executed and then the command.
	 */
	Command *Command `json:"command,omitempty"`
}

<<<<<<< HEAD
// CodeLens is:
/**
=======
/*CodeLens defined:
>>>>>>> v0.0.4
 * A code lens represents a [command](#Command) that should be shown along with
 * source text, like the number of references, a way to run tests, etc.
 *
 * A code lens is _unresolved_ when no command is associated to it. For performance
 * reasons the creation of a code lens and resolving should be done to two stages.
 */
type CodeLens struct {

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range in which this code lens is valid. Should only span a single line.
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** Command defined:
=======
	/*Command defined:
>>>>>>> v0.0.4
	 * The command this code lens represents.
	 */
	Command *Command `json:"command,omitempty"`

<<<<<<< HEAD
	/** Data defined:
=======
	/*Data defined:
>>>>>>> v0.0.4
	 * An data entry field that is preserved on a code lens item between
	 * a [CodeLensRequest](#CodeLensRequest) and a [CodeLensResolveRequest]
	 * (#CodeLensResolveRequest)
	 */
	Data interface{} `json:"data,omitempty"`
}

<<<<<<< HEAD
// FormattingOptions is:
/**
=======
/*FormattingOptions defined:
>>>>>>> v0.0.4
 * Value-object describing what options formatting should use.
 */
type FormattingOptions struct {

<<<<<<< HEAD
	/** TabSize defined:
=======
	/*TabSize defined:
>>>>>>> v0.0.4
	 * Size of a tab in spaces.
	 */
	TabSize float64 `json:"tabSize"`

<<<<<<< HEAD
	/** InsertSpaces defined:
=======
	/*InsertSpaces defined:
>>>>>>> v0.0.4
	 * Prefer spaces over tabs.
	 */
	InsertSpaces bool `json:"insertSpaces"`

<<<<<<< HEAD
	/** TrimTrailingWhitespace defined:
=======
	/*TrimTrailingWhitespace defined:
>>>>>>> v0.0.4
	 * Trim trailing whitespaces on a line.
	 */
	TrimTrailingWhitespace bool `json:"trimTrailingWhitespace,omitempty"`

<<<<<<< HEAD
	/** InsertFinalNewline defined:
=======
	/*InsertFinalNewline defined:
>>>>>>> v0.0.4
	 * Insert a newline character at the end of the file if one does not exist.
	 */
	InsertFinalNewline bool `json:"insertFinalNewline,omitempty"`

<<<<<<< HEAD
	/** TrimFinalNewlines defined:
=======
	/*TrimFinalNewlines defined:
>>>>>>> v0.0.4
	 * Trim all newlines after the final newline at the end of the file.
	 */
	TrimFinalNewlines bool `json:"trimFinalNewlines,omitempty"`

<<<<<<< HEAD
	/** Key defined:
=======
	/*Key defined:
>>>>>>> v0.0.4
	 * Signature for further properties.
	 */
	Key map[string]bool `json:"key"` // [key: string]: boolean | number | string | undefined;
}

<<<<<<< HEAD
// DocumentLink is:
/**
=======
/*DocumentLink defined:
>>>>>>> v0.0.4
 * A document link is a range in a text document that links to an internal or external resource, like another
 * text document or a web site.
 */
type DocumentLink struct {

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range this link applies to.
	 */
	Range Range `json:"range"`

<<<<<<< HEAD
	/** Target defined:
=======
	/*Target defined:
>>>>>>> v0.0.4
	 * The uri this link points to.
	 */
	Target string `json:"target,omitempty"`

<<<<<<< HEAD
	/** Data defined:
=======
	/*Data defined:
>>>>>>> v0.0.4
	 * A data entry field that is preserved on a document link between a
	 * DocumentLinkRequest and a DocumentLinkResolveRequest.
	 */
	Data interface{} `json:"data,omitempty"`
}

<<<<<<< HEAD
// TextDocument is:
/**
=======
/*TextDocument defined:
>>>>>>> v0.0.4
 * A simple text document. Not to be implemented.
 */
type TextDocument struct {

<<<<<<< HEAD
	/** URI defined:
=======
	/*URI defined:
>>>>>>> v0.0.4
	 * The associated URI for this document. Most documents have the __file__-scheme, indicating that they
	 * represent files on disk. However, some documents may have other schemes indicating that they are not
	 * available on disk.
	 *
	 * @readonly
	 */
	URI string `json:"uri"`

<<<<<<< HEAD
	/** LanguageID defined:
=======
	/*LanguageID defined:
>>>>>>> v0.0.4
	 * The identifier of the language associated with this document.
	 *
	 * @readonly
	 */
	LanguageID string `json:"languageId"`

<<<<<<< HEAD
	/** Version defined:
=======
	/*Version defined:
>>>>>>> v0.0.4
	 * The version number of this document (it will increase after each
	 * change, including undo/redo).
	 *
	 * @readonly
	 */
	Version float64 `json:"version"`

<<<<<<< HEAD
	/** LineCount defined:
=======
	/*LineCount defined:
>>>>>>> v0.0.4
	 * The number of lines in this document.
	 *
	 * @readonly
	 */
	LineCount float64 `json:"lineCount"`
}

<<<<<<< HEAD
// TextDocumentChangeEvent is:
/**
=======
/*TextDocumentChangeEvent defined:
>>>>>>> v0.0.4
 * Event to signal changes to a simple text document.
 */
type TextDocumentChangeEvent struct {

<<<<<<< HEAD
	/** Document defined:
=======
	/*Document defined:
>>>>>>> v0.0.4
	 * The document that has changed.
	 */
	Document TextDocument `json:"document"`
}

<<<<<<< HEAD
// TextDocumentWillSaveEvent is:
type TextDocumentWillSaveEvent struct {

	/** Document defined:
=======
// TextDocumentWillSaveEvent is
type TextDocumentWillSaveEvent struct {

	/*Document defined:
>>>>>>> v0.0.4
	 * The document that will be saved
	 */
	Document TextDocument `json:"document"`

<<<<<<< HEAD
	/** Reason defined:
=======
	/*Reason defined:
>>>>>>> v0.0.4
	 * The reason why save was triggered.
	 */
	Reason TextDocumentSaveReason `json:"reason"`
}

<<<<<<< HEAD
// TextDocumentContentChangeEvent is:
/**
=======
/*TextDocumentContentChangeEvent defined:
>>>>>>> v0.0.4
 * An event describing a change to a text document. If range and rangeLength are omitted
 * the new text is considered to be the full content of the document.
 */
type TextDocumentContentChangeEvent struct {

<<<<<<< HEAD
	/** Range defined:
=======
	/*Range defined:
>>>>>>> v0.0.4
	 * The range of the document that changed.
	 */
	Range *Range `json:"range,omitempty"`

<<<<<<< HEAD
	/** RangeLength defined:
=======
	/*RangeLength defined:
>>>>>>> v0.0.4
	 * The length of the range that got replaced.
	 */
	RangeLength float64 `json:"rangeLength,omitempty"`

<<<<<<< HEAD
	/** Text defined:
=======
	/*Text defined:
>>>>>>> v0.0.4
	 * The new text of the document.
	 */
	Text string `json:"text"`
}

// FoldingRangeKind defines constants
type FoldingRangeKind string

// SelectionRangeKind defines constants
type SelectionRangeKind string

// ResourceOperationKind defines constants
type ResourceOperationKind string

// FailureHandlingKind defines constants
type FailureHandlingKind string

// TextDocumentSyncKind defines constants
type TextDocumentSyncKind float64

// InitializeError defines constants
type InitializeError float64

// MessageType defines constants
type MessageType float64

// FileChangeType defines constants
type FileChangeType float64

// WatchKind defines constants
type WatchKind float64

// CompletionTriggerKind defines constants
type CompletionTriggerKind float64

// DiagnosticSeverity defines constants
type DiagnosticSeverity float64

// DiagnosticTag defines constants
type DiagnosticTag float64

// MarkupKind defines constants
type MarkupKind string

// CompletionItemKind defines constants
type CompletionItemKind float64

// InsertTextFormat defines constants
type InsertTextFormat float64

// DocumentHighlightKind defines constants
type DocumentHighlightKind float64

// SymbolKind defines constants
type SymbolKind float64

// CodeActionKind defines constants
type CodeActionKind string

// TextDocumentSaveReason defines constants
type TextDocumentSaveReason float64

const (

<<<<<<< HEAD
	// Comment is:
	/**
=======
	/*Comment defined:
>>>>>>> v0.0.4
	 * Folding range for a comment
	 */
	Comment FoldingRangeKind = "comment"

<<<<<<< HEAD
	// Imports is:
	/**
=======
	/*Imports defined:
>>>>>>> v0.0.4
	 * Folding range for a imports or includes
	 */
	Imports FoldingRangeKind = "imports"

<<<<<<< HEAD
	// Region is:
	/**
=======
	/*Region defined:
>>>>>>> v0.0.4
	 * Folding range for a region (e.g. `#region`)
	 */
	Region FoldingRangeKind = "region"

<<<<<<< HEAD
	// Empty is:
	/**
=======
	/*Empty defined:
>>>>>>> v0.0.4
	 * Empty Kind.
	 */
	Empty SelectionRangeKind = ""

<<<<<<< HEAD
	// Statement is:
	/**
=======
	/*Statement defined:
>>>>>>> v0.0.4
	 * The statment kind, its value is `statement`, possible extensions can be
	 * `statement.if` etc
	 */
	Statement SelectionRangeKind = "statement"

<<<<<<< HEAD
	// Declaration is:
	/**
=======
	/*Declaration defined:
>>>>>>> v0.0.4
	 * The declaration kind, its value is `declaration`, possible extensions can be
	 * `declaration.function`, `declaration.class` etc.
	 */
	Declaration SelectionRangeKind = "declaration"

<<<<<<< HEAD
	// Create is:
	/**
=======
	/*Create defined:
>>>>>>> v0.0.4
	 * Supports creating new files and folders.
	 */
	Create ResourceOperationKind = "create"

<<<<<<< HEAD
	// Rename is:
	/**
=======
	/*Rename defined:
>>>>>>> v0.0.4
	 * Supports renaming existing files and folders.
	 */
	Rename ResourceOperationKind = "rename"

<<<<<<< HEAD
	// Delete is:
	/**
=======
	/*Delete defined:
>>>>>>> v0.0.4
	 * Supports deleting existing files and folders.
	 */
	Delete ResourceOperationKind = "delete"

<<<<<<< HEAD
	// Abort is:
	/**
=======
	/*Abort defined:
>>>>>>> v0.0.4
	 * Applying the workspace change is simply aborted if one of the changes provided
	 * fails. All operations executed before the failing operation stay executed.
	 */
	Abort FailureHandlingKind = "abort"

<<<<<<< HEAD
	// Transactional is:
	/**
=======
	/*Transactional defined:
>>>>>>> v0.0.4
	 * All operations are executed transactional. That means they either all
	 * succeed or no changes at all are applied to the workspace.
	 */
	Transactional FailureHandlingKind = "transactional"

<<<<<<< HEAD
	// TextOnlyTransactional is:
	/**
=======
	/*TextOnlyTransactional defined:
>>>>>>> v0.0.4
	 * If the workspace edit contains only textual file changes they are executed transactional.
	 * If resource changes (create, rename or delete file) are part of the change the failure
	 * handling startegy is abort.
	 */
	TextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"

<<<<<<< HEAD
	// Undo is:
	/**
=======
	/*Undo defined:
>>>>>>> v0.0.4
	 * The client tries to undo the operations already executed. But there is no
	 * guaruntee that this is succeeding.
	 */
	Undo FailureHandlingKind = "undo"

<<<<<<< HEAD
	// None is:
	/**
=======
	/*None defined:
>>>>>>> v0.0.4
	 * Documents should not be synced at all.
	 */
	None TextDocumentSyncKind = 0

<<<<<<< HEAD
	// Full is:
	/**
=======
	/*Full defined:
>>>>>>> v0.0.4
	 * Documents are synced by always sending the full content
	 * of the document.
	 */
	Full TextDocumentSyncKind = 1

<<<<<<< HEAD
	// Incremental is:
	/**
=======
	/*Incremental defined:
>>>>>>> v0.0.4
	 * Documents are synced by sending the full content on open.
	 * After that only incremental updates to the document are
	 * send.
	 */
	Incremental TextDocumentSyncKind = 2

<<<<<<< HEAD
	// UnknownProtocolVersion is:
	/**
=======
	/*UnknownProtocolVersion defined:
>>>>>>> v0.0.4
	 * If the protocol version provided by the client can't be handled by the server.
	 * @deprecated This initialize error got replaced by client capabilities. There is
	 * no version handshake in version 3.0x
	 */
	UnknownProtocolVersion InitializeError = 1

<<<<<<< HEAD
	// Error is:
	/**
=======
	/*Error defined:
>>>>>>> v0.0.4
	 * An error message.
	 */
	Error MessageType = 1

<<<<<<< HEAD
	// Warning is:
	/**
=======
	/*Warning defined:
>>>>>>> v0.0.4
	 * A warning message.
	 */
	Warning MessageType = 2

<<<<<<< HEAD
	// Info is:
	/**
=======
	/*Info defined:
>>>>>>> v0.0.4
	 * An information message.
	 */
	Info MessageType = 3

<<<<<<< HEAD
	// Log is:
	/**
=======
	/*Log defined:
>>>>>>> v0.0.4
	 * A log message.
	 */
	Log MessageType = 4

<<<<<<< HEAD
	// Created is:
	/**
=======
	/*Created defined:
>>>>>>> v0.0.4
	 * The file got created.
	 */
	Created FileChangeType = 1

<<<<<<< HEAD
	// Changed is:
	/**
=======
	/*Changed defined:
>>>>>>> v0.0.4
	 * The file got changed.
	 */
	Changed FileChangeType = 2

<<<<<<< HEAD
	// Deleted is:
	/**
=======
	/*Deleted defined:
>>>>>>> v0.0.4
	 * The file got deleted.
	 */
	Deleted FileChangeType = 3

<<<<<<< HEAD
	// Change is:
	/**
=======
	/*Change defined:
>>>>>>> v0.0.4
	 * Interested in change events
	 */
	Change WatchKind = 2

<<<<<<< HEAD
	// Invoked is:
	/**
=======
	/*Invoked defined:
>>>>>>> v0.0.4
	 * Completion was triggered by typing an identifier (24x7 code
	 * complete), manual invocation (e.g Ctrl+Space) or via API.
	 */
	Invoked CompletionTriggerKind = 1

<<<<<<< HEAD
	// TriggerCharacter is:
	/**
=======
	/*TriggerCharacter defined:
>>>>>>> v0.0.4
	 * Completion was triggered by a trigger character specified by
	 * the `triggerCharacters` properties of the `CompletionRegistrationOptions`.
	 */
	TriggerCharacter CompletionTriggerKind = 2

<<<<<<< HEAD
	// TriggerForIncompleteCompletions is:
	/**
=======
	/*TriggerForIncompleteCompletions defined:
>>>>>>> v0.0.4
	 * Completion was re-triggered as current completion list is incomplete
	 */
	TriggerForIncompleteCompletions CompletionTriggerKind = 3

<<<<<<< HEAD
	// SeverityError is:
	/**
=======
	/*SeverityError defined:
>>>>>>> v0.0.4
	 * Reports an error.
	 */
	SeverityError DiagnosticSeverity = 1

<<<<<<< HEAD
	// SeverityWarning is:
	/**
=======
	/*SeverityWarning defined:
>>>>>>> v0.0.4
	 * Reports a warning.
	 */
	SeverityWarning DiagnosticSeverity = 2

<<<<<<< HEAD
	// SeverityInformation is:
	/**
=======
	/*SeverityInformation defined:
>>>>>>> v0.0.4
	 * Reports an information.
	 */
	SeverityInformation DiagnosticSeverity = 3

<<<<<<< HEAD
	// SeverityHint is:
	/**
=======
	/*SeverityHint defined:
>>>>>>> v0.0.4
	 * Reports a hint.
	 */
	SeverityHint DiagnosticSeverity = 4

<<<<<<< HEAD
	// Unnecessary is:
	/**
=======
	/*Unnecessary defined:
>>>>>>> v0.0.4
	 * Unused or unnecessary code.
	 *
	 * Clients are allowed to render diagnostics with this tag faded out instead of having
	 * an error squiggle.
	 */
	Unnecessary DiagnosticTag = 1

<<<<<<< HEAD
	// PlainText is:
	/**
=======
	/*PlainText defined:
>>>>>>> v0.0.4
	 * Plain text is supported as a content format
	 */
	PlainText MarkupKind = "plaintext"

<<<<<<< HEAD
	// Markdown is:
	/**
=======
	/*Markdown defined:
>>>>>>> v0.0.4
	 * Markdown is supported as a content format
	 */
	Markdown MarkupKind = "markdown"

<<<<<<< HEAD
	// TextCompletion is:
	TextCompletion CompletionItemKind = 1

	// MethodCompletion is:
	MethodCompletion CompletionItemKind = 2

	// FunctionCompletion is:
	FunctionCompletion CompletionItemKind = 3

	// ConstructorCompletion is:
	ConstructorCompletion CompletionItemKind = 4

	// FieldCompletion is:
	FieldCompletion CompletionItemKind = 5

	// VariableCompletion is:
	VariableCompletion CompletionItemKind = 6

	// ClassCompletion is:
	ClassCompletion CompletionItemKind = 7

	// InterfaceCompletion is:
	InterfaceCompletion CompletionItemKind = 8

	// ModuleCompletion is:
	ModuleCompletion CompletionItemKind = 9

	// PropertyCompletion is:
	PropertyCompletion CompletionItemKind = 10

	// UnitCompletion is:
	UnitCompletion CompletionItemKind = 11

	// ValueCompletion is:
	ValueCompletion CompletionItemKind = 12

	// EnumCompletion is:
	EnumCompletion CompletionItemKind = 13

	// KeywordCompletion is:
	KeywordCompletion CompletionItemKind = 14

	// SnippetCompletion is:
	SnippetCompletion CompletionItemKind = 15

	// ColorCompletion is:
	ColorCompletion CompletionItemKind = 16

	// FileCompletion is:
	FileCompletion CompletionItemKind = 17

	// ReferenceCompletion is:
	ReferenceCompletion CompletionItemKind = 18

	// FolderCompletion is:
	FolderCompletion CompletionItemKind = 19

	// EnumMemberCompletion is:
	EnumMemberCompletion CompletionItemKind = 20

	// ConstantCompletion is:
	ConstantCompletion CompletionItemKind = 21

	// StructCompletion is:
	StructCompletion CompletionItemKind = 22

	// EventCompletion is:
	EventCompletion CompletionItemKind = 23

	// OperatorCompletion is:
	OperatorCompletion CompletionItemKind = 24

	// TypeParameterCompletion is:
	TypeParameterCompletion CompletionItemKind = 25

	// PlainTextTextFormat is:
	/**
=======
	// TextCompletion is
	TextCompletion CompletionItemKind = 1

	// MethodCompletion is
	MethodCompletion CompletionItemKind = 2

	// FunctionCompletion is
	FunctionCompletion CompletionItemKind = 3

	// ConstructorCompletion is
	ConstructorCompletion CompletionItemKind = 4

	// FieldCompletion is
	FieldCompletion CompletionItemKind = 5

	// VariableCompletion is
	VariableCompletion CompletionItemKind = 6

	// ClassCompletion is
	ClassCompletion CompletionItemKind = 7

	// InterfaceCompletion is
	InterfaceCompletion CompletionItemKind = 8

	// ModuleCompletion is
	ModuleCompletion CompletionItemKind = 9

	// PropertyCompletion is
	PropertyCompletion CompletionItemKind = 10

	// UnitCompletion is
	UnitCompletion CompletionItemKind = 11

	// ValueCompletion is
	ValueCompletion CompletionItemKind = 12

	// EnumCompletion is
	EnumCompletion CompletionItemKind = 13

	// KeywordCompletion is
	KeywordCompletion CompletionItemKind = 14

	// SnippetCompletion is
	SnippetCompletion CompletionItemKind = 15

	// ColorCompletion is
	ColorCompletion CompletionItemKind = 16

	// FileCompletion is
	FileCompletion CompletionItemKind = 17

	// ReferenceCompletion is
	ReferenceCompletion CompletionItemKind = 18

	// FolderCompletion is
	FolderCompletion CompletionItemKind = 19

	// EnumMemberCompletion is
	EnumMemberCompletion CompletionItemKind = 20

	// ConstantCompletion is
	ConstantCompletion CompletionItemKind = 21

	// StructCompletion is
	StructCompletion CompletionItemKind = 22

	// EventCompletion is
	EventCompletion CompletionItemKind = 23

	// OperatorCompletion is
	OperatorCompletion CompletionItemKind = 24

	// TypeParameterCompletion is
	TypeParameterCompletion CompletionItemKind = 25

	/*PlainTextTextFormat defined:
>>>>>>> v0.0.4
	 * The primary text to be inserted is treated as a plain string.
	 */
	PlainTextTextFormat InsertTextFormat = 1

<<<<<<< HEAD
	// SnippetTextFormat is:
	/**
=======
	/*SnippetTextFormat defined:
>>>>>>> v0.0.4
	 * The primary text to be inserted is treated as a snippet.
	 *
	 * A snippet can define tab stops and placeholders with `$1`, `$2`
	 * and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	 * the end of the snippet. Placeholders with equal identifiers are linked,
	 * that is typing in one will update others too.
	 *
	 * See also: https://github.com/Microsoft/vscode/blob/master/src/vs/editor/contrib/snippet/common/snippet.md
	 */
	SnippetTextFormat InsertTextFormat = 2

<<<<<<< HEAD
	// Text is:
	/**
=======
	/*Text defined:
>>>>>>> v0.0.4
	 * A textual occurrence.
	 */
	Text DocumentHighlightKind = 1

<<<<<<< HEAD
	// Read is:
	/**
=======
	/*Read defined:
>>>>>>> v0.0.4
	 * Read-access of a symbol, like reading a variable.
	 */
	Read DocumentHighlightKind = 2

<<<<<<< HEAD
	// Write is:
	/**
=======
	/*Write defined:
>>>>>>> v0.0.4
	 * Write-access of a symbol, like writing to a variable.
	 */
	Write DocumentHighlightKind = 3

<<<<<<< HEAD
	// File is:
	File SymbolKind = 1

	// Module is:
	Module SymbolKind = 2

	// Namespace is:
	Namespace SymbolKind = 3

	// Package is:
	Package SymbolKind = 4

	// Class is:
	Class SymbolKind = 5

	// Method is:
	Method SymbolKind = 6

	// Property is:
	Property SymbolKind = 7

	// Field is:
	Field SymbolKind = 8

	// Constructor is:
	Constructor SymbolKind = 9

	// Enum is:
	Enum SymbolKind = 10

	// Interface is:
	Interface SymbolKind = 11

	// Function is:
	Function SymbolKind = 12

	// Variable is:
	Variable SymbolKind = 13

	// Constant is:
	Constant SymbolKind = 14

	// String is:
	String SymbolKind = 15

	// Number is:
	Number SymbolKind = 16

	// Boolean is:
	Boolean SymbolKind = 17

	// Array is:
	Array SymbolKind = 18

	// Object is:
	Object SymbolKind = 19

	// Key is:
	Key SymbolKind = 20

	// Null is:
	Null SymbolKind = 21

	// EnumMember is:
	EnumMember SymbolKind = 22

	// Struct is:
	Struct SymbolKind = 23

	// Event is:
	Event SymbolKind = 24

	// Operator is:
	Operator SymbolKind = 25

	// TypeParameter is:
	TypeParameter SymbolKind = 26

	// QuickFix is:
	/**
=======
	// File is
	File SymbolKind = 1

	// Module is
	Module SymbolKind = 2

	// Namespace is
	Namespace SymbolKind = 3

	// Package is
	Package SymbolKind = 4

	// Class is
	Class SymbolKind = 5

	// Method is
	Method SymbolKind = 6

	// Property is
	Property SymbolKind = 7

	// Field is
	Field SymbolKind = 8

	// Constructor is
	Constructor SymbolKind = 9

	// Enum is
	Enum SymbolKind = 10

	// Interface is
	Interface SymbolKind = 11

	// Function is
	Function SymbolKind = 12

	// Variable is
	Variable SymbolKind = 13

	// Constant is
	Constant SymbolKind = 14

	// String is
	String SymbolKind = 15

	// Number is
	Number SymbolKind = 16

	// Boolean is
	Boolean SymbolKind = 17

	// Array is
	Array SymbolKind = 18

	// Object is
	Object SymbolKind = 19

	// Key is
	Key SymbolKind = 20

	// Null is
	Null SymbolKind = 21

	// EnumMember is
	EnumMember SymbolKind = 22

	// Struct is
	Struct SymbolKind = 23

	// Event is
	Event SymbolKind = 24

	// Operator is
	Operator SymbolKind = 25

	// TypeParameter is
	TypeParameter SymbolKind = 26

	/*QuickFix defined:
>>>>>>> v0.0.4
	 * Base kind for quickfix actions: 'quickfix'
	 */
	QuickFix CodeActionKind = "quickfix"

<<<<<<< HEAD
	// Refactor is:
	/**
=======
	/*Refactor defined:
>>>>>>> v0.0.4
	 * Base kind for refactoring actions: 'refactor'
	 */
	Refactor CodeActionKind = "refactor"

<<<<<<< HEAD
	// RefactorExtract is:
	/**
=======
	/*RefactorExtract defined:
>>>>>>> v0.0.4
	 * Base kind for refactoring extraction actions: 'refactor.extract'
	 *
	 * Example extract actions:
	 *
	 * - Extract method
	 * - Extract function
	 * - Extract variable
	 * - Extract interface from class
	 * - ...
	 */
	RefactorExtract CodeActionKind = "refactor.extract"

<<<<<<< HEAD
	// RefactorInline is:
	/**
=======
	/*RefactorInline defined:
>>>>>>> v0.0.4
	 * Base kind for refactoring inline actions: 'refactor.inline'
	 *
	 * Example inline actions:
	 *
	 * - Inline function
	 * - Inline variable
	 * - Inline constant
	 * - ...
	 */
	RefactorInline CodeActionKind = "refactor.inline"

<<<<<<< HEAD
	// RefactorRewrite is:
	/**
=======
	/*RefactorRewrite defined:
>>>>>>> v0.0.4
	 * Base kind for refactoring rewrite actions: 'refactor.rewrite'
	 *
	 * Example rewrite actions:
	 *
	 * - Convert JavaScript function to class
	 * - Add or remove parameter
	 * - Encapsulate field
	 * - Make method static
	 * - Move method to base class
	 * - ...
	 */
	RefactorRewrite CodeActionKind = "refactor.rewrite"

<<<<<<< HEAD
	// Source is:
	/**
=======
	/*Source defined:
>>>>>>> v0.0.4
	 * Base kind for source actions: `source`
	 *
	 * Source code actions apply to the entire file.
	 */
	Source CodeActionKind = "source"

<<<<<<< HEAD
	// SourceOrganizeImports is:
	/**
=======
	/*SourceOrganizeImports defined:
>>>>>>> v0.0.4
	 * Base kind for an organize imports source action: `source.organizeImports`
	 */
	SourceOrganizeImports CodeActionKind = "source.organizeImports"

<<<<<<< HEAD
	// Manual is:
	/**
=======
	/*Manual defined:
>>>>>>> v0.0.4
	 * Manually triggered, e.g. by the user pressing save, by starting debugging,
	 * or by an API call.
	 */
	Manual TextDocumentSaveReason = 1

<<<<<<< HEAD
	// AfterDelay is:
	/**
=======
	/*AfterDelay defined:
>>>>>>> v0.0.4
	 * Automatic after a delay.
	 */
	AfterDelay TextDocumentSaveReason = 2

<<<<<<< HEAD
	// FocusOut is:
	/**
=======
	/*FocusOut defined:
>>>>>>> v0.0.4
	 * When the editor lost focus.
	 */
	FocusOut TextDocumentSaveReason = 3
)

// DocumentFilter is a type
/**
 * A document filter denotes a document by different properties like
 * the [language](#TextDocument.languageId), the [scheme](#Uri.scheme) of
 * its resource, or a glob-pattern that is applied to the [path](#TextDocument.fileName).
 *
 * Glob patterns can have the following syntax:
 * - `*` to match one or more characters in a path segment
 * - `?` to match on one character in a path segment
 * - `**` to match any number of path segments, including none
 * - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
 * - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
 * - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
 *
 * @sample A language filter that applies to typescript files on disk: `{ language: 'typescript', scheme: 'file' }`
 * @sample A language filter that applies to all package.json paths: `{ language: 'json', pattern: '**package.json' }`
 */
type DocumentFilter struct {

<<<<<<< HEAD
	/** Language defined: A language id, like `typescript`. */
	Language string `json:"language,omitempty"`

	/** Scheme defined: A Uri [scheme](#Uri.scheme), like `file` or `untitled`. */
	Scheme string `json:"scheme,omitempty"`

	/** Pattern defined: A glob pattern, like `*.{ts,js}`. */
=======
	/*Language defined: A language id, like `typescript`. */
	Language string `json:"language,omitempty"`

	/*Scheme defined: A Uri [scheme](#Uri.scheme), like `file` or `untitled`. */
	Scheme string `json:"scheme,omitempty"`

	/*Pattern defined: A glob pattern, like `*.{ts,js}`. */
>>>>>>> v0.0.4
	Pattern string `json:"pattern,omitempty"`
}

// DocumentSelector is a type
/**
 * A document selector is the combination of one or many document filters.
 *
 * @sample `let sel:DocumentSelector = [{ language: 'typescript' }, { language: 'json', pattern: '**∕tsconfig.json' }]`;
 */
type DocumentSelector []DocumentFilter

// DefinitionLink is a type
/**
 * Information about where a symbol is defined.
 *
 * Provides additional metadata over normal [location](#Location) definitions, including the range of
 * the defining symbol
 */
type DefinitionLink LocationLink

// DeclarationLink is a type
/**
 * Information about where a symbol is declared.
 *
 * Provides additional metadata over normal [location](#Location) declarations, including the range of
 * the declaring symbol.
 *
 * Servers should prefer returning `DeclarationLink` over `Declaration` if supported
 * by the client.
 */
type DeclarationLink LocationLink
