package hub

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestCleanRelativeFilePath(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"foo/bar", "foo/bar"},
		{"foo/../bar", "bar"},
		{"foo/./bar", "foo/bar"},
		{"/foo/bar", "foo/bar"},
		{"foo//bar", "foo/bar"},
		{"foo/bar/..", "foo"},
		{"../foo/bar", "foo/bar"},
		{"foo/../../../..", "."},
		{"foo/../../../bar", "bar"},
		{"", "."},
		{".", "."},
		{"..", "."},
	}

	for _, tc := range testCases {
		expected := filepath.FromSlash(tc.expected)
		got := cleanRelativeFilePath(tc.input)
		fmt.Printf("\tcleanRelativeFilePath(%q) = %q\n", tc.input, got)
		assert.Equal(t, expected, got)
	}
}
