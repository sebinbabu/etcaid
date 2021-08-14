package etcaid

import (
	"os"
	"path/filepath"
	"unicode"

	"github.com/otiai10/copy"
)

// prepareAndCopy copies src node to dest. It overwrites files in case of conflict and
// skips symlinks. It is used for Backup & Restore.
//
// First it creates destination directory, then it copies all files recursively.
func prepareAndCopy(src string, dest string) error {
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

	return nil
}

// parseFilename returns the name of file and its extension after parsing the given filename.
// The (.) dot is included in extension.
func parseFilename(filename string) (string, string) {
	ext := filepath.Ext(filename)
	return filename[:len(filename)-len(filepath.Ext(ext))], ext
}

// isValidAppName checks if the given name only contains letters and numbers.
// If not, it will return false.
func isValidAppName(name string) bool {
	if name == "" {
		return false
	}

	for _, r := range name {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}
