package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/data/binding"
)

func main() {
	val := binding.NewString()
	callback := binding.NewDataListener(func() {
		fmt.Println("String changed to:", val.Get())
	})
	val.AddListener(callback)

	time.Sleep(time.Millisecond * 100)
}
