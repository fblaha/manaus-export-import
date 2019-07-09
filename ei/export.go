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
	Purge func()
}

func configureExport(conf config.Conf) (export, error) {
	tempDir, purge, err := archive.CreateTempDir()
	if err != nil {
		return export{}, err
	}
	log.Println("temp directory created:", tempDir)
	Writer := archive.NewWriter(tempDir, ".json")

	idLoader := rest.NewIDLoader(conf.URL+"/market-ids/", rest.LoadURL)
	dataLoader := rest.NewDataLoader(conf.URL+"/footprints/", rest.LoadURL)
	transfer := NewTransfer(idLoader, dataLoader, Writer)
	return export{Writer: Writer, Transfer: transfer, Purge: purge}, nil

}

// Export does the export
func Export(conf config.Conf) {
	export, err := configureExport(conf)
	if err != nil {
		log.Fatal("unable to configure export : ", err)
	}
	defer export.Purge()
	err = export.Execute(conf.Concurrency)
	if err != nil {
		log.Fatal("data transfer failed : ", err)
	}
	err = export.MakeArchive(conf.ArchiveFile)
	if err != nil {
		log.Fatal("unable to create archive : ", err)
	}
}
