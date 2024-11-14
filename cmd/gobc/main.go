package main

import (
	"fmt"
	"sync"
)

const (
	MAX_BC_SERVERS = 3
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < MAX_BC_SERVERS; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			newServerNode(fmt.Sprintf(":500%d", port)).InitAndRun()
		}(i)
	}
	wg.Wait()
}
