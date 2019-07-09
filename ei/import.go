package ei

import (
	"log"

	"github.com/fblaha/manaus-export-import/archive"
	"github.com/fblaha/manaus-export-import/config"
	"github.com/fblaha/manaus-export-import/rest"
)

// Import does the import
func Import(conf config.Conf) {
	tmpDir, purge, err := archive.ExtractToTmpDir(conf.ArchiveFile)
	// TODO handle Ctrl+C
	defer purge()

	if err != nil {
		log.Fatal("unable to extract archive : ", err)
	}

	idLoader := archive.NewIDLoader(tmpDir, ".json")
	dataLoader := archive.NewDataLoader(tmpDir, ".json")
	writer := rest.NewWriter(conf.FootprintsURL(), "application/json")
	importer := NewTransfer(idLoader, dataLoader, writer)
	err = importer.Execute(conf.Concurrency)
	if err != nil {
		log.Fatal("import failed : ", err)
	}
}
