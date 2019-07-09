package ei

import (
	"github.com/fblaha/manaus-export-import/archive"
	"github.com/fblaha/manaus-export-import/config"
	"github.com/fblaha/manaus-export-import/rest"
	"log"
)

// Import does the import
func Import(conf config.Conf) {
	tmpDir, cleanUp, err := archive.ExtractToTmpDir(conf.ArchiveFile)
	if cleanUp != nil {
		defer cleanUp()
	}
	if err != nil {
		log.Fatal("unable to extract archive : ", err)
	}

	idLoader := archive.NewIDLoader(tmpDir, ".json")
	dataLoader := archive.NewDataLoader(tmpDir, ".json")
	writer := rest.NewWriter(conf.URL, "application/json")
	importer := NewTransfer(idLoader, dataLoader, writer)
	err = importer.Execute(conf.Concurrency)
	if err != nil {
		log.Fatal("import failed : ", err)
	}
}
