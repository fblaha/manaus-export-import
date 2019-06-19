package ei

type exportResult struct {
	id  string
	err error
}

type exportWorker struct {
	DataLoader
	DataWriter
	id string
	c  chan<- exportResult
}

func (w exportWorker) Work() {
	data, err := w.Load(w.id)
	if err != nil {
		w.c <- exportResult{w.id, err}
		return
	}
	err = w.Write(w.id, data)
	w.c <- exportResult{w.id, err}
}
