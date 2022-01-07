package dartfmt

import (
	_ "embed"
	"errors"
	"fmt"
)

// Formatter is a type of dart formatters
type Formatter func(script string) (string, error)

// Formatters are candidates for FormatDart
var Formatters = []Formatter{
	FormatDartWithFlutterCLI,
	FormatDartWithDartCLI,
	FormatDartWithDartfmtCLI,
	FormatDartStandalone,
}

// FormatDart tries to format a script with all of Formatters one-by-one
func FormatDart(script string) (string, error) {
	for _, fmtr := range Formatters {
		formatted, err := fmtr(script)

		if err != nil {
			if errors.As(err, &unavailableError{}) {
				continue
			}

			return "", fmt.Errorf("formatting failed: %w", err)
		}

		return formatted, nil
	}

	return "", fmt.Errorf("no available formatter")
}
