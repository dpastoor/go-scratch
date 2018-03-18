package main

import (
	"os"
	"os/exec"
)

// DockerExec shows execing into a docker container shell
func DockerExec() {
	cmd := exec.Command("docker",
		"run",
		"--rm",
		"-it",
		"osmosisfoundation/rocker",
		"/bin/sh",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
