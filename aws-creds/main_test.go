package main

import (
	"testing"
)

var credItemsOutputTcs = []struct {
	items    []string
	opts     *options
	expected string
}{
	{[]string{"FOO=BAR", "BAZ=BLAH"}, &options{Export: false}, "FOO=BAR BAZ=BLAH"},
	{[]string{"FOO=BAR", "BAZ=BLAH"}, &options{Export: true}, "export FOO=BAR\nexport BAZ=BLAH\n"},
}

func TestCredItemsOutput(t *testing.T) {
	for _, tc := range credItemsOutputTcs {
		result := formatCredItems(tc.items, tc.opts)
		if tc.expected != result {
			t.Fatalf("Expected '%v' got '%v' for export %t", tc.expected, result, tc.opts.Export)
		}

	}
}
