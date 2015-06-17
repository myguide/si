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
#include "register.h"
#include "return_double.h"
#include "wren.h"

#define REGISTER_TEST(name) \
  if (strcmp(testName, #name) == 0) return name##BindForeign(fullName)

const char* testName;


static WrenForeignMethodFn bindForeign(
    WrenVM* vm, const char* module, const char* className,
    bool isStatic, const char* signature)
{
  if (strcmp(module, "main") != 0) return NULL;

  // For convenience, concatenate all of the method qualifiers into a single
  // signature string.
  char fullName[256];
  fullName[0] = '\0';
  if (isStatic) strcat(fullName, "static ");
  strcat(fullName, className);
  strcat(fullName, ".");
  strcat(fullName, signature);

  REGISTER_TEST(return_double);

  exit(1);
  return NULL;
}



/*
void register_classes(WrenVM *vm)
{

    // File
    //wrenDefineMethod(vm, "BaseFile", "Exists(_)", class_file_exists);
    //wrenDefineMethod(vm, "BaseFile", "Read(_)", class_file_read);
    //wrenDefineMethod(vm, "BaseFile", "Write(_,_,_)", class_file_write);

    // Markdown
    //wrenDefineMethod(vm, "BaseMarkdown", "Parse(_)", class_markdown_parse);
}
*/


static void returnFloat(WrenVM* vm)
{
  wrenReturnDouble(vm, 123.456);
}

WrenForeignMethodFn return_doubleBindForeign(const char* signature)
{
  if (strcmp(signature, "static Api.returnFloat") == 0) return returnFloat;

  return NULL;
}
