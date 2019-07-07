package ei

import (
	"sync"

	"github.com/fblaha/manaus-export-import/pool"
	"github.com/pkg/errors"
)

// Transfer does data transfer
type Transfer struct {
	IDLoader
	DataLoader
	DataWriter
}

// NewTransfer constructor
func NewTransfer(
	idLoader IDLoader,
	dataLoader DataLoader,
	dataWriter DataWriter) Transfer {
	return Transfer{IDLoader: idLoader, DataLoader: dataLoader, DataWriter: dataWriter}
}

// Execute executes data transfer
func (t Transfer) Execute(concurrency int) error {
	executor, shutdown := pool.NewExecutor(concurrency)
	defer shutdown()
	ids, err := t.LoadIDs()
	if err != nil {
		return errors.Wrap(err, "unable to read ids to transfer")
	}

	results := make(chan transferResult, 1)
	defer close(results)

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		t.submit(ids, executor, results)
	}()

	return t.collectResults(len(ids), results)
}

func (t Transfer) submit(ids []string, executor pool.Executor, results chan<- transferResult) {
	for _, id := range ids {
		executor.Submit(transferWorker{
			DataLoader: t.DataLoader,
			DataWriter: t.DataWriter,
			id:         id,
			c:          results,
		})
	}
}

func (t Transfer) collectResults(count int, results <-chan transferResult) (err error) {
	for i := 0; i < count; i++ {
		wr := <-results
		if wr.err != nil {
			err = wr.err
		}
	}
	return err
}
