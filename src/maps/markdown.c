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

#include <string.h>
#include "markdown.h"

int is_markdown_parse(const char* className, const char* signature) {
  if ( strcmp( className, "Markdown" ) == 0 ) {
    if ( strcmp( signature, "Parse(_)" ) == 0 ) {
      return 1;
    }
  }
  return 0;
}
