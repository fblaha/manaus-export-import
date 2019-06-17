package pool

import (
	"context"
	"sync"
	"testing"
)

func TestExecutorStartShutdown(t *testing.T) {
	_, shutdown := NewExecutor(context.Background(), 10)
	shutdown()
}

type mockWork struct {
	*sync.WaitGroup
}

func (w *mockWork) Work(ctx context.Context) {
	w.Done()
}

func TestExecutorSubmit(t *testing.T) {
	executor, shutdown := NewExecutor(context.Background(), 10)
	defer shutdown()
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		work := mockWork{&wg}
		executor.Submit(&work)
	}
	wg.Wait()
}
