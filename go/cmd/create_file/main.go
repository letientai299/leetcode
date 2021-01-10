package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/pflag"
)

var (
	errExist = errors.New("file existed")
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
	fmt.Println("Try <cmd> '1386. Cinema Seat Allocation'")
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
	fs.Parse(os.Args)
	return fs.Args()[1:]
}

type application struct {
	ops *option
}

func (a *application) run(args []string) error {
	file, err := a.create(args)
	if err != nil && err != errExist {
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

func (a *application) create(ss []string) (string, error) {
	var sb strings.Builder
	sb.WriteString(ss[0])

	for _, s := range ss[1:] {
		s = strings.ToLower(s)
		sb.WriteString(s)
		sb.WriteString("-")
	}

	s := sb.String()
	s = s[:len(s)-1]
	s += ".go"

	if a.fileExist(s) {
		return s, errExist
	}
	f, err := os.Create(s)
	if err != nil {
		log.Printf("fail to craete file, err=%v", err)
		return "", err
	}

	fmt.Fprintln(f, "package main")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "// medium")
	fmt.Fprintf(f, "// %s\n", strings.Join(os.Args[1:], " "))
	_ = f.Close()
	fmt.Println(s)
	return s, nil
}

func (a *application) fileExist(s string) bool {
	_, err := os.Open(s)
	return !os.IsNotExist(err)
}
