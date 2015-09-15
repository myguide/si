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

#include <string.h>
#include "file.h"

int is_file_read(const char* className, const char* signature) {
  if ( strcmp( className, "File" ) == 0 ) {
    if ( strcmp( signature, "read(_)" ) == 0 ) {
      return 1;
    }
  }
  return 0;
}

int is_file_write(const char* className, const char* signature) {
  if ( strcmp( className, "File" ) == 0 ) {
    if ( strcmp( signature, "write(_,_,_)" ) == 0 ) {
      return 1;
    }
  }
  return 0;
}
