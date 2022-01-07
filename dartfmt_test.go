package dartfmt_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-generalize/go-dartfmt"
	"github.com/google/go-cmp/cmp"
)

func runFormat(t *testing.T, formatter dartfmt.Formatter, dir string) {
	t.Helper()

	input, err := os.ReadFile(filepath.Join(dir, "input.dart"))

	if err != nil {
		t.Fatal(err)
	}

	expectedBytes, err := os.ReadFile(filepath.Join(dir, "output.dart"))

	if err != nil {
		t.Fatal(err)
	}
	expected := string(expectedBytes)

	actual, err := dartfmt.FormatDart(string(input))

	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("the result differed: %s", diff)
	}
}

func TestDartFormat(t *testing.T) {
	runFormat(t, dartfmt.FormatDart, "testdata/001")
}
