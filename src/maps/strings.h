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

#ifndef map_strings_h
#define map_strings_h

int is_strings_contains(const char* className, const char* signature);
int is_strings_has_prefix(const char* className, const char* signature);
int is_strings_has_suffix(const char* className, const char* signature);
int is_strings_index(const char* className, const char* signature);
int is_strings_last_index(const char* className, const char* signature);

#endif
