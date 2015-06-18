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
	"github.com/hazbo/blackfriday"
)

// Markdown represents the markdown file / content that
// the user will be able to parse.
type Markdown struct {

}

// NewMarkdown returns a new empty instance of
// Markdown
func NewMarkdown() Markdown {
	return Markdown{}
}

// Parse parses the markdown string passed through
// and returns the contents as a string
func (m Markdown) Parse(contents string) string {
	return string(blackfriday.MarkdownCommon([]byte(contents)))
}
