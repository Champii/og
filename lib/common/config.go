package common

type OgConfig struct {
	Blocks      bool
	Dirty       bool
	Print       bool
	Force       bool
	Ast         bool
	SimpleAst   bool
	Quiet       bool
	Interpreter bool
	Paths       []string
	Workers     int
	OutPath     string
	NoBuild     bool
	Run         bool
}

func NewOgConfig() *OgConfig {
	return &OgConfig{
		Workers:     8,
		Interpreter: false,
		OutPath:     "./",
	}
}
