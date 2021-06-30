package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts:=make(map[string]int)

	files:=os.Args[1:]

	if len(files)==0 {
		coutLines(os.Stdin,counts)
	} else {
		for _,arg:=range files {
			f,err := os.Open(arg)
			if err!=nil {
				fmt.Println(os.Stderr,"dup2:%v\n",err)
				continue
			}
			coutLines(f,counts)
			f.Close()
		}
	}


	for line,n:=range counts{
		if n>1 {
			//fmt.Println(n)
			//fmt.Println(line)
			fmt.Printf("%s%d",line,n)
		}
	}
}

func coutLines(f *os.File,counts map[string]int)  {
	input:=bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}
