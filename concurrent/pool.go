package concurrent

import (
	"sync"
)

// Worker does work
type Worker interface {
	Work()
}

// PoolExecutor distributes works to pool of gouroutines
type PoolExecutor struct {
	workWG sync.WaitGroup
	poolWG sync.WaitGroup
	todo   chan Worker
}

// NewPoolExecutor constructor
func NewPoolExecutor(concurrency int) *PoolExecutor {
	todo := make(chan Worker)
	executor := PoolExecutor{todo: todo}
	executor.poolWG.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer executor.poolWG.Done()
			for w := range todo {
				w.Work()
				executor.workWG.Done()
			}
		}()
	}
	return &executor
}

// Submit submits work for execution
func (e *PoolExecutor) Submit(worker Worker) {
	e.workWG.Add(1)
	e.todo <- worker
}

// WaitShutdown waits for completeion of submitted work and terminates worker goroutines
func (e *PoolExecutor) WaitShutdown() {
	e.workWG.Wait()
	close(e.todo)
	e.poolWG.Wait()
}
