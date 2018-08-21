package fileutil

import (
	"path/filepath"

	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
	"github.com/sutd-statnlp/service-ladrawex/env"
)

// AbsPath returns the absolute path.
func AbsPath(path string) *string {
	path, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &path
}

// FullPath returns full path.
func FullPath(path string) string {
	return stringutil.Concat(env.BasePath, path)
}
