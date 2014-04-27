package main

import (
    "fmt"
    "flag"
    "os/exec"
    "bytes"
)

func run(command []string) {
    rest := command[1:len(command)]
    cmd := exec.Command(command[0], rest...)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out
    err := cmd.Run()
    if err != nil {
        fmt.Print(out.String())
    } else if !cmd.ProcessState.Success() {
        fmt.Print(out.String())
    }
}

func main() {
    flag.Parse()
    run(flag.Args())
}
