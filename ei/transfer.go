package ei

import (
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
func (t Transfer) Execute() error {
	// TODO hardcoded constants
	executor, shutdown := pool.NewExecutor(10)
	defer shutdown()
	ids, err := t.LoadIDs()
	if err != nil {
		return errors.Wrap(err, "unable to read ids to move")
	}

	// TODO hardcoded constants
	results := make(chan transferResult, 0)
	defer close(results)

	submitted := t.submitTransferWork(executor, ids, results)

	return t.collectResults(submitted, results)
}

func (t Transfer) submitTransferWork(executor pool.Executor, ids []string, results chan<- transferResult) int {
	for _, id := range ids {
		executor.Submit(transferWorker{
			DataLoader: t.DataLoader,
			DataWriter: t.DataWriter,
			id:         id,
			c:          results,
		})
	}
	return len(ids)
}

func (t Transfer) collectResults(count int, writeResults <-chan transferResult) (err error) {
	for i := 0; i < count; i++ {
		wr := <-writeResults
		if wr.err != nil {
			err = wr.err
		}
	}
	return err
}
