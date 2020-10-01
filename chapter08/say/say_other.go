// +build !darwin

package main

import (
	"log"
	"runtime"
)

func say(_ string) {
	log.Println("Say support is not available for", runtime.GOOS)
}
