package main

import (
	"errors"
	"fmt"
	"leetcode/cmd/create_file/lc"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func main() {
	ops, args := getOps()
	if ops.help {
		usage()
		return
	}

	app := &application{ops: ops}
	if err := app.run(args); err != nil {
		log.Fatalf("err=%v", err)
	}
}

func usage() {
	fmt.Println("Usage: <cmd> <leetcode_problem_url>")
}

func getOps() (*option, []string) {
	ops := &option{}
	return ops, ops.bindFlags()
}

type option struct {
	help     bool
	openWith string
}

func (o *option) bindFlags() []string {
	fs := pflag.NewFlagSet("cli", pflag.ContinueOnError)
	fs.BoolVarP(&o.help, "help", "h", o.help, "show help")
	fs.StringVarP(&o.openWith, "open", "o", o.openWith, "open created file with given application, default is 'goland'")
	_ = fs.Parse(os.Args)
	return fs.Args()[1:]
}

type application struct {
	ops *option
}

func (a *application) run(args []string) error {
	if len(args) == 0 {
		return errors.New("not enough argument, need an URL")
	}

	file, err := lc.New().Prepare(args[0], lc.Go)
	if err != nil && err != lc.ErrExist {
		return err
	}

	needOpen := false
	for _, s := range os.Args {
		if s == "-o" || s == "--open" {
			needOpen = true
			break
		}
	}

	if !needOpen {
		return nil
	}
	openCmd := a.ops.openWith
	if openCmd == "" {
		openCmd = "goland"
	}
	return exec.Command(openCmd, file).Run()
}
