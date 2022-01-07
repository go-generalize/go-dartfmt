package dartfmt_test

import (
	"testing"

	"github.com/go-generalize/go-dartfmt"
)

func TestFormatDartWithDartCLI(t *testing.T) {
	runFormat(t, dartfmt.FormatDartWithDartCLI, "testdata/001")
}

func TestFormatDartWithDartfmtCLI(t *testing.T) {
	runFormat(t, dartfmt.FormatDartWithDartfmtCLI, "testdata/001")
}

func TestFormatDartWithFlutterCLI(t *testing.T) {
	runFormat(t, dartfmt.FormatDartWithFlutterCLI, "testdata/001")
}
