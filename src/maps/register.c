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
#include "register.h"
#include "strings.h"

WrenForeignMethodFn findForeignMethods( WrenVM* vm, const char* module,
  const char* className, bool isStatic, const char* signature )
{
  // Strings
  if (is_strings_contains(className, signature) == 1) { return class_strings_contains; }
  if (is_strings_has_prefix(className, signature) == 1) { return class_strings_has_prefix; }
  if (is_strings_has_suffix(className, signature) == 1) { return class_strings_has_suffix; }
  if (is_strings_index(className, signature) == 1) { return class_strings_index; }
  if (is_strings_last_index(className, signature) == 1) { return class_strings_last_index; }

  return NULL;
}
