package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func src_api_strings_wren() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4a, 0xce,
		0x49, 0x2c, 0x2e, 0x56, 0x08, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x2f, 0x56,
		0xa8, 0xe6, 0x52, 0x00, 0x82, 0xb4, 0xfc, 0xa2, 0xd4, 0xcc, 0xf4, 0x3c,
		0x85, 0xe4, 0xfc, 0xbc, 0x92, 0xc4, 0xcc, 0xbc, 0x62, 0x8d, 0x62, 0x1d,
		0x85, 0xe2, 0xd2, 0xa4, 0xe2, 0x92, 0x22, 0x4d, 0x14, 0xf9, 0x8c, 0xc4,
		0xe2, 0x80, 0xa2, 0xd4, 0xb4, 0xcc, 0x0a, 0x90, 0x82, 0x02, 0x30, 0x0b,
		0x43, 0x41, 0x70, 0x69, 0x1a, 0x54, 0x41, 0x31, 0x98, 0xa5, 0xc9, 0x55,
		0xcb, 0x05, 0x08, 0x00, 0x00, 0xff, 0xff, 0x29, 0x8c, 0x66, 0x4f, 0x74,
		0x00, 0x00, 0x00,
	},
		"src/api/strings.wren",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"src/api/strings.wren": src_api_strings_wren,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
func AssetDir(name string) ([]string, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	pathList := strings.Split(cannonicalName, "/")
	node := _bintree
	for _, p := range pathList {
		node = node.Children[p]
		if node == nil {
			return nil, fmt.Errorf("Asset %s not found", name)
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"src": &_bintree_t{nil, map[string]*_bintree_t{
		"api": &_bintree_t{nil, map[string]*_bintree_t{
			"strings.wren": &_bintree_t{src_api_strings_wren, map[string]*_bintree_t{
			}},
		}},
	}},
}}
