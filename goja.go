package dartfmt

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"golang.org/x/sync/singleflight"
)

//go:embed internal/entrypoint/dist/entrypoint.js
var formatScript string
var formatProgram *goja.Program

const formatScriptFormat = `
var self = {location: {href: '/'}};
%s
var formatDart = self.formatDart;
`

func compileFormatterForSingleflight() (interface{}, error) {
	if formatProgram != nil {
		return formatProgram, nil
	}

	prog, err := goja.Parse(
		"entrypoint.js",
		fmt.Sprintf(formatScriptFormat, formatScript),
		parser.WithDisableSourceMaps,
	)

	if err != nil {
		return nil, fmt.Errorf("parsing a script failed: %w", err)
	}

	formatProgram, err = goja.CompileAST(prog, false)

	if err != nil {
		return nil, fmt.Errorf("compiling AST failed: %w", err)
	}

	return formatProgram, nil
}

var group singleflight.Group
var pool = sync.Pool{
	New: func() interface{} {
		fg, err, _ := group.Do("entrypoint.js", compileFormatterForSingleflight)

		if err != nil {
			panic(err)
		}

		vm := goja.New()

		_, err = vm.RunProgram(fg.(*goja.Program))

		if err != nil {
			panic(err)
		}

		return vm
	},
}

// FormatDartStandalone formats a script without dependencies on external commands
func FormatDartStandalone(script string) (string, error) {
	vm := pool.Get().(*goja.Runtime)
	defer pool.Put(vm)

	dartFormat := vm.GlobalObject().Get("formatDart")

	if dartFormat == nil {
		return "", fmt.Errorf("formatDart is nil")
	}

	var fn func(string) string
	err := vm.ExportTo(dartFormat, &fn)
	if err != nil {
		return "", fmt.Errorf("failed to execute dart_style formatter: %w", err)
	}

	return fn(script), nil
}
