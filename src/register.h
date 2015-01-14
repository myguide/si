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

#ifndef REGISTER_
#define REGISTER_

#include <wren.h>

// File
extern void class_file_read(WrenVM *vm);
extern void class_file_write(WrenVM *vm);

// Markdown
extern void class_markdown_parse(WrenVM *vm);

void register_classes(WrenVM *vm);

#endif
