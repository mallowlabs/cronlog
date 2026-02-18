package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"
	"syscall"
)

func output(command []string, text string, config Config) {
	fmt.Print(text)

	if config.Slack.Url != "" {
		message := "```\n"
		message += text
		message += "```\n"

		hostname, _ := os.Hostname()
		attributes := map[string]string{"Host": hostname, "Command": strings.Join(command, " ")}

		PostToSlack(message, attributes, config.Slack)
	}
}

func run(command []string) {
	if len(command) == 0 {
		if info, ok := debug.ReadBuildInfo(); ok {
			fmt.Println(info.Main.Version)
		} else {
			fmt.Fprintln(os.Stderr, "Error: could not read build info")
			os.Exit(1)
		}
		return
	}
	rest := command[1:len(command)]
	cmd := exec.Command(command[0], rest...)

	config := ReadConfig("/etc/cronlog.toml")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()

	if err != nil {
		exitCode := 0
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		}

		if exitCode != config.FindCommand(command[0]).SuccessCode {
			output(command, out.String(), config)
		}
	} else if !cmd.ProcessState.Success() {
		output(command, out.String(), config)
	}
}

func main() {
	flag.Parse()
	run(flag.Args())
}
