package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	context:=make(map[string]int)

	input:=bufio.NewScanner(os.Stdin)

	for input.Scan() {
		context[input.Text()]++
		fmt.Println(">")
	}

	for line,n:=range context {
		if n>1 {
			fmt.Printf("%s%d",line,n)
		}
	}

}
