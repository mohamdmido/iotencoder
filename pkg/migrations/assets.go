// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sql/20180525115614_create_device_table.down.sql
// sql/20180525115614_create_device_table.up.sql
package migrations

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

var __20180525115614_create_device_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x49\x2d\xcb\x4c\x4e\x2d\x8e\x2f\xc9\x2f\xc8\x4c\x8e\xcf\x4c\xa9\x50\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\xe2\xc2\xab\xa1\xb4\x38\xb5\x28\xbe\x34\x33\x05\x9b\x9e\x10\x47\x27\x1f\x57\x4c\x3d\x18\xca\x22\x03\x50\x54\x65\x16\x17\xe4\x17\x67\x96\x64\xe6\xe7\xc1\x55\x02\x02\x00\x00\xff\xff\xb8\xa0\xd1\x7a\xb4\x00\x00\x00")

func _20180525115614_create_device_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180525115614_create_device_tableDownSql,
		"20180525115614_create_device_table.down.sql",
	)
}

func _20180525115614_create_device_tableDownSql() (*asset, error) {
	bytes, err := _20180525115614_create_device_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180525115614_create_device_table.down.sql", size: 180, mode: os.FileMode(420), modTime: time.Unix(1527249767, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20180525115614_create_device_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x6a\xf3\x30\x10\x84\xef\x7a\x8a\xb9\x25\x81\xbc\x41\x4e\xfa\x93\x0d\xbf\xa8\x2d\xbb\x96\x44\x92\x5e\x8c\x1b\x89\x22\x52\xa2\x20\xcb\xa1\x7d\xfb\x62\x83\xa9\x8b\x43\xe9\x6d\x77\x66\x76\x58\xbe\x6d\x45\x5c\x13\xf4\xa9\x24\x58\xdf\xde\x42\xeb\x93\x0f\x57\x70\x05\x92\x26\xc7\x72\xe1\xaf\x36\x84\xb8\x58\x63\x11\xba\x34\x8c\xab\x0d\x63\xe3\x1d\xff\x97\x11\xc4\x1e\xb2\xd0\xa0\xa3\x50\x5a\xc1\xba\xbb\x3f\xbb\x16\x4b\x06\x78\x0b\x45\x95\xe0\x19\xca\x4a\xe4\xbc\x3a\xe1\x89\x4e\x6b\x06\xbc\xc6\x70\x71\x11\x9a\x8e\x7a\xb8\x95\x26\xcb\x7a\x3d\x85\x9b\x3f\xcf\xe5\x5b\xf4\xf7\x26\xb9\xfa\xe2\x3e\xe7\x66\xd7\xba\x58\x77\xde\xce\x9d\xf7\x70\x7d\xf3\xa9\xb3\x0e\xbb\xc2\xf4\x9f\x96\x15\x6d\x85\x12\x85\xfc\x19\x6b\xd2\x1f\x52\x53\x3c\xd3\x79\x9a\x39\x47\xd7\x24\x67\xeb\x26\x41\x8b\x9c\x94\xe6\x79\x89\x83\xd0\xff\x87\x15\x2f\x85\x24\xec\x68\xcf\x4d\xd6\x3f\x7a\x58\xae\xd8\x04\xa6\x91\xe2\xd9\x10\x84\xdc\xd1\xf1\x31\xd3\x7a\xc0\x53\x7b\xfb\xc1\x80\x42\x7e\xa3\x1e\xf4\x49\xd5\x6f\x1d\x23\xae\x07\x35\xa3\xb5\xda\x7c\x05\x00\x00\xff\xff\x45\xfe\x8e\x12\x18\x02\x00\x00")

func _20180525115614_create_device_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180525115614_create_device_tableUpSql,
		"20180525115614_create_device_table.up.sql",
	)
}

func _20180525115614_create_device_tableUpSql() (*asset, error) {
	bytes, err := _20180525115614_create_device_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180525115614_create_device_table.up.sql", size: 536, mode: os.FileMode(420), modTime: time.Unix(1527249703, 0)}
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
	"20180525115614_create_device_table.down.sql": _20180525115614_create_device_tableDownSql,
	"20180525115614_create_device_table.up.sql": _20180525115614_create_device_tableUpSql,
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
	"20180525115614_create_device_table.down.sql": &bintree{_20180525115614_create_device_tableDownSql, map[string]*bintree{}},
	"20180525115614_create_device_table.up.sql": &bintree{_20180525115614_create_device_tableUpSql, map[string]*bintree{}},
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
