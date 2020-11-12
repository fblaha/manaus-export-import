package config

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

// Conf parsed cli configuration
type Conf struct {
	URL         string
	ArchiveFile string
	Concurrency int
}

// FootprintsURL according to the name
func (c Conf) FootprintsURL() string {
	return strings.TrimSuffix(c.URL, "/") + "/footprints/"
}

// MarketIDsURL according to the name
func (c Conf) MarketIDsURL() string {
	return strings.TrimSuffix(c.URL, "/") + "/market-ids/"
}

// LoadConfig parse cli args
func LoadConfig() Conf {
	var url string
	var archive string
	var concurrency int

	flag.StringVar(&url, "url", "http://manaus:8080/", "rest URL")
	flag.StringVar(&archive, "archive", archiveFile(), "archive file")
	flag.IntVar(&concurrency, "concurrency", runtime.NumCPU(), "goroutines count")
	flag.Parse()
	conf := Conf{ArchiveFile: archive, Concurrency: concurrency, URL: url}
	log.Printf("loaded config : %+v", conf)
	return conf
}

func archiveFile() string {
	now := time.Now().Format("2006-01-02-15-04-05")
	return fmt.Sprintf("export%s.zip", now)
}
