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
	"strings"
)

type Strings struct {

}

func NewStrings() Strings {
	return Strings{}
}

func (s Strings) Contains(st, substr string) bool {
	return strings.Contains(st, substr)
}

func (s Strings) HasPrefix(st, prefix string) bool {
	return strings.HasPrefix(st, prefix)
}

func (s Strings) HasSuffix(st, suffix string) bool {
	return strings.HasSuffix(st, suffix)
}

func (s Strings) Index(st, sep string) int {
	return strings.Index(st, sep)
}


