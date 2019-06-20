package ei

import (
	"github.com/fblaha/manaus-export-import/config"
	"github.com/fblaha/manaus-export-import/rest"
	"github.com/fblaha/manaus-export-import/storage"
	"io/ioutil"
	"log"
)

type export struct {
	storage.DirectoryWriter
	Transfer
}

func configureExport(conf config.Conf) (export, error) {
	tempDir, err := ioutil.TempDir("", "mns-export")
	if err != nil {
		return export{}, err
	}
	log.Println("temp directory created : ", tempDir)
	directoryWriter := storage.NewDirectoryWriter(tempDir, ".json")

	idLoader := rest.NewIDLoader(conf.URL+"/market-ids/", rest.LoadURL)
	dataLoader := rest.NewDataLoader(conf.URL+"/footprints/", rest.LoadURL)
	transfer := NewTransfer(idLoader, dataLoader, directoryWriter)
	return export{DirectoryWriter: directoryWriter, Transfer: transfer}, nil

}

// Export does the export
func Export(conf config.Conf) {
	export, err := configureExport(conf)
	if err != nil {
		log.Fatal("unable to configure export : ", err)
	}
	defer func() {
		if err := export.Purge(); err != nil {
			log.Fatal("unable delete temp dir : ", err)
		}
	}()
	err = export.Execute(conf.Concurrency)
	if err != nil {
		log.Fatal("data transfer failed : ", err)
	}
	err = export.MakeArchive(conf.ArchiveFile)
	if err != nil {
		log.Fatal("unable to create archive : ", err)
	}
}
