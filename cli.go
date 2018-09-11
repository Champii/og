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
			Blocks:      c.Bool("b"),
			Dirty:       c.Bool("d"),
			Print:       c.Bool("p"),
			Ast:         c.Bool("a"),
			Quiet:       c.Bool("q"),
			Workers:     c.Int("w"),
			OutPath:     c.String("o"),
			Interpreter: c.Bool("i"),
			NoBuild:     c.Bool("n"),
			Run:         c.Bool("r"),
			Paths:       []string(c.Args()),
		}

		done(options)

		return nil
	}

	cli_.Run(os.Args)
}

func setupCli() *cli.App {
	cli.AppHelpTemplate = `NAME:
	{{.Name}} - {{.Usage}}

VERSION:
	{{.Version}}

USAGE:
	{{if .VisibleFlags}}{{.HelpName}} [options] [folders...|files...]

	If a Print argument (-p, -b, -d, -a) is given, NO COMPILATION is done.

	If run without files, it will compile and build '.'{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Copyright }}

COPYRIGHT:
	{{.Copyright}}
	{{end}}{{if .Version}}
	{{end}}
`

	cli.VersionFlag = cli.BoolFlag{
		Name:  "v, version",
		Usage: "Print version",
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "h, help",
		Usage: "Print help",
	}

	app := cli.NewApp()

	app.Name = "Oglang"
	app.Version = "v0.6.2"
	app.Compiled = time.Now()

	app.Usage = "Golang on steroids"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "r, run",
			Usage: "Run the binary",
		},
		cli.StringFlag{
			Name:  "o, out",
			Usage: "Output `directory`. If input is recursive folder, the tree is recreated",
			Value: "./",
		},
		cli.IntFlag{
			Name:  "w, workers",
			Usage: "Set the number of `jobs`",
			Value: 8,
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
			Name:  "i, interpreter",
			Usage: "Run a small interpreter (ALPHA)",
		},
		cli.BoolFlag{
			Name:  "q, quiet",
			Usage: "Hide the progress output",
		},
		cli.BoolFlag{
			Name:  "n, no-build",
			Usage: "Don't run 'go build'",
		},
	}

	app.UsageText = "og [options] [folders|files]"

	return app
}
