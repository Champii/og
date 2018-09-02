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
	{{if .VisibleFlags}}{{.HelpName}} [options] Folders|Files{{end}}
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
	app.Version = "v0.1.5"
	app.Compiled = time.Now()

	app.Usage = "Golang on steroids"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "o, out",
			Usage: "Output directory. If input is recursive folder, the tree is recreated",
			Value: "./",
		},
		cli.BoolFlag{
			Name:  "p, print",
			Usage: "Print only to stdout. No files created",
		},
		cli.BoolFlag{
			Name:  "d, dirty",
			Usage: "Don't use 'go fmt'",
		},
		cli.BoolFlag{
			Name:  "b, blocks",
			Usage: "Get only the generated blocks from indent. No compilation to go.",
		},
		cli.BoolFlag{
			Name:  "v, verbose",
			Usage: "Show the filenames",
		},
	}

	app.UsageText = "og [options] folders|files"

	return app
}
