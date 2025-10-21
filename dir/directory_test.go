package dir

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	tempDir := t.TempDir()

	// Existing directory should return true
	assert.True(t, Exists(tempDir), "Expected directory to exist")

	// Non-existent directory should return false
	nonExistent := tempDir + "/does-not-exist"
	assert.False(t, Exists(nonExistent), "Expected non-existent directory to return false")

	// File (not a directory) should return false
	tempFile := tempDir + "/file"
	f, err := os.Create(tempFile)
	assert.NoError(t, err, "Failed to create temp file")
	f.Close()
	assert.False(t, Exists(tempFile), "Expected file to not be recognized as a directory")
}
