package main

import (
	"github.com/fblaha/manaus-export-import/config"
	"github.com/fblaha/manaus-export-import/ei"
)

func main() {
	ei.Import(config.LoadConfig())
}
