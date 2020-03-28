package commands

import "fmt"

type GenerateSpecificScriptCmd struct {
	CarConfigPath string `arg:"" name:"config" help:"The path to the CarConfig file" type:"path"`
}

func (g *GenerateSpecificScriptCmd) Run(ctx *Context) error {
	fmt.Printf("Generate specific script from config file: %s\n", g.CarConfigPath)

	return nil
}
