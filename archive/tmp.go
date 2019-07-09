package archive

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
)

// CreateTempDir creates and logs
func CreateTempDir() (string, func(), error) {
	tempDir, err := ioutil.TempDir("", "mns-export-import-")
	if err != nil {
		log.Println("unable to create tmp dir:", err)
		return "", nop, errors.Wrap(err, "unable to create tmp dir")
	}
	log.Println("using tmp dir:", tempDir)
	return tempDir, purgeFunc(tempDir), err
}

func nop() {

}

func purgeFunc(dir string) func() {
	return func() {
		if err := os.RemoveAll(dir); err != nil {
			log.Panic("unable to delete tmp dir:", dir)
		}
	}
}
