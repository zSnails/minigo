package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const EXTENSION = ".exe"

func run(executable string) int {
	path, err := filepath.Abs(executable)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return InternalError
	}
	cmd := exec.Command(fmt.Sprintf("%s.exe", path))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExternalError
	}

	return 0
}
