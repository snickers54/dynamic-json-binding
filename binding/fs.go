package binding

import (
	"os"
	"path/filepath"
	"strings"
)

func listFS(startingPath string) {
	filepath.Walk(startingPath, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			bindingHelper.GoFiles = append(bindingHelper.GoFiles, path)
		}
		return nil
	})
}
