package concurrent

import (
	"sync"
)

// Worker does the work
type Worker interface {
	Work()
}

// PoolExecutor distributes works to the pool of gouroutines
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
		go executor.handleWork()
	}
	return &executor
}

func (e *PoolExecutor) handleWork() {
	defer e.poolWG.Done()
	for w := range e.todo {
		w.Work()
		e.workWG.Done()
	}
}

// Submit submits the work for execution
func (e *PoolExecutor) Submit(worker Worker) {
	e.workWG.Add(1)
	e.todo <- worker
}

// ShutdownGracefully waits for completeion of the submitted work and terminates worker goroutines
func (e *PoolExecutor) ShutdownGracefully() {
	e.workWG.Wait()
	close(e.todo)
	e.poolWG.Wait()
}
