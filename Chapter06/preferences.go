package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
)

func main() {
	a := app.NewWithID("com.example.preferences")

	key := "demokey"
	a.Preferences().SetString(key, "somevalue")
	val := a.Preferences().String(key)
	fmt.Println("Value is:", val)

	key = "anotherkey"
	val = a.Preferences().String(key)
	fmt.Println("Value is:", val)
	val = a.Preferences().StringWithFallback(key, "missing")
	fmt.Println("Value is:", val)

	fmt.Println("Removing")
	a.Preferences().RemoveValue(key)
	val = a.Preferences().String(key)
	fmt.Println("Value is:", val)

	data := binding.BindPreferenceString("demokey", a.Preferences())
	val, _ = data.Get()
	fmt.Println("Bound value:", val)
}
