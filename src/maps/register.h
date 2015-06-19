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

#ifndef register_h
#define register_h

#include <wren.h>

WrenForeignMethodFn findForeignMethods( WrenVM* vm, const char* module,
  const char* className, bool isStatic, const char* signature );

// Strings
extern void class_strings_contains(WrenVM *vm);
extern void class_strings_has_prefix(WrenVM *vm);
extern void class_strings_has_suffix(WrenVM *vm);
extern void class_strings_index(WrenVM *vm);

#endif
