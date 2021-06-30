package main

import (
	"fmt"
	"testing"
	"time"
)

func TestEx1(t *testing.T)  {
	var data int

	go func() {
		data++
	}()

	time.Sleep(time.Microsecond*200)

	if data==0 {
		fmt.Printf("the value is %v.\n",data)
	}
}