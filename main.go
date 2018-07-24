package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	err := npm()
	if err != nil {
		log.Fatal(err)
	}
}

func npm() error {
	cmd := exec.Command("npm", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return os.RemoveAll("./node_modules")
}
