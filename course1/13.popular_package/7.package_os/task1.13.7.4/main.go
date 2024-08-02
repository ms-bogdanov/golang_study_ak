package main

import (
	"fmt"
	"os/exec"
)

func ExecBin(binPath string, args ...string) string {
	out, err := exec.Command(binPath, args...).Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(out)
}
