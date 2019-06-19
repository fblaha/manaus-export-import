package main

import "fmt"
import "flag"

func main() {
	var url string
	var archive string

	flag.StringVar(&url, "url", "http://localhost:7777", "rest URL")
	flag.StringVar(&archive, "archive", "export.zip", "archive file")
	fmt.Println(url)
	fmt.Println(archive)
}
