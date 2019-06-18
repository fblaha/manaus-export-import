package rest

type dataLoader struct {
	url string
	loader urlLoader
}

func newDataLoader(url string, loader urlLoader) dataLoader {
	return dataLoader{url: url, loader: loader}
}

func (p dataLoader) Load(id string) ([]byte, error) {
	return p.loader(p.url + id)
}
