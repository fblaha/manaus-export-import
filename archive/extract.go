package archive

import (
	"github.com/mholt/archiver"
	"io/ioutil"
	"log"
	"os"
)

// ExtractToTmpDir extracts given archive to temp dir
func ExtractToTmpDir(archive string) (string, func(), error) {
	dir, err := ioutil.TempDir("", "mns-import")
	if err != nil {
		return "", deleteFunc(dir), err
	}
	if err = archiver.Unarchive(archive, dir); err != nil {
		return dir, deleteFunc(dir), err
	}

	return dir, deleteFunc(dir), nil
}

func deleteFunc(dir string) func() {
	if dir == "" {
		return func() {

		}
	}
	return func() {
		if err := os.RemoveAll(dir); err != nil {
			log.Fatal("unable to delete tmp dir:", dir)
		}
	}
}
