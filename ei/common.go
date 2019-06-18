package ei

// IDLoader capable to load list of IDs
type IDLoader interface {
	LoadIDs() ([]string, error)
}

// DataLoader capable of loading data associaated with given ID
type DataLoader interface {
	Load(id string) ([]byte, error)
}

// DataWriter writes data associated with given ID
type DataWriter interface {
	Write(id string, data []byte) error
}
