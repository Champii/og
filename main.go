package main

import (
	"fmt"

	og "github.com/champii/og/lib"
)

//go:generate java -Xmx500M -cp "./parser/antlr4-4.7.1-SNAPSHOT-complete.jar" org.antlr.v4.Tool -Dlanguage=Go ./parser/Og.g4 -visitor -o .
//go:generate ./scripts/fix_parser_imports.sh
func main() {
	parseArgs(func(options og.OgConfig) {
		if err := og.Compile(options); err != nil {
			fmt.Println(err)
		}
	})

}
