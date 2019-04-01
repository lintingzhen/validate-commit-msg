package main

import (
    "os"
    "fmt"
    "flag"
)

type Arguments struct {
    install bool
    debug bool
    args []string
}

var revision string

var usage string = 
`Validate-commit-msg(golang) v0.1:

Usage: validate-commit-msg [-i] [-d] [-v] [<args>]

Options:
    -i, -install    Install this tool to .git/hooks/ as commit-msg
    -d, -debug      Debug mode
    -v, -version    Print version information and quit
    args            Commit message temp file path, use by git-commit
`

func ParseArgs(args *Arguments) {
    var (
        help bool
        version bool
    )
    flag.BoolVar(&args.install, "i", false, "Install this tool to git-core as git-cz.")
    flag.BoolVar(&args.install, "install", false, "Install this tool to git-core as git-cz.")
    flag.BoolVar(&args.debug, "d", false, "Debug mode.")
    flag.BoolVar(&args.debug, "debug", false, "Debug mode.")
    flag.BoolVar(&help, "h", false, "Show the help.")
    flag.BoolVar(&help, "help", false, "Show the help.")
    flag.BoolVar(&version, "v", false, "")
    flag.BoolVar(&version, "version", false, "")

    flag.Usage = func() {
        fmt.Println(usage)
    }

    flag.Parse()
    if help {
        flag.Usage()
        os.Exit(0)
    } else if version {
        fmt.Printf("Validate-commit-msg(golang) version 0.1, build revision %s\n", revision)
        os.Exit(0)
    }

    args.args = flag.Args()
}

