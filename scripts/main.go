package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Specify the name of the environment variable you want to check
	envVarName := "FOO_BAR"

	// Get the value of the environment variable
	envVarValue := os.Getenv(envVarName)

	// Check if the environment variable is empty
	if len(envVarValue) == 0 {
		// If it's empty, return an error
		fmt.Printf("%s environment variable is not set.\n", envVarName)
		os.Exit(1)
	} else {
		// If it's not empty, you can run your desired command
		// For example, let's say you want to run "echo Hello, World!"
		cmd := exec.Command("echo", "Hello, World!")

		// Set the command's environment to inherit the current environment
		cmd.Env = os.Environ()

		// Run the command and capture its output
		output, err := cmd.CombinedOutput()

		if err != nil {
			// If there was an error running the command, print the error
			fmt.Printf("Error executing the command: %v\n", err)
			os.Exit(1)
		}

		// Print the command's output
		fmt.Printf("Command output:\n%s\n", output)
	}
}
