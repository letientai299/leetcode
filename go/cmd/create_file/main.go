package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func main() {
	ops := getOps()
	if ops.help {
		usage()
		return
	}

	app := &application{ops: ops}
	if err := app.run(); err != nil {
		log.Fatalf("err=%v", err)
	}
}

func usage() {
	fmt.Println("Try <cmd> '1386. Cinema Seat Allocation'")
}

func getOps() *option {
	ops := &option{}
	ops.bindFlags()
	pflag.Parse()
	return ops
}

type option struct {
	help bool
}

func (o *option) bindFlags() {
	pflag.BoolVarP(&o.help, "help", "h", o.help, "show help")
}

type application struct {
	ops *option
}

func (a *application) run() error {
	if len(os.Args) == 1 {
		return errors.New("need file name")
	}

	ss := os.Args[1:]
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
		return errors.New("file existed")
	}
	f, err := os.Create(s)
	if err != nil {
		log.Printf("fail to craete file, err=%v", err)
		return err
	}

	fmt.Fprintln(f, "package main")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "// medium")
	fmt.Fprintf(f, "// %s\n", strings.Join(os.Args[1:], " "))
	_ = f.Close()
	fmt.Println(s)
	return nil
}

func (a *application) fileExist(s string) bool {
	_, err := os.Open(s)
	return !os.IsNotExist(err)
}
