// wrengo v0.0.1-dev
//
// (c) Harry Lawrence 2015
//
// @package wrengo
// @version 0.0.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package module

import (
	"os"
	"io/ioutil"
)

// File represents a file that the user could either
// be reading from or writing to.
type File struct {
	path string
}

// NewFile returns an empty instance of File
func NewFile() File {
	return File{}
}

// Exists checks to see if a file exists, if so
// returns true and assigns the path of File
// with the path found.
func (f *File) Exists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		f.path = filename
		return true
	}
	return false
}

// Read reads a file and returns the contents as a string.
func (f File) Read(filename string) string {
	c, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(c)
}

// Write writes data to a given file, taking into
// consideration the file permissions.
func (f File) Write(filename string, data string, mode int) {
	m := uint32(mode)
	if err := ioutil.WriteFile(filename, []byte(data), os.FileMode(m)); err != nil {
		panic(err)
	}
}
