// si v0.0.1-dev
//
// (c) Harry Lawrence 2015
//
// @package si
// @version 0.0.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package class

import (
	"os"
	"io/ioutil"
)

type File struct {

}

func NewFile() File {
	return File{}
}

func (f File) Read(filename string) string {
	c, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(c)
}

func (f File) Write(filename, contents string, perm int) bool {
	err := ioutil.WriteFile(filename, []byte(contents), os.FileMode(perm))
	if err != nil {
		return false
	}
	return true
}
