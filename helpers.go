package etcaid

import (
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

// prepareAndCopy copies src node to dest. It overwrites files in case of conflict.
// It is used by Backup & Restore.
//
// First it creates destination directory, then it copies all files recursively.
func prepareAndCopy(src string, dest string, logger logger) error {
	_, err := os.Stat(src)
	if err != nil {
		return &ApplicationError{
			Op:      "copy",
			Message: "failed to access source node",
			Path:    src,
			Err:     err,
		}
	}

	destDir := filepath.Dir(dest)
	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return &ApplicationError{
			Op:      "copy",
			Message: "failed to create destination directory",
			Path:    destDir,
			Err:     err,
		}
	}

	err = copy.Copy(src, dest, copy.Options{
		PreserveTimes: true,
		OnSymlink: func(_ string) copy.SymlinkAction {
			return copy.Skip
		},
		OnDirExists: func(_ string, d string) copy.DirExistsAction {
			return copy.Replace
		},
	})
	if err != nil {
		return &ApplicationError{
			Op:      "copy",
			Message: "failed to copy to destination",
			Path:    dest,
			Err:     err,
		}
	}

	logger.Info("copied ", src, "to", dest)

	return nil
}

// parseFilename returns the name of file and its extension after parsing the given filename.
// The (.) dot is included in extension.
func parseFilename(filename string) (string, string) {
	ext := filepath.Ext(filename)
	return filename[:len(filename)-len(filepath.Ext(ext))], ext
}
