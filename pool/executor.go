package pool

import (
	"context"
	"sync"
)

// Worker does work
type Worker interface {
	Work(ctx context.Context)
}

// Executor distributes works to pool of gouroutines
type Executor struct {
	todo chan Worker
}

// NewExecutor constructor
func NewExecutor(parentCtx context.Context, concurrency int) (Executor, func()) {
	todo := make(chan Worker)
	ctx, cancel := context.WithCancel(parentCtx)

	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case w, ok := <-todo:
					if ok {
						w.Work(ctx)
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	return Executor{todo: todo}, func() {
		defer close(todo)
		cancel()
		wg.Wait()
	}
}

// Submit submits work for execution
func (e Executor) Submit(worker Worker) {
	e.todo <- worker
}
