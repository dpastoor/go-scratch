package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func ExecWithContextTimer() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	fmt.Println("started...")
	start_time := time.Now()
	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
		fmt.Println("ended...", time.Since(start_time))
	}
}

func ExecWithExternalCancellation() {
	cmd := exec.Command("sleep", "5")

	// Use a bytes.Buffer to get the output
	var buf bytes.Buffer
	cmd.Stdout = &buf

	cmd.Start()

	// Use a channel to signal completion so we can use a select statement
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	// Start a timer
	timeout := time.After(500 * time.Millisecond)
	startTime := time.Now()

	// The select statement allows us to execute based on which channel
	// we get a message from first.
	select {
	case <-timeout:
		// Timeout happened first, kill the process and print a message.
		cmd.Process.Kill()
		fmt.Println("Command timed out")
		fmt.Println("ended...", time.Since(startTime))
	case err := <-done:
		// Command completed before timeout. Print output and error if it exists.
		fmt.Println("Output:", "success")
		fmt.Println("ended...", time.Since(startTime))
		if err != nil {
			fmt.Println("Non-zero exit code:", err)
		}
	}
}
