//go:generate goversioninfo -icon=favicon.ico -manifest=goversioninfo.exe.manifest

package main

import (
	"time"

	"github.com/alexflint/go-arg"
	"github.com/gookit/color"
)

type UserInput struct {
	Level  int    `arg:"-l,required"`
	Word   string `arg:"--wordlist,-w"`
	Output string `arg:"-o,required"`
	D      bool   `arg:"-d"`
	Input  string `arg:"-i"`
}

func main() {
	userInput := UserInput{}
	arg.MustParse(&userInput)

	color.Yellowln("Processing...")
	start := time.Now()

	mungeInit(userInput)

	duration := time.Since(start)
	color.Cyanln("Finished in:", duration)
}
