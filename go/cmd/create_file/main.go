package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"syscall"

	"leetcode/cmd/create_file/lc"

	"github.com/spf13/pflag"
)

func main() {
	ops, args := getOps()
	if ops.help || len(args) == 0 {
		usage(ops)
		return
	}

	app := &application{ops: ops}
	if err := app.run(args); err != nil {
		log.Fatalf("err=%v", err)
	}
}

func usage(o *option) {
	fmt.Println("Usage: <cmd> <leetcode_problem_url> [-l lang] [-o]")
	fmt.Println(o.fs.FlagUsagesWrapped(80))
}

func getOps() (*option, []string) {
	ops := &option{
		lang: "go",
	}
	return ops, ops.bindFlags()
}

type option struct {
	help     bool
	openWith string
	fs       *pflag.FlagSet
	lang     string
	force    bool
}

func (o *option) bindFlags() []string {
	fs := pflag.NewFlagSet("cli", pflag.ContinueOnError)
	fs.BoolVarP(&o.help, "help", "h", o.help, "show help")
	fs.StringVarP(&o.lang, "lang", "l", o.lang, "language to generate")
	fs.StringVarP(&o.openWith, "open", "o", o.openWith,
		"open created file with given command, if empty, use 'vim'")
	fs.BoolVarP(&o.force, "force", "f", o.force, "force create file, even if it exists")
	_ = fs.Parse(os.Args)
	o.fs = fs
	return fs.Args()[1:]
}

type application struct {
	ops *option
}

func (a *application) run(args []string) error {
	if len(args) == 0 {
		return errors.New("not enough argument, need an URL")
	}

	lang, err := lc.ValidateLang(a.ops.lang)
	if err != nil {
		return err
	}

	file, err := lc.New().Prepare(args[0], lang, a.ops.force)
	if err != nil && !errors.Is(err, lc.ErrExist) {
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

	return a.open(file)
}

func (a *application) open(file string) error {
	openCmd := a.ops.openWith
	if openCmd == "" {
		openCmd = "vim"
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("fail to get working dir, err=%v", err)
		return err
	}
	file = path.Join(wd, file)

	cmd := exec.Command("sh", "-c", openCmd+" "+file)
	log.Println("starting", cmd)

	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cred := syscall.Credential{
		Uid:         uint32(os.Getuid()),
		Gid:         uint32(os.Getgid()),
		NoSetGroups: true,
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{Credential: &cred}

	return cmd.Run()
}
