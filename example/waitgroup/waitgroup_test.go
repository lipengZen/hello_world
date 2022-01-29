package waitgroup

import (
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestWaitGroup(t *testing.T) {

	go func() {
		time.Sleep(time.Second)
		wg.Add(1)
		t.Log("add 1")
	}()
	wg.Wait()
	time.Sleep(time.Second * 2)

	// wg.Done()

	// wg.Add(1)

	t.Log("end")
}
