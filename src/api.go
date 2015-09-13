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

var _src_api_strings_wren = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x94\xcf\x6e\xdb\x30\x0c\xc6\xef\x7e\x0a\x1e\x63\x20\xb3\xd3\x01\xbb\x14\xdb\x50\xa0\xe8\xd6\x00\xd9\x1f\x20\xbb\x07\x8a\x4c\xdb\xc2\x14\xc9\x90\xe4\x64\xe9\xd0\x77\x1f\x2d\x4b\x8b\xe3\xd4\x1b\x52\xa3\x07\x97\xfc\xf8\xe9\x27\x9a\x4c\x9e\x83\x15\xb0\x5f\x64\x8b\xec\xe6\x4d\x81\xfb\x24\xcf\xe9\x0f\x66\x3c\x85\x47\x66\xcc\x11\x56\xec\x60\x50\x71\x84\xb7\x8b\x9b\x77\x21\x7b\xd7\x30\xfe\x93\x55\x48\xa5\xfe\xdf\x3d\x1a\x2b\xb4\x82\xb1\xcb\x1d\x6b\x5d\xad\xcd\xd8\xe9\x7d\xed\x5c\x73\x9b\xe7\x95\x70\x75\xbb\xcd\xb8\xde\xe5\x35\x7b\xda\xea\x8f\xa1\x6c\x25\x38\x2a\x8b\xb7\xf0\x65\xf9\x23\x84\x3e\x91\x8b\xab\x11\xca\x56\x4a\xe0\xba\x39\x1a\x51\xd5\x0e\x98\x2a\x40\xf6\x6a\x10\xaa\xd4\x66\xc7\x1c\x81\xcc\xa1\x91\xc8\x28\xb6\x17\x78\xf0\x75\xab\xe5\xfd\xc3\xd7\xf5\x43\x67\x55\x0a\x89\x14\x63\x0e\x0e\xcc\x42\x21\xac\x33\x62\xdb\x3a\x2c\xe0\x40\x3c\x94\x11\x16\xac\x6e\x0d\x81\x72\x5d\x60\x96\x24\x5c\x32\x6b\x61\x4d\x3a\x55\x59\xf8\x9d\x00\x3d\x3d\x13\x8a\x4a\xc1\x0e\xe9\x92\x05\x59\x61\x29\x14\xd9\x08\x05\x9f\xb5\x17\x95\x41\xc1\xb5\x72\x4c\x28\x3b\xb3\x73\xb0\xed\x96\x4e\x4c\xcf\xf2\x35\xb3\xdf\x0d\x55\xff\xea\x04\x8d\x7f\xbb\x10\xac\xdb\x32\x08\xac\x7f\x3b\x17\x08\x55\x60\x9f\xc4\xe6\x3c\x43\xec\x6e\x79\x96\xf5\x69\x42\x22\x8c\x96\x3b\x50\x78\x98\xa5\x74\x2b\x1f\x7e\x4e\xe2\xed\xee\xa3\x80\x3a\xef\x1b\xb3\xb1\xbe\x01\x74\x94\x70\x82\x49\xf1\x84\xe6\x05\xa7\x5e\x94\x86\x2e\x75\x4f\xac\xfb\x00\xfd\xcb\xe8\x9c\x98\xb6\xe8\x5c\x70\xec\x23\xaf\xb7\xaa\xc6\x56\x03\x0b\x83\xae\x35\x2a\x4a\x47\x06\x8f\x28\x1b\x34\xe1\x83\x76\x1d\xa4\x8f\x21\x94\xeb\x1c\x86\x05\x3e\x38\xf0\x5c\x7e\xcb\x7c\x68\x16\x44\xe9\xb8\x95\x4c\x4a\x7b\x39\x0c\x9b\xf9\x26\xa5\xf9\x33\xc8\x9d\x3c\x82\x56\x51\xde\xcd\x6b\xbc\x4a\x63\x34\x11\xb9\x63\xec\x74\x18\xa3\x7e\x86\x2e\xaf\xd5\x0d\x6f\x76\x3a\xa0\x37\x39\x9f\xb9\x29\xac\xd3\x0c\x5e\xcd\x75\x2a\x0d\xb3\x3b\x01\x36\x38\x22\x92\x0d\x87\xfd\x1f\x64\x61\xf8\x5f\x43\x16\xf7\xa6\x5f\x9a\x69\xb2\x78\xc4\xa9\x67\xe5\x7f\xc9\xfa\xad\xbb\x9a\x2a\x2c\x2b\xed\xe2\x04\x4e\xf0\xfd\x8b\x12\x77\x7a\x8a\xe3\xb4\xe3\x57\xb3\x0c\x7e\x1e\xa6\x79\x06\xfe\x2f\x30\x3d\x27\x7f\x02\x00\x00\xff\xff\xfc\xfd\x1f\xdd\x45\x06\x00\x00")

func src_api_strings_wren_bytes() ([]byte, error) {
	return bindata_read(
		_src_api_strings_wren,
		"src/api/strings.wren",
	)
}

func src_api_strings_wren() (*asset, error) {
	bytes, err := src_api_strings_wren_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "src/api/strings.wren", size: 1605, mode: os.FileMode(420), modTime: time.Unix(1442159907, 0)}
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
			"strings.wren": &_bintree_t{src_api_strings_wren, map[string]*_bintree_t{
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

