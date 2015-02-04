package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _src_api_file_wren = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xce\x49\x2c\x2e\x56\x70\xcb\xcc\x49\x55\xc8\x2c\x56\x70\x4a\x2c\x4e\x05\xb3\xab\xb9\x14\x80\x20\x2f\xb5\x1c\xc8\x02\x33\x6b\x21\x54\x50\x6a\x62\x8a\x46\x1a\x50\x45\x5e\x62\x6e\xaa\x26\x54\x19\x08\x14\xa5\x96\x94\x16\xe5\x29\x14\x97\x16\xa4\x16\xe9\xa1\xaa\x42\xd6\x1f\x5e\x94\x59\x92\x0a\x97\xd2\x51\x48\x49\x2c\x49\xd4\x51\x00\xea\xc9\xc5\x69\x18\x1e\x2d\x50\x93\x6b\xb9\x00\x01\x00\x00\xff\xff\xb3\xbf\x6e\x2c\xc6\x00\x00\x00")

func src_api_file_wren_bytes() ([]byte, error) {
	return bindata_read(
		_src_api_file_wren,
		"src/api/file.wren",
	)
}

func src_api_file_wren() (*asset, error) {
	bytes, err := src_api_file_wren_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "src/api/file.wren", size: 198, mode: os.FileMode(420), modTime: time.Unix(1421797492, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _src_api_main_wren = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xca\x4d\xcc\xcc\xd3\x4b\x4e\xcc\xc9\xe1\x02\x04\x00\x00\xff\xff\x2b\x79\xfa\xfa\x0b\x00\x00\x00")

func src_api_main_wren_bytes() ([]byte, error) {
	return bindata_read(
		_src_api_main_wren,
		"src/api/main.wren",
	)
}

func src_api_main_wren() (*asset, error) {
	bytes, err := src_api_main_wren_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "src/api/main.wren", size: 11, mode: os.FileMode(420), modTime: time.Unix(1421797492, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _src_api_markdown_wren = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8e\x31\xca\xc3\x30\x0c\x85\xf7\x9c\x42\xff\xe6\x2c\x7f\x4e\xe0\xa5\x43\xa1\x43\xa1\xd0\x03\x14\x93\x88\xc6\xd4\x76\x8c\xa4\x90\xa1\xe4\xee\xb5\x83\x93\x26\x54\xd3\x93\xf4\xde\xc7\x6b\x9d\x61\x86\xab\xa1\x57\x37\x4c\x01\x2c\xc3\xc9\x30\x6e\xfb\xbb\x82\x34\x7e\x5b\xe1\xb1\xe9\x79\x79\x45\x43\x8c\x5d\x12\xe9\x55\xf4\x5c\x1d\x42\x5a\xb5\x43\x10\x0c\x52\x1f\xe2\x1a\xca\xf9\x00\xd2\xaa\x17\xef\xea\x1d\x4d\x43\xbe\xac\xd0\x80\x53\xea\xb4\xc8\x72\xb9\x65\xdb\x5d\xc8\x86\xa7\x5a\xe1\x75\xe9\x9d\x87\x50\x46\x0a\xc0\x63\x44\xfa\x5f\xcc\x5f\xdb\x0f\xe7\x6c\x1d\xaa\x68\xa4\xdf\x13\x9a\x06\x2e\x3e\x3a\xf4\xb9\xac\xf4\x96\xff\x4a\x6e\xae\x3e\x01\x00\x00\xff\xff\x82\xe6\x00\x27\x3e\x01\x00\x00")

func src_api_markdown_wren_bytes() ([]byte, error) {
	return bindata_read(
		_src_api_markdown_wren,
		"src/api/markdown.wren",
	)
}

func src_api_markdown_wren() (*asset, error) {
	bytes, err := src_api_markdown_wren_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "src/api/markdown.wren", size: 318, mode: os.FileMode(420), modTime: time.Unix(1421797492, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"src/api/file.wren": src_api_file_wren,
	"src/api/main.wren": src_api_main_wren,
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
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"src": &_bintree_t{nil, map[string]*_bintree_t{
		"api": &_bintree_t{nil, map[string]*_bintree_t{
			"file.wren": &_bintree_t{src_api_file_wren, map[string]*_bintree_t{
			}},
			"main.wren": &_bintree_t{src_api_main_wren, map[string]*_bintree_t{
			}},
			"markdown.wren": &_bintree_t{src_api_markdown_wren, map[string]*_bintree_t{
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

