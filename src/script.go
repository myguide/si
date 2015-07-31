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

import (
    "fmt"
    "io/ioutil"
)

type Script struct {
    Filename string
    Src      []byte
}

// NewScript creates an instance of the Script object
// to store information on the source to be ran.
func NewScript(filename string) (Script, error) {
    c, err := ioutil.ReadFile(filename)
    if err != nil {
        return Script{}, fmt.Errorf("File '%s' does not exist.", filename)
    }
    s := Script{
        Filename: filename,
        Src: c,
    }
    return s, nil
}


// readApi takes in the account the code written in wren to
// produce a cleaner interface for the end developer
func (s *Script) readApi(paths []string) {
    for _, path := range paths {
        api, err := Asset(path)
        if err != nil {
            panic(err)
        }
        s.Src = append(api, s.Src...)
    }
}
