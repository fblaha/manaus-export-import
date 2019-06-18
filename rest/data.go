package rest

type dataLoader struct {
	mnsURL string
	loader urlLoader
}

func newDataLoader(mnsURL string, loader urlLoader) dataLoader {
	return dataLoader{mnsURL: mnsURL, loader: loader}
}

func (p dataLoader) Load(id string) ([]byte, error) {
	return p.loader(p.mnsURL + id)
}
