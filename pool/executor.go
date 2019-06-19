package pool

import (
	"sync"
)

// Worker does work
type Worker interface {
	Work()
}

// Executor distributes works to pool of gouroutines
type Executor struct {
	todo chan Worker
}

// NewExecutor constructor
func NewExecutor(concurrency int) (Executor, func()) {
	todo := make(chan Worker)

	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for w := range todo {
				w.Work()
			}
		}()
	}
	return Executor{todo: todo}, func() {
		close(todo)
		wg.Wait()
	}
}

// Submit submits work for execution
func (e Executor) Submit(worker Worker) {
	e.todo <- worker
}
