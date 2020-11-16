package main

import (
	"fmt"

	"fyne.io/fyne/app"
)

func main() {
	a := app.NewWithID("com.example.preferences")

	key := "anotherkey"
	val := a.Preferences().String(key)
	fmt.Println("Value is:", val)
	val = a.Preferences().StringWithFallback(key, "missing")
	fmt.Println("Value is:", val)

	a.Preferences().SetString(key, "somevalue")
	val = a.Preferences().String(key)
	fmt.Println("Value is:", val)

	fmt.Println("Removing")
	a.Preferences().RemoveValue(key)
	val = a.Preferences().String(key)
	fmt.Println("Value is:", val)
}
