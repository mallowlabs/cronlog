package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func output(command []string, text string) {
	fmt.Print(text)

	config := ReadConfig("/etc/cronlog.toml")
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
		return
	}
	rest := command[1:len(command)]
	cmd := exec.Command(command[0], rest...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		output(command, out.String())
	} else if !cmd.ProcessState.Success() {
		output(command, out.String())
	}
}

func main() {
	flag.Parse()
	run(flag.Args())
}

