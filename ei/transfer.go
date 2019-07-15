package ei

import (
	"log"
	"sync"

	"github.com/fblaha/pool"
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
	ids, err := t.LoadIDs()
	if err != nil {
		return errors.Wrap(err, "unable to read ids to transfer")
	}

	results := make(chan transferResult, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		t.submitWait(ids, concurrency, results)
		close(results)
	}()

	return t.collectResults(results)
}

func (t Transfer) submitWait(ids []string, concurrency int, results chan<- transferResult) {
	executor := pool.NewExecutor(concurrency)
	for _, id := range ids {
		executor.Submit(transferWorker{
			DataLoader: t.DataLoader,
			DataWriter: t.DataWriter,
			id:         id,
			c:          results,
		})
	}
	executor.ShutdownGracefully()
}

func (t Transfer) collectResults(results <-chan transferResult) (err error) {
	for wr := range results {
		if wr.err != nil {
			log.Printf("data transfer failed id: %s  error: %+v", wr.id, wr.err)
			err = wr.err
		}
	}
	return err
}
