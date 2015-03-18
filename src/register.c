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

#include "register.h"

void register_classes(WrenVM *vm)
{
    // File
    wrenDefineMethod(vm, "BaseFile", "Exists(_)", class_file_exists);
    wrenDefineMethod(vm, "BaseFile", "Read(_)", class_file_read);
    wrenDefineMethod(vm, "BaseFile", "Write(_,_,_)", class_file_write);

    // Markdown
    wrenDefineMethod(vm, "BaseMarkdown", "Parse(_)", class_markdown_parse);
}
