package dartfmt

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func runCLI(script string, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = strings.NewReader(script)

	stdout, stderr := bytes.Buffer{}, bytes.Buffer{}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return "", unavailableError{err: err}
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf(
			"formatter failed(%s): %w",
			stdout.String()+"\n"+stderr.String(),
			err,
		)
	}

	return stdout.String(), nil
}

// FormatDartWithDartCLI formats a script with `dart format`
func FormatDartWithDartCLI(script string) (string, error) {
	formatted, err := runCLI(script, "dart", "format")

	if err != nil {
		return "", fmt.Errorf("formatting with `dart format` failed: %w", err)
	}

	return formatted, nil
}

// FormatDartWithDartfmtCLI formats a script with `dartfmt`
func FormatDartWithDartfmtCLI(script string) (string, error) {
	formatted, err := runCLI(script, "dartfmt")

	if err != nil {
		return "", fmt.Errorf("formatting with `dartfmt` failed: %w", err)
	}

	return formatted, nil
}
