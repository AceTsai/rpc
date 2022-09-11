package X

import (
	"fmt"
	"sync"
)

func M() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go c1(wg)
	wg.Wait()
	fmt.Println(11111)
}

func c1(w *sync.WaitGroup) {
	w.Done()
}
