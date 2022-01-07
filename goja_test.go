package dartfmt_test

import (
	"testing"

	"github.com/go-generalize/go-dartfmt"
)

func TestFormatDartStandalone(t *testing.T) {
	runFormat(t, dartfmt.FormatDartStandalone, "testdata/001")
}
