package main

import (
	"fmt"

	"github.com/champii/og/lib/common"
	"github.com/champii/og/lib/og"
)

//go:generate java -Xmx500M -cp "./parser/antlr4-4.7.2-SNAPSHOT-complete.jar" org.antlr.v4.Tool -Dlanguage=Go ./parser/Og.g4 -visitor -o .
//go:generate ./scripts/fix_parser_imports.sh
func main() {
	parseArgs(func(options *common.OgConfig) {
		ogLang := og.NewOg(options)

		if err := ogLang.Run(); err != nil {
			fmt.Println(err)
		}
	})
}
