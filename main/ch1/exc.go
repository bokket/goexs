package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Println(strings.Join(os.Args[0:1],""))
	fmt.Println(os.Args[0])

	for i,s:=range os.Args[1:] {
		fmt.Println(i,s)
	}

}