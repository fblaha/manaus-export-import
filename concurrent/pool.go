package concurrent

import (
	"sync"
)

// Worker does the work
type Worker interface {
	Work()
}

// PoolExecutor distributes works to the pool of goroutines
type PoolExecutor struct {
	// tracks completion of submitted work
	workWG sync.WaitGroup
	// tracks pool goroutines which process the incoming work
	poolWG sync.WaitGroup
	// incoming work
	todo chan Worker
}

// NewPoolExecutor constructor
func NewPoolExecutor(concurrency int) *PoolExecutor {
	todo := make(chan Worker)
	executor := PoolExecutor{todo: todo}
	for i := 0; i < concurrency; i++ {
		executor.poolWG.Add(1)
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
func (e *PoolExecutor) Submit(workers ...Worker) {
	for _, w := range workers {
		// ensures that shutdown waits for completion of submitted work
		e.workWG.Add(1)
		// submits work
		e.todo <- w
	}
}

// ShutdownGracefully waits for completion of the submitted work and terminates worker goroutines
// and frees allocated resources. The executor can no longer be used after this call.
func (e *PoolExecutor) ShutdownGracefully() {
	// waits for completion of submitted work
	e.workWG.Wait()
	close(e.todo)
	// waits for completion of the pool goroutines
	e.poolWG.Wait()
}
