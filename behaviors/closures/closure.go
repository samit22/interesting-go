/* interesting go memory management, Go routine is smart enough to know
that a reference to the name is still being used so it transfers
the memory to the heap so that the routine can continue to access it.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, name := range []string{"ram", "shyam", "hari"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(name)
		}()
	}
	wg.Wait()
}
