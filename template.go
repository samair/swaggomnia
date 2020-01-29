// Code generated by go-bindata. DO NOT EDIT.
// sources:
// tmpl/swagger.yaml

package main


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataTmplSwaggeryaml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x52\x4d\x8f\xd3\x30\x10\xbd\xe7\x57\x8c\xaa\x3d\x6e\x43\x05\xb7\x48" +
	"\x5c\x16\xd0\x2e\x12\x8b\x2a\x54\x71\xb7\x92\x49\x63\x6d\xfd\xc1\x8c\x53\x88\x5a\xff\x77\x64\x27\x8d\xf3\x01\xec" +
	"\x91\x9e\x9a\xe7\x37\xef\xcd\x9b\x19\x63\x51\x0b\x2b\x0b\x78\x97\xef\xf2\x5d\x26\x75\x6d\x8a\x0c\xa0\x42\x2e\x49" +
	"\x5a\x27\x8d\x2e\x60\x73\xb9\x40\xfe\xc1\xe8\x5a\x1e\xf3\x8f\xe9\x01\xbc\xdf\x64\x00\x67\x24\x5e\xd1\xbe\xf7\xe0" +
	"\x40\x71\xd2\x9d\x70\x4e\x38\x04\x28\x3e\x33\x52\x90\x08\xae\x5b\x68\xe9\x34\xe7\x3d\x19\x76\xe0\xfd\x9b\x09\xf4" +
	"\x20\x18\xf7\xc2\x35\xb1\xda\x89\x23\x17\xd9\xe5\x02\x24\xf4\x11\xe1\xee\x05\xbb\x7b\xb8\x3b\x8b\x53\x8b\x50\xbc" +
	"\x87\xfc\x93\x76\xd2\x49\x64\xf0\x3e\xdb\x82\x16\x6a\xe8\x23\x10\xe1\x0a\x8f\xe8\x1e\xc9\xb4\xf6\xab\x50\x38\x34" +
	"\x3b\x8f\xbe\x09\xda\xa8\xab\x50\x6f\x85\x6b\xa2\xd9\x76\xee\x16\xf1\x95\xdb\x84\x16\x08\xf7\x70\x47\xc8\xa6\xa5" +
	"\x12\x23\x77\xa8\xf2\x3e\x03\x08\xfd\x84\x4f\xb8\xc2\x37\x54\xe6\x1c\xe3\xed\x09\x6b\xf9\xcb\xfb\x30\x98\x40\x49" +
	"\xa6\x2f\x13\xad\x28\x95\x84\xa3\x5c\x2f\x78\x03\xf3\x67\x74\x8d\xa9\xe0\x0a\x07\xf3\xc5\xfc\x44\x82\x9b\x26\x40" +
	"\x3f\xbd\xfe\xff\xf6\x95\xb9\x84\x1f\xb7\x4a\x09\xea\x86\x11\x8e\x0e\x0b\x56\xe8\x55\xd6\x60\x68\x42\xf9\xac\xd9" +
	"\x28\x2d\xc5\x5e\x90\x50\x3c\x79\x78\x30\x55\x97\xaf\xd0\x27\x14\x15\xd2\x98\x67\xd4\x5c\xd4\x3d\x4b\x85\x87\xce" +
	"\x62\xe2\xad\x0e\x37\x55\xac\x4f\xf7\x26\x3c\xec\xb7\x07\x6c\xe8\x05\xdd\x70\x92\x8b\xd9\xdb\xb8\x6e\x12\x6a\x36" +
	"\xf7\x45\xba\x51\x6a\x76\x70\x7d\x59\x32\x06\x90\xa1\xc5\xb0\xf7\x04\x71\xd9\xa0\x12\xc5\xf8\x0d\xe0\x3a\x1b\x14" +
	"\xd8\x91\xd4\xc7\x44\x5c\x5e\xe9\x0d\x27\xfc\xd1\x4a\xc2\xaa\x00\x47\x2d\xfe\x2d\xe3\xeb\x89\x86\x05\xfc\x23\xca" +
	"\x72\xeb\x43\x9e\x26\x16\xfe\x97\x44\x33\x80\x90\xad\xd1\x8c\x9c\xac\xdf\xee\x76\xd3\x3e\xe6\x86\xdc\x96\x25\x32" +
	"\xd7\xed\x09\x8c\x45\x12\x01\xde\x64\x0b\xe9\x3f\xff\xfd\x1d\x00\x00\xff\xff\x79\x9c\xca\x95\x44\x05\x00\x00")

func bindataTmplSwaggeryamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTmplSwaggeryaml,
		"tmpl/swagger.yaml",
	)
}



func bindataTmplSwaggeryaml() (*asset, error) {
	bytes, err := bindataTmplSwaggeryamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "tmpl/swagger.yaml",
		size: 1348,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1580332212, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"tmpl/swagger.yaml": bindataTmplSwaggeryaml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"tmpl": {Func: nil, Children: map[string]*bintree{
		"swagger.yaml": {Func: bindataTmplSwaggeryaml, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
