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

package main
// #cgo CFLAGS: -std=c99 -Wall -Werror -I./wren/src/include
// #cgo LDFLAGS: -L. wren/lib/libwren.a
//
// #include <wren.h>
// #include "register.h"
//
// static inline WrenVM* createVM() {
//   WrenConfiguration config = {
//     .reallocateFn = NULL,
//     .heapGrowthPercent = 0,
//     .minHeapSize = 0,
//     .initialHeapSize = 1024 * 1024 * 100
//   };
//   config.bindForeignMethodFn = findForeignMethods;
//   return wrenNewVM(&config);
// }
// 
// static inline int interpret(WrenVM *vm, char *path, char *src) {
//   int c = 0;
//   switch(wrenInterpret(vm, path, src)) {
//     case WREN_RESULT_SUCCESS: c = 0; break;
//     case WREN_RESULT_COMPILE_ERROR: c = 1; break;
//     case WREN_RESULT_RUNTIME_ERROR: c = 2; break;
//     default: break;
//   }
//   wrenFreeVM(vm);
//   return c;
// }
import "C"

type VM struct {
    Wren   *C.WrenVM
    Script Script
}

// NewVM creates an instance of the VM object.
func NewVM() VM {
    return VM{
        Wren: C.createVM(),
    }
}

// Interpret interprets the Wren script
func (m VM) interpret() {
    c := C.interpret(
        m.Wren,
        C.CString(m.Script.Filename),
        C.CString(string(m.Script.Src)),
    )
    if c == 0 {
        // Success!
    }
}

// wrenGetArgumentDouble is the Go binding for the
// C implementation of wrenGetArgumentDouble
func wrenGetArgumentDouble(vm *C.WrenVM, index int) int {
    return int(C.wrenGetArgumentDouble(vm, C.int(index)))
}

// wrenGetArgumentString is the Go binding for the
// C implementation of wrenGetArgumentString
func wrenGetArgumentString(vm *C.WrenVM, index int) string {
    return C.GoString(C.wrenGetArgumentString(vm, C.int(index)))
}

// wrenGetArgumentBool is the Go binding for the
// C implementation of wrenGetArgumentBool
func wrenGetArgumentBool(vm *C.WrenVM, index int) bool {
    return bool(bool(C.wrenGetArgumentBool(vm, C.int(index))))
}

// wrenReturnDouble is the Go binding for the
// C implementation of wrenReturnDouble
func wrenReturnDouble(vm *C.WrenVM, value float64) {
    C.wrenReturnDouble(vm, C.double(value))
}

// wrenReturnString is the Go binding for the
// C implementation of wrenReturnString
func wrenReturnString(vm *C.WrenVM, value string) {
    C.wrenReturnString(vm, C.CString(value), -1)
}

// wrenReturnBool is the Go binding for the
// C implementation of wrenReturnBool
func wrenReturnBool(vm *C.WrenVM, value bool) {
    C.wrenReturnBool(vm, C._Bool(value))
}
