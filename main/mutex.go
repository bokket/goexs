package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex

	var value int

	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()

	if value==0 {
		fmt.Printf("the value is %d",value)
	} else {
		fmt.Printf("the value is %d\n",value)
	}

	memoryAccess.Unlock()
}