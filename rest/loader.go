package rest

import "log"

// DataLoader loads data from given URL
type DataLoader struct {
	url    string
	loader URLLoader
}

// NewDataLoader constructor of URL based data loader
func NewDataLoader(url string, loader URLLoader) DataLoader {
	return DataLoader{url: url, loader: loader}
}

// Load loads data for given ID
func (l DataLoader) Load(id string) ([]byte, error) {
	url := l.url + id
	log.Println("loading data from:", url)
	return l.loader(url)
}
