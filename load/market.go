package load

type marketLoader struct {
	mnsUrl string
	loader urlLoader
}

func newMarketLoader(mnsUrl string, loader urlLoader) marketLoader {
	return marketLoader{mnsUrl: mnsUrl, loader: loader}
}

func (p marketLoader) load(marketID string) ([]byte, error) {
	return p.loader(p.mnsUrl + "/markets/" + marketID)
}
