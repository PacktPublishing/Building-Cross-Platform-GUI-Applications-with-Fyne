package main

import (
	"log"
	"os/exec"
)

func say(text string) {
	cmd := exec.Command("say", text)

	err := cmd.Run()
	if err != nil {
		log.Println("Error saying text", err)
	}
}
