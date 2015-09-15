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
	"testing"
	"os"
)

func TestRead(t *testing.T) {
	s   := NewFile()
	res := s.Read("../../build/hello.wren")

	if "IO.print(\"Hello, World!\")\n" != res {
		t.Error("Expected true, got ", res)
	}
}

func TestWrite(t *testing.T) {
	s := NewFile()
	res := s.Write("test.txt", "Test Content", 755)

	if _, err := os.Stat("test.txt"); err != nil {
		t.Error("test.txt could not be found, expecting true, got ", res)
	} else {
		os.Remove("test.txt")
	}
}
