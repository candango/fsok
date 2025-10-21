package dir

import (
	"os"
)

// Exists reports whether the specified path exists and is a directory.
//
// It returns true if the path exists and is a directory, and false otherwise.
// This function does not follow symlinks; it checks the path as provided.
//
// Example usage:
//
//	if dir.Exists("/tmp/mydir") {
//	    // Directory exists
//	}
func Exists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
