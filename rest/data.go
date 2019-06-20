package rest

import "log"

// URLDataLoader loads data from given URL
type URLDataLoader struct {
	url    string
	loader urlLoader
}

// NewDataLoader constructor of URL based data loader
func NewDataLoader(url string, loader urlLoader) URLDataLoader {
	return URLDataLoader{url: url, loader: loader}
}

// Load loads data for given ID
func (p URLDataLoader) Load(id string) ([]byte, error) {
	url := p.url + id
	log.Println("loading data from :", url)
	return p.loader(url)
}
