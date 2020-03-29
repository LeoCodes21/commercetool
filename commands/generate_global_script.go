package commands

import (
	"bufio"
	"fmt"
	"github.com/LeoCodes21/commercetool/lib"
	"os"
	"strings"
)

type GenerateGlobalScriptCommand struct {
	CarName          string `arg:"" name:"car-name" help:"The name of the car to generate a script for"`
	ScriptOutputPath string `arg:"" name:"out-path" help:"The path (including file name) where the script should be created" type:"path"`
}

func (g *GenerateGlobalScriptCommand) Run(_ *Context) error {
	f, err := os.Create(g.ScriptOutputPath)

	fmt.Printf("Generating script for car %s to %s\n", g.CarName, g.ScriptOutputPath)

	if err != nil {
		return err
	}

	defer f.Close()

	listFile, err := os.Open("global_visualpart_list.txt")
	if err != nil {
		return err
	}

	defer listFile.Close()

	scanner := bufio.NewScanner(listFile)
	scriptLines := []string{"# Generated by CommerceTool by heyitsleo"}

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "0x") {
			return fmt.Errorf("invalid line in part list: %s", line)
		}
		scriptLines = append(scriptLines, fmt.Sprintf("append_array visualpart %s baseCarHashes %d", line, lib.BinHash(g.CarName)))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	for _, l := range scriptLines {
		_, err = fmt.Fprintln(f, l)
		if err != nil {
			return err
		}
	}

	return nil
}
