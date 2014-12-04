package main

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"syscall"
	"testing"

	"github.com/thbishop/aws-creds/Godeps/_workspace/src/github.com/vaughan0/go-ini"
)

func TestCredItems(t *testing.T) {
	content := `
[default]
aws_access_key_id=foo
aws_secret_access_key=bar
aws_session_token=baz
`
	expected := []string{
		"AWS_ACCESS_KEY_ID=foo",
		"AWS_SECRET_ACCESS_KEY=bar",
		"AWS_SESSION_TOKEN=baz",
	}

	f, err := ioutil.TempFile("", "aws-creds-test")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer syscall.Unlink(f.Name())
	ioutil.WriteFile(f.Name(), []byte(content), 0644)

	path, err := filepath.Abs(f.Name())
	if err != nil {
		t.Fatalf("Error getting temp file path: %s", err)
	}

	ini, err := ini.LoadFile(path)
	if err != nil {
		t.Fatalf("Unable to load ini file: %s", err)
	}

	result := credItems("default", ini)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Excepted '%v' got '%v'", expected, result)
	}
}

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
