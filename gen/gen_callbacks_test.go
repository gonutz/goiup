package gen

import (
	"os"
	"path/filepath"
	"testing"
)

// TestGenerateCallbackWrapper is a test case abused to re-generate the callbacks.go
// file containing the C callback wrappers for Go. For each function signature it
// creates the same number of callback wrappers. The number can be increased if
// necessary. genCallbacks takes the number of callbacks per signature as the first
// parameter.
func TestGenerateCallbackWrapper(t *testing.T) {
	path := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "gonutz", "goiup", "gen", "callbacks.go")
	f, err := os.Create(path)
	if err == nil {
		defer f.Close()
		genCallbacks(100, f)
	}
}
