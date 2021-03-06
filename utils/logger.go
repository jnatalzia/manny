package utils

import (
	"log"
	"os"
	"regexp"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

var (
	blueFg    = color.New(color.FgBlue)
	magentaFg = color.New(color.FgMagenta)
)

var Logger = log.New(os.Stdout, "", log.Lshortfile)

func InitLogger() {
	outWrite := marker.NewStdoutMarker()
	rex, _ := regexp.Compile("\\w+\\.go")
	markRules := []marker.MarkRule{
		{marker.MatchBracketSurrounded(), magentaFg},
		{marker.MatchRegexp(rex), blueFg},
	}

	outWrite.AddRules(markRules)

	Logger.SetOutput(outWrite)
}
