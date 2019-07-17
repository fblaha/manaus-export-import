package archive

import (
	"github.com/mholt/archiver"
	"log"
)

// ExtractToTmpDir extracts given archive to temp dir
func ExtractToTmpDir(archive string) (string, func(), error) {
	dir, purge, err := CreateTempDir()
	if err != nil {
		return "", purge, err
	}
	log.Println("extracting archive to temp dir:", dir)
	if err = archiver.Unarchive(archive, dir); err != nil {
		return dir, purge, err
	}

	return dir, purge, nil
}
