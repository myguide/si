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

// #cgo CFLAGS: -std=c99 -Wall -Werror -I./wren/include
// #cgo LDFLAGS: -L. wren/libwren.a
// #include <wren.h>
// #include "register.h"
// static inline WrenVM* createVM() {
//   WrenConfiguration config = {
//     .reallocateFn = NULL,
//     .heapGrowthPercent = 0,
//     .minHeapSize = 0,
//     .initialHeapSize = 1024 * 1024 * 100
//   };
//   return wrenNewVM(&config);
// }
// 
// static inline int interpret(WrenVM *vm, char *path, char *src) {
//   register_classes(vm);
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

func (m VM) Interpret() {
    c := C.interpret(
        m.Wren,
        C.CString(m.Script.Filename),
        C.CString(string(m.Script.Src)),
    )
    if c == 0 {
        // Success!
    }
}
