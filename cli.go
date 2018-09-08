package main

import (
	"os"
	"time"

	og "github.com/champii/og/lib"

	"github.com/urfave/cli"
)

func parseArgs(done func(og.OgConfig)) {
	cli_ := setupCli()

	cli_.Action = func(c *cli.Context) error {
		options := og.OgConfig{
			Blocks:  c.Bool("b"),
			Dirty:   c.Bool("d"),
			Print:   c.Bool("p"),
			Ast:     c.Bool("a"),
			Verbose: c.Bool("v"),
			OutPath: c.String("o"),
			Paths:   []string(c.Args()),
		}

		done(options)

		return nil
	}

	cli_.Run(os.Args)
}

func setupCli() *cli.App {
	cli.AppHelpTemplate = `NAME:
	{{.Name}} - {{.Usage}}

USAGE:
	{{if .VisibleFlags}}{{.HelpName}} [options] [folders...|files...]

	By default it compiles the given files.
	If a Print argument (-p, -b, -d, -a) is given, NO COMPILATION is done.

	If run without any arguments, a small interpreter is spawn (ALPHA){{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
VERSION:
	{{.Version}}

OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Copyright }}

COPYRIGHT:
	{{.Copyright}}
	{{end}}{{if .Version}}
	{{end}}`

	cli.VersionFlag = cli.BoolFlag{
		Name:  "V, version",
		Usage: "Print version",
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "h, help",
		Usage: "Print help",
	}

	app := cli.NewApp()

	app.Name = "Oglang"
	app.Version = "v0.4.2"
	app.Compiled = time.Now()

	app.Usage = "Golang on steroids"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "o, out",
			Usage: "Output `directory`. If input is recursive folder, the tree is recreated",
			Value: "./",
		},
		cli.BoolFlag{
			Name:  "p, print",
			Usage: "Print the file",
		},
		cli.BoolFlag{
			Name:  "d, dirty",
			Usage: "Print the file before going through 'go fmt'",
		},
		cli.BoolFlag{
			Name:  "b, blocks",
			Usage: "Print the file after it goes to preprocessor. Shows only block-based indentation",
		},
		cli.BoolFlag{
			Name:  "a, ast",
			Usage: "Print the generated AST",
		},
		cli.BoolFlag{
			Name:  "v, verbose",
			Usage: "Show the filenames",
		},
	}

	app.UsageText = "og [options] folders|files"

	return app
}
