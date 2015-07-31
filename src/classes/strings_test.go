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

package class

import (
	"testing"
)

func TestContains(t *testing.T) {
	s    := NewStrings()
	text := "hello si"
	res  := s.Contains(text, "ello s")

	if true != res {
		t.Error("Expected true, got ", res)
	}
}

func TestHasPrefix(t *testing.T) {
	s    := NewStrings()
	text := "hello si"
	res  := s.HasPrefix(text, "hel")

	if true != res {
		t.Error("Expected true, got", res)
	}
}

func TestHasSuffix(t *testing.T) {
	s    := NewStrings()
	text := "hello si"
	res  := s.HasSuffix(text, " si")

	if true != res {
		t.Error("Expected true, got", res)
	}
}

func TestIndex(t *testing.T) {
	s    := NewStrings()
	text := "hello si"
	res  := s.Index(text, "o si")

	if 4 != res {
		t.Error("Expected 4, got", res)
	}
}

func TestLastIndex(t *testing.T) {
	s := NewStrings()
	text := "hello si hello again"
	res  := s.LastIndex(text, "ello ")

	if 10 != res {
		t.Error("Expected 10, got", res)
	}
}
