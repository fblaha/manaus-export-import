package pool

import (
	"sync"
	"testing"
)

func TestExecutorStartShutdown(t *testing.T) {
	_, shutdown := NewExecutor(10)
	shutdown()
}

type mockWork struct {
	*sync.WaitGroup
}

func (w *mockWork) Work() {
	w.Done()
}

func TestExecutorSubmit(t *testing.T) {
	executor, shutdown := NewExecutor(10)
	defer shutdown()
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		work := mockWork{&wg}
		executor.Submit(&work)
	}
	wg.Wait()
}
