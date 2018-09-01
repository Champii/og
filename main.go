package main

import (
	og "github.com/champii/og/lib"
)

//go:generate java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar" org.antlr.v4.Tool -Dlanguage=Go ./parser/Og.g4 -visitor -o .
//go:generate ./scripts/fix_parser_imports.sh
func main() {
	parseArgs(func(options og.OgConfig) {
		og.Compile(options)
	})

}
