package main

import (
	"fmt"
)

func main() {
	strings := []string{"a", "b", "c"}
	fmt.Printf("%v\n", strings)
	fmt.Printf("%v\n", &strings[0])
	echo(strings)
	fmt.Printf("%v\n", strings)
}

func echo(param []string) {
	fmt.Printf("%v\n", &param[0])
	param[0] = "nb"
	param = append(param, "d")
	fmt.Printf("%v\n", param)
}
