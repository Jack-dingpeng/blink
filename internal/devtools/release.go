// Code generated by go-bindata.
// sources:
// devtools/inspector.html
// DO NOT EDIT!

// +build !bdebug

package devtools

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _inspectorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\x31\x6b\x14\x41\x1c\xc5\xfb\xfd\x14\x2f\xdb\x1c\x08\x73\xeb\x89\x85\x24\xbb\x07\xf1\x4c\x21\x88\x08\xa7\x1f\x60\x6f\xe6\xbf\x37\x63\xe6\x66\xd6\x99\xff\x2e\x6e\x97\x26\x95\x45\x10\x22\x96\x6a\x21\xa6\xb7\x88\x48\xc4\x2f\x73\x77\xf1\x63\xc8\xae\x2b\xd8\x58\x0d\x33\xfc\xdf\xef\xff\xde\x9b\xfc\x40\x88\x04\x77\xb0\xf0\x75\x17\xcc\x5a\x33\xee\xdd\x9d\xdd\xc7\x73\x4d\x58\xe8\xe0\x37\xa6\xd9\xe0\xb8\x61\xed\x43\x9c\xe2\xd8\x5a\x0c\x43\x11\x81\x22\x85\x96\xd4\xb4\x17\xbf\x88\x04\x5f\x81\xb5\x89\x88\xbe\x09\x92\x20\xbd\x22\x98\x88\xb5\x6f\x29\x38\x52\x58\x75\x28\xf1\x70\xf9\x48\x44\xee\x2c\xc1\x1a\x49\x2e\x12\x58\x97\x0c\x59\x3a\xac\xa8\x27\x55\xbe\x71\x0a\xc6\x81\x35\xe1\xc9\xe3\xc5\xc9\xd3\xe5\x09\x2a\x63\x69\x9a\x08\x31\x4f\xf2\x03\xe5\x25\x77\x35\x41\xf3\xc6\xce\x93\xfc\xef\x41\xa5\x9a\x27\x00\x90\x6f\x88\x4b\x68\xe6\x5a\xd0\xab\xc6\xb4\x45\x2a\xbd\x63\x72\x2c\x7a\x59\x8a\xf1\x56\xa4\x4c\xaf\x39\xeb\xe5\x47\x90\xba\x0c\x91\xb8\x68\xb8\x12\x0f\xd2\xff\x71\x16\x23\x67\x49\xb2\x09\x86\x3b\xf1\xcc\x5b\x23\xbb\x7f\x90\x7e\xf5\x92\x24\x8b\x18\x24\x26\xce\x3b\x9a\x1c\x21\xca\x60\xea\xf1\x29\x92\xad\x26\x98\x34\x2e\x96\x15\x09\x6a\x4b\x3b\x19\x16\xc4\xc3\x2c\x93\x7d\xd7\x24\x14\xb5\xec\xbd\x8d\xa2\x0a\x03\x54\x4d\xcb\xba\x8e\xb5\xe7\xa9\xf4\x9b\x74\x9e\xe4\xd9\x9f\xa4\xf9\xca\xab\x0e\xd2\x96\x31\x16\x69\xe3\x94\x97\xa7\xa4\x52\x18\x55\xa4\x62\x65\x8d\x3b\xed\x49\x62\x40\xa5\x18\x03\x29\xd3\x62\xe8\xbe\x48\xd1\x87\x17\xa5\x35\x6b\x77\x08\x49\x8e\x29\x8c\xb1\x87\x49\x3d\x9b\xdf\x5e\x7e\xd8\x7e\xff\xb2\xbf\xfa\xb4\xbb\xb9\xd8\x7e\x7b\xb3\x7f\xff\x71\xff\xf5\xdd\xf6\xc7\xcf\xdb\xcb\xab\xdd\xcd\xd9\xee\xe2\xed\xaf\xb3\xf3\xdd\xf5\xe7\xdd\xf9\x75\x9e\xe9\xd9\xb8\x21\x53\xa6\xed\x3d\xf6\xe6\x06\xaf\xc3\xe7\xfc\x0e\x00\x00\xff\xff\xe9\x67\x29\x7c\x62\x02\x00\x00")

func inspectorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_inspectorHtml,
		"inspector.html",
	)
}

func inspectorHtml() (*asset, error) {
	bytes, err := inspectorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "inspector.html", size: 610, mode: os.FileMode(420), modTime: time.Unix(1545296931, 0)}
	a := &asset{bytes: bytes, info: info}
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

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
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
	"inspector.html": inspectorHtml,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"inspector.html": &bintree{inspectorHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
