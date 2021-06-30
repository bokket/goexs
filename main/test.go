package main

import (
	"fmt"
	"strings"
)

/*
func main()  {
	fmt.Print("hello,world")
	s := "hello"
	s="c"+s[1:]
	fmt.Println(s)

	c :=s[:len(s)]
	d :=s[len(s)/2:]
	fmt.Println(c)
	fmt.Println(d)
}*/

// Example URL: https://bucket.zone.qingstor.com
// Example URL: https://zone.qingstor.com/bucket

func main()  {
	/*s := "https://bucket.zone.qingstor.com"
	u := "https://zone.qingstor.com/bucket"

	fmt.Println(s)
	fmt.Println(u)

	location := strings.Split(s,".")[1]
	fmt.Println(location)

	location1 := strings.Split(u,".")[0]
	location1 = strings.Split(location1,"//")[1]
	fmt.Println(location1)
	fmt.Println(location1)

	base,_ :=url.Parse(u)

	fmt.Println(base.Host)
	location = strings.Split(base.Hostname(),".")[0]
	fmt.Println(location)*/
	//<type>://[<name>][<work_dir>][?key1=value1&...&keyN=valueN]
	s :="<type>://[<name>][<work_dir>][?key1=value1&...&keyN=valueN]"
	colon :=strings.Index(s,":");
	fmt.Println(colon)

	ty:=s[:colon]
	reset:=s[colon+1:]
	fmt.Println(ty)
	fmt.Println(reset)

	v:=strings.HasPrefix(reset,"//")
	fmt.Println(v)

	reset=reset[2:]
	fmt.Println(reset)

	question:=strings.Index(reset,"?")
	fmt.Println("reset:",question)

	path:=reset[:question]
	fmt.Println("path:",path)

	reset=reset[question+1:]
	fmt.Println(reset)

}