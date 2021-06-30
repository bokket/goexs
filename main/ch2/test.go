package main

import "fmt"

func main()  {

	var p=f()

	fmt.Printf("%v\n",p)

	var s=f()
	fmt.Printf("%v",s)
}

func f()*int  {
	v:=1
	return &v
}