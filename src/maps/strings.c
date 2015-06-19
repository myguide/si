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
#include "strings.h"

int is_strings_contains(const char* className, const char* signature) {
  if ( strcmp( className, "Strings" ) == 0 ) {
    if ( strcmp( signature, "contains(_,_)" ) == 0 ) {
      return 1;
    }
  }
  return 0;
}

int is_strings_has_prefix(const char* className, const char* signature) {
  if ( strcmp( className, "Strings" ) == 0 ) {
    if ( strcmp( signature, "hasPrefix(_,_)" ) == 0 ) {
      return 1;
    }
  }
  return 0;
}

int is_strings_has_suffix(const char* className, const char* signature) {
  if ( strcmp( className, "Strings" ) == 0 ) {
    if ( strcmp( signature, "hasSuffix(_,_)" ) == 0 ) {
      return 1;
    }
  }
  return 0;
}

int is_strings_index(const char* className, const char* signature) {
  if ( strcmp( className, "Strings" ) == 0 ) {
    if ( strcmp( signature, "index(_,_)" ) == 0 ) {
      return 1;
    }
  }
  return 0;
}

int is_strings_last_index(const char* className, const char* signature) {
  if ( strcmp( className, "Strings" ) == 0 ) {
    if ( strcmp( signature, "lastIndex(_,_)" ) == 0 ) {
      return 1;
    }
  }
  return 0;
}
