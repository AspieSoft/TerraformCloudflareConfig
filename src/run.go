package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	regex "github.com/AspieSoft/go-regex/v5/re2"
	"github.com/klauspost/cpuid/v2"
)

func main(){
	// detect windows
	var winDir string
	if runtime.GOOS == "windows" {
		winPath, err := filepath.Abs("./windows")
		if err != nil {
			panic(err)
		}

		if regex.Comp(`(?i)(^|[^A-Za-z])amd($|[^A-Za-z])`).Match([]byte(cpuid.CPU.BrandName)) {
			winDir = winPath+"/amd64"
		}else if regex.Comp(`(?i)(^|[^A-Za-z])intel($|[^A-Za-z])`).Match([]byte(cpuid.CPU.BrandName)) {
			winDir = winPath+"/386"
		}

		if file, err := os.ReadFile(winDir+"/terraform.exe"); err == nil {
			os.WriteFile("terraform.exe", file, 100755)
		}
	}


	resetTerraform()

	autoYes := false
	for _, arg := range os.Args[1:] {
		if arg == "-y" || arg == "--autoyes" {
			autoYes = true
		}
	}

	file, err := os.ReadFile("cloudflare.tf")
	if err != nil {
		panic(err)
	}

	oldFile := file

	hasModifiedFile := false
	if stat, err := os.Stat("cloudflare_api_token.key"); err == nil && !stat.IsDir() {
		key, err := os.ReadFile("cloudflare_api_token.key")
		if err != nil {
			panic(err)
		}
		key = bytes.TrimRight(key, "\r\n")

		if len(key) != 0 {
			hasModifiedFile = true
			file = regexp.MustCompile(`#\s*api_token\s*=`).ReplaceAll(file, []byte("api_token ="))
			file = bytes.ReplaceAll(file, []byte("<Insert Cloudflare API Token>"), key)
		}
	}

	if bytes.Contains(file, []byte("<Insert Zone ID>")) {
		if zoneListFile, err := os.ReadFile("zone.list"); err == nil && len(zoneListFile) != 0 && len(bytes.TrimRight(zoneListFile, "\r\n")) != 0 {
			defer os.WriteFile("cloudflare.tf", oldFile, 0644)

			curFile := file

			zoneList := bytes.Split(zoneListFile, []byte{'\n'})
			for _, zone := range zoneList {
				zone = bytes.TrimRight(zone, "\r\n")
				if len(zone) == 0 {
					break
				}

				file = bytes.ReplaceAll(file, []byte("<Insert Zone ID>"), []byte(zone))
				os.WriteFile("cloudflare.tf", file, 0644)

				run(true, winDir)

				file = curFile
				resetTerraform()
			}

			os.WriteFile("cloudflare.tf", oldFile, 0644)
			resetTerraform()
			return
		}else{
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

			hasModifiedFile = true
			file = bytes.ReplaceAll(file, []byte("<Insert Zone ID>"), []byte(text))
		}
	}

	if hasModifiedFile {
		defer os.WriteFile("cloudflare.tf", oldFile, 0644)
		os.WriteFile("cloudflare.tf", file, 0644)
	}

	run(autoYes, winDir)

	if hasModifiedFile {
		os.WriteFile("cloudflare.tf", oldFile, 0644)
		resetTerraform()
	}
}

func run(autoYes bool, winDir string){
	var cmd *exec.Cmd
	
	if winDir != "" {
		cmd = exec.Command(`.\terraform`, `init`)
		// cmd.Dir = winDir
	}else{
		cmd = exec.Command(`terraform`, `init`)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err == nil {
		if winDir != "" {
			cmd = exec.Command(`.\terraform`, `plan`)
			// cmd.Dir = winDir
		}else{
			cmd = exec.Command(`terraform`, `plan`)
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	if err == nil {
		if winDir != "" {
			if autoYes {
				cmd = exec.Command(`.\terraform`, `apply`, `-auto-approve`)
			}else{
				cmd = exec.Command(`.\terraform`, `apply`)
			}
			// cmd.Dir = winDir
		}else if autoYes {
			cmd = exec.Command(`terraform`, `apply`, `-auto-approve`)
		}else{
			cmd = exec.Command(`terraform`, `apply`)
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

func resetTerraform(){
	os.RemoveAll(".terraform")
	os.RemoveAll(".terraform.lock.hcl")
	os.RemoveAll(".terraform.tfstate")
	os.RemoveAll("terraform.tfstate")
	os.RemoveAll("terraform.tfstate.backup")
}
