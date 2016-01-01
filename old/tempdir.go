// 7 october 2014
package main

import (
	"os"
	"io"
	"io/ioutil"
	"path/filepath"
)

var tempdir string
var tempfiles []string

// TODO adorn errors
func StartTempDir() (string, error) {
	if tempdir != "" {
		for _, f := range tempfiles {
			err := os.Remove(f)
			if err != nil {
				return "", err
			}
		}
		tempfiles = tempfiles[0:0]			// make len 0 without freeing preexisting space
		return tempdir, nil
	}
	d, err := ioutil.TempDir("", "ohvtemp")
	if err != nil {
		return "", err
	}
	tempdir = d
	tempfiles = nil
	return tempdir, nil
}

// TODO adorn errors
func TempFile(name string, r io.Reader) (string, error) {
	name = filepath.Join(tempdir, name)
	f, err := os.Create(name)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	if err != nil {
		return "", err
	}
	tempfiles = append(tempfiles, name)
	return name, nil
}
