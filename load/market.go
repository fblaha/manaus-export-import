package load

type marketLoader struct {
	mnsURL string
	loader urlLoader
}

func newMarketLoader(mnsURL string, loader urlLoader) marketLoader {
	return marketLoader{mnsURL: mnsURL, loader: loader}
}

func (p marketLoader) load(marketID string) ([]byte, error) {
	return p.loader(p.mnsURL + "/markets/" + marketID)
}
