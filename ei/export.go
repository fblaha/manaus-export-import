package ei

import (
	"log"

	"github.com/fblaha/manaus-export-import/archive"
	"github.com/fblaha/manaus-export-import/config"
	"github.com/fblaha/manaus-export-import/rest"
)

type export struct {
	archive.Writer
	Transfer
}

func configureExport(conf config.Conf) (export, func(), error) {
	tempDir, purge, err := archive.CreateTempDir()
	if err != nil {
		return export{}, purge, err
	}
	log.Println("temp directory created:", tempDir)
	Writer := archive.NewWriter(tempDir, ".json")

	idLoader := rest.NewIDLoader(conf.MarketIDsURL(), rest.LoadURL)
	dataLoader := rest.NewDataLoader(conf.FootprintsURL(), rest.LoadURL)
	transfer := NewTransfer(idLoader, dataLoader, Writer)
	return export{Writer: Writer, Transfer: transfer}, purge, nil

}

// Export does the export
func Export(conf config.Conf) {
	export, purge, err := configureExport(conf)
	if err != nil {
		log.Fatal("unable to configure export:", err)
	}
	defer purge()
	err = export.Execute(conf.Concurrency)
	if err != nil {
		log.Fatal("data transfer failed:", err)
	}
	err = export.MakeArchive(conf.ArchiveFile)
	if err != nil {
		log.Fatal("unable to create archive:", err)
	}
}
