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

func src_api_file_wren() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4a, 0xce,
		0x49, 0x2c, 0x2e, 0x56, 0x70, 0xcb, 0xcc, 0x49, 0x55, 0xa8, 0xe6, 0x52,
		0x00, 0x82, 0xb4, 0xfc, 0xa2, 0xd4, 0xcc, 0xf4, 0x3c, 0x85, 0xa0, 0xd4,
		0xc4, 0x14, 0x8d, 0x34, 0xa0, 0x78, 0x5e, 0x62, 0x6e, 0xaa, 0x26, 0x8a,
		0x54, 0x78, 0x51, 0x66, 0x49, 0x2a, 0x5c, 0x4e, 0x47, 0x21, 0x25, 0xb1,
		0x24, 0x51, 0x47, 0xa1, 0x20, 0xb5, 0x28, 0x17, 0x55, 0x9d, 0x6b, 0x45,
		0x66, 0x71, 0x49, 0x31, 0x92, 0x21, 0xb5, 0x5c, 0x80, 0x00, 0x00, 0x00,
		0xff, 0xff, 0x5d, 0x75, 0xd5, 0x13, 0x6f, 0x00, 0x00, 0x00,
	},
		"src/api/file.wren",
	)
}

func src_api_markdown_wren() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4a, 0xce,
		0x49, 0x2c, 0x2e, 0x56, 0xf0, 0x4d, 0x2c, 0xca, 0x4e, 0xc9, 0x2f, 0xcf,
		0x53, 0xa8, 0xe6, 0x52, 0x00, 0x82, 0xb4, 0xfc, 0xa2, 0xd4, 0xcc, 0xf4,
		0x3c, 0x85, 0x80, 0xc4, 0xa2, 0xe2, 0x54, 0x8d, 0x5c, 0xa8, 0xa4, 0x26,
		0x57, 0x2d, 0x17, 0x20, 0x00, 0x00, 0xff, 0xff, 0x56, 0x31, 0x83, 0x07,
		0x2f, 0x00, 0x00, 0x00,
	},
		"src/api/markdown.wren",
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
	"src/api/file.wren": src_api_file_wren,
	"src/api/markdown.wren": src_api_markdown_wren,
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
			"file.wren": &_bintree_t{src_api_file_wren, map[string]*_bintree_t{
			}},
			"markdown.wren": &_bintree_t{src_api_markdown_wren, map[string]*_bintree_t{
			}},
		}},
	}},
}}
