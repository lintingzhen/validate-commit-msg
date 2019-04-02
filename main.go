package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "regexp"
    "strings"
    "io/ioutil"
    "flag"
)

func main() {
    var args Arguments
    ParseArgs(&args)

    // open debug switch
    if args.debug {
        log.SetFlags(log.Lshortfile | log.LstdFlags)
    } else {
        log.SetFlags(0)
        log.SetOutput(ioutil.Discard)
    }

    if args.install {
        // exit if not git directory
        ExitIfNotGitDirectory()

        // copy to .git/hooks/commit-msg
        HookCommitMsg()

        fmt.Println("Install git commit-msg hook success.")
        return 
    }

    if len(flag.Args()) != 1 {
        fmt.Println("Only for commit msg.")
        os.Exit(1)
    }
    log.SetOutput(ioutil.Discard)
    commitMsgTempFilePath := flag.Args()[0]

    f, err := os.Open(commitMsgTempFilePath)
    if err != nil {
        fmt.Println("Open git commit msg file failed.")
        os.Exit(1)
    }

    defer f.Close()

    sc := bufio.NewScanner(f)
    if !sc.Scan() {
        fmt.Println("Empty commit msg.")
        os.Exit(1)
    }
    line := sc.Text()
    log.Printf("%s\n", line)
    reg := regexp.MustCompile(`^\w+(\([\w\s]*\))?:\s*([\w\s]*?)\s*$`) 
    result := reg.FindAllStringSubmatch(string(line), -1)
    if result == nil {
        fmt.Println("Invalid commit msg, valid msg format: \n<type>[(<scope>)]: <subject>\n\n[<body>]\n\n[<footer>]")
        os.Exit(1)
    }

    log.Printf("%q\n", result)
    scope_with_parent := result[0][1]
    subject := result[0][2]
    log.Printf("scope=%s|subject=%s\n", scope_with_parent, subject)

    if scope_with_parent != "" {
        scope := strings.TrimSpace(scope_with_parent[1:(len(scope_with_parent)) - 1])
        if scope == "" {
            fmt.Println("Empty scope in () is valid.")
            os.Exit(1)
        }
    }
    if subject == "" {
        fmt.Println("Empty subject is invalid.")
        os.Exit(1)
    }
}

