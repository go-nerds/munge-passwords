//go:generate goversioninfo -icon=favicon.ico -manifest=goversioninfo.exe.manifest

package main

import (
	"github.com/alexflint/go-arg"
	"github.com/gookit/color"
)

var args struct {
	Level  int    `arg:"-l,required"`
	Word   string `arg:"--wordlist,-w"`
	Output string `arg:"-o,required"`
	D      bool   `arg:"-d"`
	Input  string `arg:"-i"`
}

func main() {
	arg.MustParse(&args)
	color.Yellowln("Processing...")
	mungeInit(args.Level, args.Word, args.Output, args.D, args.Input)
}
