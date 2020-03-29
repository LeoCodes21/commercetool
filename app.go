package main

import (
	"fmt"
	"github.com/LeoCodes21/commercetool/commands"
	"github.com/alecthomas/kong"
)

var cli struct {
	Debug        bool                                 `help:"Enable debug mode."`
	CarScript    commands.GenerateSpecificScriptCmd   `cmd:"" help:"Generate a ModScript for customization specific to a car."`
	GlobalScript commands.GenerateGlobalScriptCommand `cmd:"" help:"Generate a ModScript to make \"universal\" parts compatible with a car."`
}

func main() {
	fmt.Println("CommerceTool v1.0.0 by heyitsleo")
	ctx := kong.Parse(&cli)
	err := ctx.Run(&commands.Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
