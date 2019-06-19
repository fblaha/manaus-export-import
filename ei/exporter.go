package ei

import (
	"github.com/fblaha/manaus-export-import/pool"
	"github.com/pkg/errors"
	"sync"
)

// Exporter does export
type Exporter struct {
	IDLoader
	DataLoader
	DataWriter
}

// NewExporter constructor
func NewExporter(
	idLoader IDLoader,
	dataLoader DataLoader,
	dataWriter DataWriter) Exporter {
	return Exporter{IDLoader: idLoader, DataLoader: dataLoader, DataWriter: dataWriter}
}

// Execute executes export
func (e Exporter) Execute() error {
	executor, shutdown := pool.NewExecutor(10)
	defer shutdown()
	ids, err := e.LoadIDs()
	if err != nil {
		return errors.Wrap(err, "unable to read ids to export")
	}
	var wg sync.WaitGroup
	defer wg.Wait()

	loadResults := make(chan loadResult, 0)
	defer close(loadResults)

	submitted := e.submitLoadDataWork(executor, ids, loadResults)

	writeResults := make(chan writeResult, 0)
	defer close(writeResults)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for lr := range loadResults {
			executor.Submit(writeWorker{DataWriter: e.DataWriter, loadResult: lr, c: writeResults})
		}
	}()
	return e.collectWriteResults(submitted, writeResults)
}

func (e Exporter) submitLoadDataWork(executor pool.Executor, ids []string, loadResults chan<- loadResult) int {
	for _, id := range ids {
		executor.Submit(loadWorker{DataLoader: e.DataLoader, id: id, c: loadResults})
	}
	return len(ids)
}

func (e Exporter) collectWriteResults(count int, writeResults <-chan writeResult) (err error) {
	for i := 0; i < count; i++ {
		wr := <-writeResults
		if wr.err != nil {
			err = wr.err
		}
	}
	return err
}

type loadResult struct {
	id   string
	data []byte
	err  error
}

type loadWorker struct {
	DataLoader
	id string
	c  chan<- loadResult
}

func (w loadWorker) Work() {
	data, err := w.Load(w.id)
	w.c <- loadResult{w.id, data, err}
}

type writeResult struct {
	id  string
	err error
}

type writeWorker struct {
	DataWriter
	loadResult
	c chan<- writeResult
}

func (w writeWorker) Work() {
	err := w.err
	if err == nil {
		err = w.Write(w.id, w.data)
	}
	w.c <- writeResult{w.id, err}
}
