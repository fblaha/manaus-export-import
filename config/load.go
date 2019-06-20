package config

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"
)

// Conf parsed cli configuration
type Conf struct {
	URL         string
	ArchiveFile string
	Concurrency int
}

// LoadConfig parse cli args
func LoadConfig() Conf {
	var url string
	var archive string
	var concurrency int

	flag.StringVar(&url, "url", "http://localhost:7777", "rest URL")
	flag.StringVar(&archive, "archive", archiveFile(), "archive file")
	flag.IntVar(&concurrency, "concurrency", runtime.NumCPU(), "goroutines count")
	conf := Conf{ArchiveFile: archive, Concurrency: concurrency, URL: url}
	log.Printf("loaded config : %+v", conf)
	return conf
}

func archiveFile() string {
	time := time.Now().Format("2006-01-02-15-04-05")
	return fmt.Sprintf("export%s.zip", time)
}
