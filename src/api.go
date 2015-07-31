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

var _src_api_strings_wren = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x94\x4f\x6f\xdb\x3c\x0c\xc6\xef\xf9\x14\x3c\x26\x40\x5e\x3b\x7d\x81\x5d\x8a\x6d\x28\x50\x64\x6b\x80\xec\x0f\x90\xdd\x03\xc5\xa6\x6d\x61\x8a\x64\x48\x74\xb2\xb4\xe8\x77\x1f\x6d\xc9\xb1\xe3\xcc\x1b\x52\xa3\x07\x97\x7c\xf8\xe8\x27\x9a\x4c\x1c\x83\x93\x70\x58\x44\x8b\xe8\xee\xbf\x14\x0f\x93\x38\xe6\x3f\x98\x26\x33\x78\x12\xd6\x9e\x60\x2d\x8e\x16\x75\x82\xf0\xff\xe2\xee\x5d\xc8\x3e\x94\x22\xf9\x29\x72\xe4\xd2\xe6\xdf\x03\x5a\x27\x8d\x86\xa1\xcb\x83\xa8\xa8\x30\x76\xe8\xf4\xbe\x20\x2a\xef\xe3\x38\x97\x54\x54\xbb\x28\x31\xfb\xb8\x10\xcf\x3b\xf3\x31\x94\xad\x65\x82\xda\xe1\x3d\x7c\x59\xfd\x08\xa1\x4f\xec\x42\x05\x42\x56\x29\x05\x89\x29\x4f\x56\xe6\x05\x81\xd0\x29\x28\xaf\x06\xa9\x33\x63\xf7\x82\x18\x64\x0e\xa5\x42\xc1\xb1\x83\xc4\x63\x53\xb7\x5e\x3d\x2e\xbf\x6e\x96\xb5\x55\x26\x15\x72\x4c\x10\x1c\x85\x83\x54\x3a\xb2\x72\x57\x11\xa6\x70\x64\x1e\xce\x48\x07\xce\x54\x96\x41\x13\x93\x62\x34\x99\x24\x4a\x38\x07\x1b\xd6\xe9\xdc\xc1\xcb\x04\xf8\xf1\x4c\x28\x73\x0d\x7b\xe4\x4b\xa6\x6c\x85\x99\xd4\x6c\x23\x35\x7c\x36\x8d\x28\x0b\x8a\xc4\x68\x12\x52\xbb\xa9\x9b\x83\xab\x76\x7c\xe2\xec\x22\x5f\x08\xf7\xdd\x72\xf5\xaf\x5a\x50\x36\x6f\x57\x82\x4d\x95\x05\x81\x6b\xde\x2e\x05\x52\xa7\xe8\x93\x58\x5e\x66\x98\x9d\x56\x17\xd9\x96\x7f\xb9\x2f\xe9\x54\xb3\x31\x4f\x95\x90\xb1\x4d\x42\x73\xc7\x5e\xbc\xe6\xf5\x2c\x7d\xec\x44\xbe\x4b\x5b\xd7\x74\x83\xcf\x95\x24\x85\x92\xcf\x78\xae\x9e\xfa\xd4\x2c\x34\xaa\x7e\x5a\xf5\x07\xf0\x2f\x03\xf7\x36\xed\x90\x28\xf8\xf8\xc8\xdb\xad\xf2\xa1\x55\xcf\xc2\x22\x55\x56\xb7\xd2\x81\xc1\x13\xaa\x12\x6d\xf8\xa6\x75\x13\xf9\x7b\x48\x4d\xb5\x43\xbf\xa0\x09\xf6\x3c\x57\xdf\xa2\x26\x34\x0d\xa2\xd9\xb0\x81\x42\x29\x77\x3d\x0f\xdb\xf9\x76\xc6\x23\x68\x31\x21\x75\x02\xa3\x5b\x79\x3d\xb2\xed\x55\x4a\x6b\x98\x88\x4e\x4d\xae\x9b\x24\x3f\x46\xd7\xd7\xaa\xe7\x37\xea\x0e\xf0\x26\x97\x63\x37\x86\xd5\x8d\xe1\xcd\x5c\x5d\x69\x18\xdf\x11\xb0\xde\x11\x2d\x59\x7f\xde\xff\x42\x16\xe6\xff\x2d\x64\xed\xea\xf8\xbd\x19\x27\x6b\x8f\xe8\x7a\x96\xfd\x93\xcc\x2f\xde\xcd\x54\x61\x5f\x79\x1d\x47\x70\x82\xef\x19\xa5\x5d\xeb\x31\x8e\x6e\xcd\x6f\x66\xe9\xfd\x42\x8c\xf3\xf4\xfc\xff\xc0\xf4\x3a\xf9\x1d\x00\x00\xff\xff\xd9\x8b\x91\xe8\x48\x06\x00\x00")

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

	info := bindata_file_info{name: "src/api/strings.wren", size: 1608, mode: os.FileMode(420), modTime: time.Unix(1438380478, 0)}
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

