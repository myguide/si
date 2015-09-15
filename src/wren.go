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

package main

// #cgo CFLAGS: -std=c99 -Wall -Werror -I./wren/src/include
// #cgo LDFLAGS: -L. wren/lib/libwren.a
// #include <wren.h>
import "C"

import (
    "os"
    "fmt"
    "./classes"
    "github.com/myguide/cli"
)

func main() {
    app := cli.NewApp()

    // constants located in meta.go
    app.Name    = appName
    app.Version = version
    app.Usage   = usage
    app.Author  = author
    app.Email   = email

    app.Commands = []cli.Command {
        {
            Name: "run",
            Usage: "your wren file!",
            Action: func (c *cli.Context) {
                s, err := NewScript(os.Args[2]);
                if err != nil {
                    fmt.Println(err); os.Exit(1)
                }
                vm := NewVM()
                vm.Script = s
                vm.Script.readApi([]string{
                    "src/api/testing.wren",
                    "src/api/strings.wren",
                    "src/api/file.wren",
                })
                vm.interpret()
            },
        },
    }

    app.Run(os.Args)
}

// Strings
//export class_strings_contains
func class_strings_contains(vm *C.WrenVM) {
    obj    := class.NewStrings()
    s      := wrenGetArgumentString(vm, 1)
    substr := wrenGetArgumentString(vm, 2)
    wrenReturnBool(vm, obj.Contains(s, substr))
}

//export class_strings_has_prefix
func class_strings_has_prefix(vm *C.WrenVM) {
    obj    := class.NewStrings()
    s      := wrenGetArgumentString(vm, 1)
    prefix := wrenGetArgumentString(vm, 2)
    wrenReturnBool(vm, obj.HasPrefix(s, prefix))
}

//export class_strings_has_suffix
func class_strings_has_suffix(vm *C.WrenVM) {
    obj    := class.NewStrings()
    s      := wrenGetArgumentString(vm, 1)
    suffix := wrenGetArgumentString(vm, 2)
    wrenReturnBool(vm, obj.HasSuffix(s, suffix))
}

//export class_strings_index
func class_strings_index(vm *C.WrenVM) {
    obj := class.NewStrings()
    s   := wrenGetArgumentString(vm, 1)
    sep := wrenGetArgumentString(vm, 2)
    wrenReturnDouble(vm, float64(obj.Index(s, sep)))
}

//export class_strings_last_index
func class_strings_last_index(vm *C.WrenVM) {
    obj := class.NewStrings()
    s   := wrenGetArgumentString(vm, 1)
    sep := wrenGetArgumentString(vm, 2)
    wrenReturnDouble(vm, float64(obj.LastIndex(s, sep)))
}

// File
//export class_file_read
func class_file_read(vm *C.WrenVM) {
    obj := class.NewFile()
    f   := wrenGetArgumentString(vm, 1)
    wrenReturnString(vm, obj.Read(f))
}

//export class_file_write
func class_file_write(vm *C.WrenVM) {
    obj := class.NewFile()
    f   := wrenGetArgumentString(vm, 1)
    c   := wrenGetArgumentString(vm, 2)
    p   := wrenGetArgumentDouble(vm, 3)
    wrenReturnBool(vm, obj.Write(f, c, p))
}
