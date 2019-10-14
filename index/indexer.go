// Original source: https://github.com/sourcegraph/lsif-go/blob/master/index/indexer.go
// Package index is used to generate an LSIF dump for a workspace.
package index

import (
	//"encoding/json"
	"fmt"
	"io"
	//"io/ioutil"
	"path/filepath"
	//"strconv"
	//"strings"

	//"github.com/hashicorp/hcl2/hcl"
	//"github.com/hashicorp/hcl2/hcldec"
	"github.com/hashicorp/terraform/configs"
	//"github.com/hashicorp/terraform/lang"
	"github.com/juliosueiras/lsif-terraform/log"
	"github.com/juliosueiras/lsif-terraform/protocol"
	//"golang.org/x/tools/go/packages"
)

// Index generates an LSIF dump for a workspace by traversing through source files
// and storing LSP responses to output source that implements io.Writer. It is
// caller's responsibility to close the output source if applicable.
func Index(workspace string, excludeContent bool, w io.Writer, toolInfo protocol.ToolInfo) (*Stats, error) {
	projectRoot, err := filepath.Abs(workspace)
	if err != nil {
		return nil, fmt.Errorf("get abspath of project root: %v", err)
	}

	parser := configs.NewParser(nil)

	log.Infoln(parser.LoadConfigDir(projectRoot))

	e := &indexer{
		projectRoot:    projectRoot,
		excludeContent: excludeContent,
		w:              w,

		//pkgs:    pkgs,
		//files:   make(map[string]*fileInfo),
		//imports: make(map[token.Pos]*defInfo),
		//funcs:   make(map[string]*defInfo),
		//consts:  make(map[token.Pos]*defInfo),
		//vars:    make(map[token.Pos]*defInfo),
		//types:   make(map[string]*defInfo),
		//labels:  make(map[token.Pos]*defInfo),
		//refs:    make(map[string]*refResultInfo),
	}
	return e.index(toolInfo)
}

// indexer keeps track of all information needed to generate a LSIF dump.
type indexer struct {
	projectRoot    string
	excludeContent bool
	w              io.Writer

	id int // The ID counter of the last element emitted
	//pkgs    []*packages.Package
	//files   map[string]*fileInfo      // Keys: filename
	//imports map[token.Pos]*defInfo    // Keys: definition position
	//funcs   map[string]*defInfo       // Keys: full name (with receiver for methods)
	//consts  map[token.Pos]*defInfo    // Keys: definition position
	//vars    map[token.Pos]*defInfo    // Keys: definition position
	//types   map[string]*defInfo       // Keys: type name
	//labels  map[token.Pos]*defInfo    // Keys: definition position
	//refs    map[string]*refResultInfo // Keys: definition range ID
}

// Stats contains statistics of data processed during index.
type Stats struct {
	NumPkgs     int
	NumFiles    int
	NumDefs     int
	NumElements int
}

func (e *indexer) index(info protocol.ToolInfo) (*Stats, error) {
	//	_, err := e.emitMetaData("file://"+e.projectRoot, info)
	//	if err != nil {
	//		return nil, fmt.Errorf(`emit "metadata": %v`, err)
	//	}
	//	proID, err := e.emitProject()
	//	if err != nil {
	//		return nil, fmt.Errorf(`emit "project": %v`, err)
	//
	//	}
	//
	//	_, err = e.emitBeginEvent("project", proID)
	//	if err != nil {
	//		return nil, fmt.Errorf(`emit "begin": %v`, err)
	//	}
	//
	//	for _, p := range e.pkgs {
	//		if err := e.indexPkg(p, proID); err != nil {
	//			return nil, fmt.Errorf("index package %q: %v", p.Name, err)
	//		}
	//	}
	//
	//	for _, f := range e.files {
	//		for _, rangeID := range f.defRangeIDs {
	//			refResultID, err := e.emitReferenceResult()
	//			if err != nil {
	//				return nil, fmt.Errorf(`emit "referenceResult": %v`, err)
	//			}
	//
	//			_, err = e.emitTextDocumentReferences(e.refs[rangeID].resultSetID, refResultID)
	//			if err != nil {
	//				return nil, fmt.Errorf(`emit "textDocument/references": %v`, err)
	//			}
	//
	//			for docID, rangeIDs := range e.refs[rangeID].defRangeIDs {
	//				_, err = e.emitItemOfDefinitions(refResultID, rangeIDs, docID)
	//				if err != nil {
	//					return nil, fmt.Errorf(`emit "item": %v`, err)
	//				}
	//			}
	//
	//			for docID, rangeIDs := range e.refs[rangeID].refRangeIDs {
	//				_, err = e.emitItemOfReferences(refResultID, rangeIDs, docID)
	//				if err != nil {
	//					return nil, fmt.Errorf(`emit "item": %v`, err)
	//				}
	//			}
	//		}
	//
	//		if len(f.defRangeIDs) > 0 || len(f.useRangeIDs) > 0 {
	//			_, err = e.emitContains(f.docID, append(f.defRangeIDs, f.useRangeIDs...))
	//			if err != nil {
	//				return nil, fmt.Errorf(`emit "contains": %v`, err)
	//			}
	//		}
	//	}
	//
	//	// Close all documents. This must be done as a last step as we need
	//	// to emit everything about a document before sending the end event.
	//
	//	// TODO(efritz) - see if we can rearrange the outputs so that
	//	// all of the output for a document is contained in one segment
	//	// that does not interfere with emission of other document
	//	// properties.
	//
	//	for _, info := range e.files {
	//		_, err = e.emitEndEvent("document", info.docID)
	//		if err != nil {
	//			return nil, fmt.Errorf(`emit "end": %v`, err)
	//		}
	//	}
	//
	//	_, err = e.emitEndEvent("project", proID)
	//	if err != nil {
	//		return nil, fmt.Errorf(`emit "end": %v`, err)
	//	}
	//
	//return &Stats{
	//	NumPkgs:     len(e.pkgs),
	//	NumFiles:    len(e.files),
	//	NumDefs:     len(e.imports) + len(e.funcs) + len(e.consts) + len(e.vars) + len(e.types) + len(e.labels),
	//	NumElements: e.id,
	//}, nil
	return &Stats{}, nil
}
