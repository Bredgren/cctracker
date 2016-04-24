// Code generated by go-bindata.
// sources:
// ../static/css/style.css
// ../tmpl/index.tmpl
// DO NOT EDIT!

package resources

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

var _staticCssStyleCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xca\x4f\xa9\x54\xa8\xe6\xe5\xe2\x4c\x4a\x4c\xce\x4e\x2f\xca\x2f\xcd\x4b\xb1\x52\x28\x4a\x4f\xd2\x30\x32\xd5\x51\x00\x61\x63\x53\x4d\x6b\xa0\x74\x72\x7e\x4e\x7e\x11\x54\xc6\x18\x24\x65\x62\x09\x92\x07\x4b\x2a\x28\xa4\xe5\xe7\x95\xe8\xa6\x25\xe6\x66\xe6\x54\x5a\x29\xe4\xe6\xe7\xe5\x17\x17\x24\x26\xa7\x02\xa5\x6a\x79\xb9\x00\x01\x00\x00\xff\xff\xb5\xf1\xa8\x17\x62\x00\x00\x00")

func staticCssStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_staticCssStyleCss,
		"static/css/style.css",
	)
}

func staticCssStyleCss() (*asset, error) {
	bytes, err := staticCssStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/css/style.css", size: 98, mode: os.FileMode(438), modTime: time.Unix(1461469869, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplIndexTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x92\xb1\x6e\xf3\x20\x14\x85\x67\xff\xd2\xff\x0e\x94\xa9\x1d\x62\xa4\x4c\x1d\x70\x86\xa6\x51\x3b\xb5\x55\x95\xa5\x23\x85\x9b\x18\x85\xe0\x88\x7b\x13\x25\xb2\xfc\xee\x25\xd8\x46\x8e\xda\x89\x03\xe7\xe3\xdc\x83\x6c\x79\xf7\xfc\xbe\x5c\x7f\x7d\xac\x58\x4d\x7b\xb7\xf8\xff\x4f\x0e\x6b\x21\x6b\x50\xe6\x2a\x0a\xb9\x07\x52\x4c\xd7\x2a\x20\x50\xc5\x8f\xb4\x99\x3d\xf2\xde\x21\x4b\x0e\x16\xcb\xc6\x9f\xac\x26\x16\x57\x63\xc9\x36\xde\xfa\x2d\x5b\x07\xa5\x77\x10\xa4\xe8\x99\x84\x3b\xeb\x77\xac\x0e\xb0\xa9\xb8\x40\x52\x64\xb5\xd0\x88\x51\x5e\x1c\x94\x51\x71\x16\xc0\x55\x3c\xed\xb1\x06\xa0\x34\x46\x8a\xb1\x8a\xfc\x6e\xcc\xa5\x8f\x32\xf6\xc4\xac\xa9\xf8\xd5\x82\xd0\xd7\x29\x5e\xd3\x26\xf9\x22\x02\xb7\xa4\x6e\x3c\x81\xa7\x01\x6d\xdb\xa0\xfc\x16\x58\xf9\x12\x9a\xe3\x01\xbb\x2e\x9d\x26\x58\x3b\x85\x58\xf1\xed\xd5\x18\xe8\xdf\xc6\xec\x66\x70\x04\xea\xf9\xa2\x6d\xcb\x37\xb5\x87\xae\x8b\x8d\xe7\xe3\xc5\x5c\xe4\xaf\x10\x67\x91\x72\x44\xae\xb4\x3a\x43\xd0\x16\x61\x6c\x75\x7b\x13\x06\x37\xdf\x2b\xf2\x5c\x16\xd5\x13\x20\x75\x9d\x68\xdb\x7b\xeb\x0d\x9c\xe3\x03\x1b\xe5\x90\xcd\x1f\xca\x4f\x38\x4c\x12\x27\xbd\x62\x00\x78\x33\x5a\x13\x67\x22\x27\x48\x3e\x95\x62\xf8\x1e\xf1\xbd\xfd\x5f\xf3\x13\x00\x00\xff\xff\x75\x7d\x38\x67\x4f\x02\x00\x00")

func tmplIndexTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplIndexTmpl,
		"tmpl/index.tmpl",
	)
}

func tmplIndexTmpl() (*asset, error) {
	bytes, err := tmplIndexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/index.tmpl", size: 591, mode: os.FileMode(438), modTime: time.Unix(1461469699, 0)}
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
	"static/css/style.css": staticCssStyleCss,
	"tmpl/index.tmpl": tmplIndexTmpl,
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
	"static": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"style.css": &bintree{staticCssStyleCss, map[string]*bintree{}},
		}},
	}},
	"tmpl": &bintree{nil, map[string]*bintree{
		"index.tmpl": &bintree{tmplIndexTmpl, map[string]*bintree{}},
	}},
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

