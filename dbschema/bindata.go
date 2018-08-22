// Code generated by go-bindata.
// sources:
// migrations/001_base.sql
// migrations/002_add_buzz_type.sql
// migrations/003_add_created_at_updated_at.sql
// DO NOT EDIT!

package dbschema

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

var _migrations001_baseSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x4d\x4b\x85\x40\x18\x85\xf7\xef\xaf\x38\x3b\xef\xa5\xee\x26\xb8\x2b\x57\xa3\xbe\xd2\xd4\xa4\x32\x33\x86\xae\xc2\x72\x8a\x81\xfc\x40\x8d\xc0\xe8\xbf\x47\x46\x1f\xd0\x3d\xdb\xf3\x70\x0e\xcf\xe1\x80\xb3\xce\x3f\x4d\xcd\xe2\x50\x8e\x14\x6b\x16\x96\x61\x45\xa4\x18\x32\x45\x96\x5b\x70\x25\x8d\x35\xb8\x7f\x59\xd7\xbb\x47\xe7\xda\x19\x3b\x02\x7c\x8b\x9f\x18\xd6\x52\x28\x14\x5a\xde\x08\x5d\xe3\x9a\xeb\x73\x02\xfa\xa6\x73\xdf\xc4\xad\xd0\xf1\xa5\xd0\xbb\x8b\xe3\x71\xff\xd9\xb5\x6e\x7e\x98\xfc\xb8\xf8\xa1\x87\xe5\xca\x62\x7b\xca\x4a\xa5\x90\x70\x2a\x4a\x65\x11\x04\x5f\xe0\xd2\xf8\xe7\x79\x1b\xb9\x32\x79\x16\x9d\x00\xdf\xde\x03\xda\x87\x44\x7f\x5d\x92\xe1\xb5\xa7\x44\xe7\xc5\xaf\xcb\x3f\x8f\x90\x3e\x02\x00\x00\xff\xff\x0e\xa4\x82\x74\xfe\x00\x00\x00")

func migrations001_baseSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations001_baseSql,
		"migrations/001_base.sql",
	)
}

func migrations001_baseSql() (*asset, error) {
	bytes, err := migrations001_baseSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/001_base.sql", size: 254, mode: os.FileMode(420), modTime: time.Unix(1534923158, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations002_add_buzz_typeSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\x0e\x72\x75\x0c\x71\x55\x08\x89\x0c\x70\x55\x28\x49\x2d\x2e\x89\x2f\xa9\x2c\x48\x55\x70\x0c\x56\x70\xf5\x0b\xf5\x55\xd0\x50\x07\x89\x39\xaa\xeb\x28\x80\x19\x4e\x30\x86\x33\x8c\xe1\xa2\xae\x69\xcd\xe5\xe8\x13\xe2\x1a\xa4\x10\xe2\xe8\xe4\xe3\xaa\x90\x54\x5a\x55\x15\x9f\x96\x9a\x9a\x52\xac\xe0\xe8\xe2\xa2\xe0\xec\xef\x13\xea\xeb\x07\x11\x05\x9b\x0c\xb7\xc3\x9a\x8b\x0b\xd9\x29\x2e\xf9\xe5\x79\x5c\x2e\x41\xfe\x01\x68\x4e\xb1\x06\x04\x00\x00\xff\xff\xd7\xea\x00\xc7\xb0\x00\x00\x00")

func migrations002_add_buzz_typeSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations002_add_buzz_typeSql,
		"migrations/002_add_buzz_type.sql",
	)
}

func migrations002_add_buzz_typeSql() (*asset, error) {
	bytes, err := migrations002_add_buzz_typeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/002_add_buzz_type.sql", size: 176, mode: os.FileMode(420), modTime: time.Unix(1534930195, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations003_add_created_at_updated_atSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2a\xad\xaa\x8a\x4f\x4b\x4d\x4d\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2e\x4a\x4d\x2c\x49\x4d\x89\x4f\x2c\x51\xc8\xcc\x2b\x49\x4d\x4f\x2d\xb2\x26\x42\x57\x69\x41\x0a\x86\x2e\x2e\x64\xbb\x5d\xf2\xcb\xf3\x00\x01\x00\x00\xff\xff\x03\x17\x35\x92\x8c\x00\x00\x00")

func migrations003_add_created_at_updated_atSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations003_add_created_at_updated_atSql,
		"migrations/003_add_created_at_updated_at.sql",
	)
}

func migrations003_add_created_at_updated_atSql() (*asset, error) {
	bytes, err := migrations003_add_created_at_updated_atSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/003_add_created_at_updated_at.sql", size: 140, mode: os.FileMode(420), modTime: time.Unix(1534928685, 0)}
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
	"migrations/001_base.sql":                      migrations001_baseSql,
	"migrations/002_add_buzz_type.sql":             migrations002_add_buzz_typeSql,
	"migrations/003_add_created_at_updated_at.sql": migrations003_add_created_at_updated_atSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"001_base.sql":                      &bintree{migrations001_baseSql, map[string]*bintree{}},
		"002_add_buzz_type.sql":             &bintree{migrations002_add_buzz_typeSql, map[string]*bintree{}},
		"003_add_created_at_updated_at.sql": &bintree{migrations003_add_created_at_updated_atSql, map[string]*bintree{}},
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
