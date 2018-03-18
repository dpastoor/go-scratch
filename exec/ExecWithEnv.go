package main

import (
	"fmt"
	"os"
	"os/exec"
)

// ExecWithEnv execs while attempting to set an env variable
func ExecWithEnv() {
	cmds := []string{
		"-e -q --vanilla",
		"packageVersion(\"rlang\")",
	}
	cmd := exec.Command("R", cmds...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	fmt.Println("running cmd")
	cmd.Run()
	cmd2 := exec.Command("R", cmds...)
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	cmd2.Stdin = os.Stdin
	env := os.Environ()
	env = append(env, "R_LIBS_SITE=/Users/devin/clients/amgen/pkgsync/lib",
		"R_LIBS_USER=\"\"",
	)
	cmd2.Env = env
	fmt.Println("running cmd2")
	cmd2.Run()
}
