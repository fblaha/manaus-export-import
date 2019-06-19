package ei

import (
	"github.com/fblaha/manaus-export-import/pool"
	"github.com/pkg/errors"
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

	exportResults := make(chan exportResult, 0)
	defer close(exportResults)

	submitted := e.submitExportWork(executor, ids, exportResults)

	return e.collectResults(submitted, exportResults)
}

func (e Exporter) submitExportWork(executor pool.Executor, ids []string, exportResults chan<- exportResult) int {
	for _, id := range ids {
		executor.Submit(exportWorker{
			DataLoader: e.DataLoader,
			DataWriter: e.DataWriter,
			id:         id,
			c:          exportResults,
		})
	}
	return len(ids)
}

func (e Exporter) collectResults(count int, writeResults <-chan exportResult) (err error) {
	for i := 0; i < count; i++ {
		wr := <-writeResults
		if wr.err != nil {
			err = wr.err
		}
	}
	return err
}
