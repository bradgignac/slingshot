package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// File represents a configuration file on disk.
type File struct {
	path string
	info os.FileInfo
}

// String provides custom formatting for the File struct.
func (f *File) String() string {
	return f.path
}

// Key builds a key for a configuration file.
func (f *File) Key() string {
	ext := filepath.Ext(f.path)
	return strings.TrimSuffix(f.path, ext)
}

// Read reads the contents of the file.
func (f *File) Read() (string, error) {
	body, err := ioutil.ReadFile(f.path)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// FileSet represents a set of configuration files guaranteed to be unique.
type FileSet struct {
	files map[string]*File
}

// Union creates a new FileSet from a slice of FileSets.
func Union(sets []*FileSet) *FileSet {
	totalLen := 0

	for _, set := range sets {
		totalLen += len(set.files)
	}

	union := &FileSet{}
	for _, set := range sets {
		for _, file := range set.files {
			union.Add(file)
		}
	}

	return union
}

// Add adds a new File to the set.
func (s *FileSet) Add(file *File) {
	if s.files == nil {
		s.files = make(map[string]*File)
	}

	k := file.path
	s.files[k] = file
}

// ToSlice converts a FileSet into a slice.
func (s *FileSet) ToSlice() []*File {
	i := 0
	slice := make([]*File, len(s.files))

	for _, file := range s.files {
		slice[i] = file
		i++
	}

	return slice
}
