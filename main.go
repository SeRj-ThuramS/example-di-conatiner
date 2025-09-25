package main

import (
	"example-di-conatiner/container"
	"example-di-conatiner/module"
	"example-di-conatiner/str"
	"fmt"
)

func main() {
	_, err := container.RegisterRef[str.NewStruct](&str.NewStruct{
		Name: "new",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	module.Run()
}
