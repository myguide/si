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

// File
extern void class_file_exists(WrenVM *vm);
extern void class_file_read(WrenVM *vm);
extern void class_file_write(WrenVM *vm);

// Markdown
extern void class_markdown_parse(WrenVM *vm);

// Strings
extern void class_strings_contains(WrenVM *vm);

#endif
