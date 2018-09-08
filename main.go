package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	og "github.com/champii/og/lib"
)

// // //go:generate java -Xmx1000M -cp "/usr/local/lib/antlr-4.7.1-complete.jar" org.antlr.v4.Tool -Dlanguage=Go ./parser/Og.g4 -visitor -o .
//go:generate java -Xmx500M -cp "./parser/antlr4-4.7.2-SNAPSHOT-complete.jar" org.antlr.v4.Tool -Dlanguage=Go ./parser/Og.g4 -visitor -o .
//go:generate ./scripts/fix_parser_imports.sh
func main() {
	f, err := os.Create("/tmp/profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	parseArgs(func(options og.OgConfig) {
		if len(options.Paths) == 0 {
			og.RunInterpreter()
		}
		if err := og.Compile(options); err != nil {
			fmt.Println(err)
		}
	})

}
