package config

import (
	"os"
	"path/filepath"
)

// FindFiles returns a FileSets containing all supported config files in the
// provided slice of paths.
func FindFiles(paths []string) (*FileSet, error) {
	sets := make([]*FileSet, len(paths))

	for i, path := range paths {
		setForPath, err := findFilesInPath(path)
		if err != nil {
			return nil, err
		}

		sets[i] = setForPath
	}

	return Union(sets), nil
}

func findFilesInPath(path string) (*FileSet, error) {
	set := &FileSet{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if isConfigFile(path, info) {
			file := &File{path: path, info: info}
			set.Add(file)
		}

		return nil
	})

	return set, err
}

func isConfigFile(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}

	ext := filepath.Ext(path)
	switch ext {
	case ".json", ".yaml", ".txt":
		return true
	}

	return false
}
