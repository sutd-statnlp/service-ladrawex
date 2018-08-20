package fileutil

import (
	"path/filepath"

	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/env"
	"github.com/sutd-statnlp/service-ladrawex/log"
	"github.com/sutd-statnlp/service-ladrawex/util/stringutil"
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

// FullPath returns full path.
func FullPath(path string) string {
	return stringutil.Concat(env.GoPath, "/src/", constant.BasePath, path)
}
