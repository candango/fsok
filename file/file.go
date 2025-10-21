package file

import (
	"io"
	"os"
)

// Copy copies a file from src to dst, preserving its content and permissions.
// If dst does not exist, it will be created. If it does exist, it will be
// overwritten. Returns an error if any file operation fails.
//
// Example usage:
//
//	err := file.Copy("foo.txt", "bar.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	fi, err := in.Stat()
	if err != nil {
		return err
	}
	return os.Chmod(dst, fi.Mode())
}
