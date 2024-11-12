package cli

import (
	"fmt"

	"github.com/fatih/color"
)

type CLI struct {
	version string
}

func New() (*CLI, error) {
	return &CLI{
		version: "1.0.0",
	}, nil
}

func (c *CLI) Run() {
	color.Cyan(`
╔══════════════════Gilfoyle CLI ═══════════════╗
║              Because deploying is not        ║
║              an unacceptable strategy        ║
╚══════════════════════════════════════════════╝
	`)
	fmt.Println("Version:", color.GreenString(c.version))
}

func (c *CLI) Cleanup() {
	fmt.Println("\nNettoyage avant de partir (comme si vous vous en souciez)...")
}