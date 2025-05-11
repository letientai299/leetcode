package main

import (
  "errors"
  "fmt"
  "log"
  "os"
  "os/exec"
  "path"
  "syscall"

  "leetcode/create_file/lc"

  "github.com/spf13/pflag"
)

func main() {
  ops := getOps()
  if ops.help || len(ops.url) == 0 {
    usage(ops)
    return
  }

  app := &application{ops: ops}
  if err := app.run(); err != nil {
    log.Fatalf("err=%v", err)
  }
}

func usage(o *option) {
  fmt.Println("Usage: <cmd> <leetcode_problem_url> [-l lang] [-o]")
  fmt.Println(o.fs.FlagUsagesWrapped(80))
}

func getOps() *option {
  ops := &option{
    lang: "go",
  }

  remainArgs := ops.bindFlags()
  if len(remainArgs) > 0 {
    ops.url = remainArgs[0]
  }

  for _, s := range remainArgs {
    if s == "-o" || s == "--open" {
      ops.open = true
      break
    }
  }

  return ops
}

type option struct {
  help     bool
  url      string
  open     bool
  openWith string
  fs       *pflag.FlagSet
  lang     string
  force    bool
  read     bool
}

func (o *option) bindFlags() []string {
  fs := pflag.NewFlagSet("cli", pflag.ContinueOnError)
  fs.BoolVarP(&o.help, "help", "h", o.help, "show help")
  fs.StringVarP(&o.lang, "lang", "l", o.lang, "language to generate")
  fs.StringVarP(&o.openWith, "open", "o", o.openWith,
    "open created file with given command, if empty, use 'vim'")
  fs.BoolVarP(&o.force, "force", "f", o.force, "force create file, even if it exists")
  fs.BoolVarP(&o.read, "read", "r", o.read,
    "read the problem and its supported languages from the URL, "+
      "don't create file")

  _ = fs.Parse(os.Args)
  o.fs = fs
  return fs.Args()[1:]
}

type application struct {
  ops *option
}

func (a *application) run() error {
  url := a.ops.url
  if a.ops.read {
    return lc.New().Read(url)
  }

  lang, err := lc.ValidateLang(a.ops.lang)
  if err != nil {
    return err
  }

  file, err := lc.New().Prepare(url, lang, a.ops.force)
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
    editor, ok := os.LookupEnv("EDITOR")
    if !ok {
      editor = "nvim"
    }
    openCmd = editor
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
