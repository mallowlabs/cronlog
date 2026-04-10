package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"
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

func run(command []string) int {
	if len(command) == 0 {
		if info, ok := debug.ReadBuildInfo(); ok {
			fmt.Println(info.Main.Version)
			return 0
		} else {
			fmt.Fprintln(os.Stderr, "Error: could not read build info")
			return 1
		}
	}
	rest := command[1:]
	cmd := exec.Command(command[0], rest...)

	config := ReadConfig("/etc/cronlog.toml")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()

	exitCode := cmd.ProcessState.ExitCode()

	if err != nil {
		if exitCode != config.FindCommand(command[0]).SuccessCode {
			output(command, out.String(), config)
		}
	} else if !cmd.ProcessState.Success() {
		output(command, out.String(), config)
	}
	return exitCode
}

func main() {
	flag.Parse()
	os.Exit(run(flag.Args()))
}
