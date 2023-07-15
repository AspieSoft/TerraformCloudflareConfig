package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main(){
	resetTerraform()

	file, err := os.ReadFile("cloudflare.tf")
	if err != nil {
		panic(err)
	}

	hasModifiedFile := false
	if bytes.Contains(file, []byte("<Insert Zone ID>")) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter Zone ID: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		text = strings.TrimRight(text, "\r\n")

		if text == "" {
			panic("error: Zone ID Not Specified")
		}

		os.WriteFile("", bytes.ReplaceAll(file, []byte("<Insert Zone ID>"), []byte(text)), 0644)

		hasModifiedFile = true
	}

	cmd := exec.Command(`terraform`, `init`)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err == nil {
		cmd = exec.Command(`terraform`, `plan`)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	if err == nil {
		cmd = exec.Command(`terraform`, `apply`)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	if hasModifiedFile {
		os.WriteFile("", file, 0644)
		resetTerraform()
	}
}

func resetTerraform(){
	os.RemoveAll(".terraform")
	os.RemoveAll(".terraform.lock.hcl")
	os.RemoveAll(".terraform.tfstate")
	os.RemoveAll("terraform.tfstate")
	os.RemoveAll("terraform.tfstate.backup")
}
