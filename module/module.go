package module

import (
	"example-di-conatiner/container"
	"example-di-conatiner/str"
	"fmt"
)

func Run() {
	fmt.Println(container.AssignRef[str.NewStruct]())
}
