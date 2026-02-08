package cli

import (
	"fmt"

	"github.com/colfarl/budgy/internal/store"
)

type Context struct {
	Debug bool
	Store store.Store
}

type RmCmd struct {
  Force     bool `help:"Force removal."`
  Recursive bool `help:"Recursively remove files."`

  Paths []string `arg:"" name:"path" help:"Paths to remove." type:"path"`
}

func (r *RmCmd) Run(ctx *Context) error {
  fmt.Println("rm", r.Paths)
  return nil
}

type LsCmd struct {
  Paths []string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
}

func (l *LsCmd) Run(ctx *Context) error {
  fmt.Println("ls", l.Paths)
  return nil
}

var CliCommands struct {
  Debug bool `help:"Enable debug mode."`

  Rm RmCmd `cmd:"" help:"Remove files."`
  Ls LsCmd `cmd:"" help:"List paths."`
}
