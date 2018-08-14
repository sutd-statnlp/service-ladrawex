package fileutil

import (
	"path/filepath"

	"github.com/sutd-statnlp/service-ladrawex/log"
)

// AbsPath returns the absolute path.
func AbsPath(path string) *string {
	log.Debug("Request to get absolute path: ", path)
	path, err := filepath.Abs(path)
	if err != nil {
		log.Error(err)
	}
	return &path
}
