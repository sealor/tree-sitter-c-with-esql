package tree_sitter_c_with_esql_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c_with_esql "github.com/tree-sitter/tree-sitter-c_with_esql/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_c_with_esql.Language())
	if language == nil {
		t.Errorf("Error loading C grammar")
	}
}
