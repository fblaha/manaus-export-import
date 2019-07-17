package ei

import (
	"github.com/fblaha/pool"
)

type transferResult struct {
	id  string
	err error
}

func createTransferWork(id string, loader DataLoader, writer DataWriter, output chan<- transferResult) pool.WorkerFunc {
	return func() {
		data, err := loader.Load(id)
		if err != nil {
			output <- transferResult{id, err}
			return
		}
		err = writer.Write(id, data)
		output <- transferResult{id, err}
	}
}
