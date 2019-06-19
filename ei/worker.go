package ei

type transferResult struct {
	id  string
	err error
}

type transferWorker struct {
	DataLoader
	DataWriter
	id string
	c  chan<- transferResult
}

func (w transferWorker) Work() {
	data, err := w.Load(w.id)
	if err != nil {
		w.c <- transferResult{w.id, err}
		return
	}
	err = w.Write(w.id, data)
	w.c <- transferResult{w.id, err}
}
