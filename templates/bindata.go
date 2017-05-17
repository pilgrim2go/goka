// Code generated by go-bindata.
// sources:
// templates/common/base.go.html
// templates/common/head.go.html
// templates/common/menu.go.html
// templates/monitor/index.go.html
// templates/monitor/menu.go.html
// templates/monitor/processor.go.html
// templates/query/index.go.html
// DO NOT EDIT!

package templates

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

var _templatesCommonBaseGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4a\x2c\x4e\x55\xaa\xad\xe5\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x01\xaa\xc9\x48\x4d\x4c\x51\x52\xd0\x03\xa9\x82\x48\xda\x24\xe5\xa7\x54\xa2\xaa\xc9\x4d\xcd\x2b\x85\xa8\x41\x16\x4d\xce\xcf\x2b\x49\xcd\x2b\x81\x6a\xd6\x87\x68\xb3\xd1\x87\x59\x91\x9a\x97\x02\x14\x07\x04\x00\x00\xff\xff\xfd\x8f\xc0\x67\x8d\x00\x00\x00")

func templatesCommonBaseGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonBaseGoHtml,
		"templates/common/base.go.html",
	)
}

func templatesCommonBaseGoHtml() (*asset, error) {
	bytes, err := templatesCommonBaseGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common/base.go.html", size: 141, mode: os.FileMode(420), modTime: time.Unix(1490699267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCommonHeadGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\xcb\x72\xda\x30\x14\xdd\xe7\x2b\x3c\xde\xb0\x48\x6d\xf1\x6a\x20\x19\x9c\x4e\x4a\x5a\x1e\x79\x11\x48\x20\x74\xd3\x11\xd6\xb5\x2d\x47\x96\x1c\x49\x06\x5c\xca\xbf\x57\x0e\x25\x61\xd2\xf4\x91\x45\x59\xe9\x1e\xee\x3d\xf7\x9c\x63\x8d\x56\x2b\x02\x01\xe5\x60\xd9\x11\x60\x62\xaf\xd7\x7b\xad\xe2\x70\xbc\x67\x99\x5f\x2b\x01\x8d\x2d\x3f\xc2\x52\x81\xf6\xec\x4c\x07\x4e\xd3\xde\xfd\x2b\xd2\x3a\x75\xe0\x21\xa3\x73\xcf\x5e\x3a\x19\x76\x7c\x91\xa4\x58\xd3\x19\x03\xdb\xf2\x05\xd7\xc0\xcd\x1c\x05\x0f\x48\x08\xdb\x49\x4d\x35\x83\xe3\xd5\xca\x4d\x71\x08\x5f\x1f\xab\xf5\xba\x85\x36\xf0\x0e\x39\xc7\x09\x78\x36\x01\xe5\x4b\x9a\x6a\x2a\xf8\x0e\xa5\xfd\x6b\xe3\x9c\xc2\x22\x15\x52\xef\x74\x2d\x28\xd1\x91\x47\x60\x4e\x7d\x70\x1e\x8b\x77\x16\xe5\x54\x53\xcc\x1c\xe5\x63\x06\x5e\x65\x4b\xc4\x28\xbf\xb7\x22\x09\x81\x57\x2a\x4c\xa9\x23\x84\x02\x43\xa3\xdc\x50\x88\x90\x01\x4e\xa9\x72\x8d\x39\xe4\x2b\xf5\x21\xc0\x09\x65\xb9\x77\x95\x02\xdf\x1f\x61\xae\xf6\xdb\x82\x13\xe0\x0a\xc8\x51\xad\x5c\xfe\xfe\x84\x97\x2c\x09\xcc\x2b\x29\x9d\x33\x50\x11\x80\x2e\x59\x3a\x4f\xc1\x2b\x69\x58\xea\x82\xa9\xb4\xbb\xbc\xe8\xb5\x9f\x7b\xed\x8d\x1a\x7b\xab\x26\xc1\x4b\x9f\x70\x77\x26\x84\x56\x5a\xe2\xb4\x28\x0a\x41\x4f\x00\xaa\xb9\x35\xb7\x51\xd0\x3e\x63\x6e\x42\x4d\x97\x52\xb6\xb1\xad\x21\x94\x54\xe7\x66\x47\x84\x6b\xcd\xba\xf3\x71\x3c\xa5\x74\xd4\xfb\x0c\x67\x15\xd2\x49\xfa\xc3\x93\xfb\xdc\xcf\xba\x27\xdd\x61\x58\xab\x5e\x25\xb7\xfe\x62\xd1\x10\xbc\x36\x9c\x92\xb0\x3e\xc6\xfb\x83\x64\x74\xa3\xbe\xa1\xb3\x83\xe6\x7c\x46\x3e\xc5\x51\x3d\x33\x39\x4b\xa1\x94\x90\x34\xa4\xdc\xb3\x31\x17\x3c\x4f\x44\xa6\xec\xff\x6d\xca\xd1\x11\x24\xf0\x27\x6b\xb2\x9b\x8b\xcb\x0a\x1d\xaa\xf1\xdd\xb8\xce\x4f\xcb\xfd\x4c\x33\xde\xc1\x8a\xb5\xfb\x59\xbb\x91\x2d\x62\x92\x4d\x0e\x47\x63\x79\x3e\x1f\x4e\x85\x18\xa4\xd5\xd9\x64\x1a\x26\x61\xff\xba\x77\xb7\x60\x68\x94\xfe\xcd\xda\xe6\x46\x5a\x4a\xfa\x9e\x8d\x10\x8e\xf1\xf2\xe5\x35\x29\x30\xc4\xe8\x4c\xa1\xf8\x21\x03\x99\xa3\x8a\x5b\xa9\xb8\xe5\x9f\xd5\xa3\xf6\xd8\xd0\xb5\xd0\x86\xea\x15\xde\xb7\x46\x14\xbf\xfc\xec\xf1\xab\xd1\xdc\xf8\xef\x7b\xd7\x74\x56\xae\x36\x1e\xe6\x79\x3c\xba\x08\xba\xf1\xd5\x05\x3e\xbf\x0f\xb2\xc9\x78\xf9\x65\x79\x3b\xe0\xed\xfe\x49\x83\x55\x93\xf6\xe4\xb2\x97\x76\x0e\x93\x4e\xfb\xb4\xb9\xe8\x5c\xf6\xfc\xc1\x69\xe3\x66\x89\x7f\x1f\xcd\x3f\x78\x31\xda\x63\x93\x0f\x13\x19\x09\x18\x96\xf0\x22\x2a\x26\x08\x56\x91\x11\x8e\xea\x6e\xe5\xc0\xad\x6f\x81\x37\xa4\x45\x6a\x66\x81\x90\xa1\x39\xb8\xf3\xfa\x2b\x93\x2d\xb4\x79\xde\x56\x2b\xe0\xc4\xbc\x76\x3f\x02\x00\x00\xff\xff\xee\x02\xd0\x16\x00\x05\x00\x00")

func templatesCommonHeadGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonHeadGoHtml,
		"templates/common/head.go.html",
	)
}

func templatesCommonHeadGoHtml() (*asset, error) {
	bytes, err := templatesCommonHeadGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common/head.go.html", size: 1280, mode: os.FileMode(420), modTime: time.Unix(1490699267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCommonMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x72\x83\x30\x0c\x45\xf7\x9c\x42\xa3\x7d\xe2\x0b\x38\x5c\x85\x11\x58\x0c\x4c\x55\x93\x31\x76\x37\x1a\xdf\xbd\x4a\x83\x69\x9b\xac\x24\xa1\xc7\xff\xdf\x52\x0d\x3c\xaf\x91\x01\x3f\x39\x16\xac\xb5\xf3\x91\xbe\x60\x12\xda\xf7\x1b\x5a\x3b\x52\x82\x67\xb9\x18\x49\x45\x32\xf6\x1d\x80\x0f\xeb\x49\x4d\x5b\xcc\x64\x1a\xe9\x32\x4b\x59\xc3\xcf\xfe\x3f\x71\x08\x2c\x4c\x81\xd3\xb1\x37\x82\x5e\xf6\x63\xa2\x18\x10\x96\xc4\xf3\x0d\x55\xaf\x23\xed\x3c\xdc\x29\x2f\xb5\x3a\xec\xed\xc3\x23\xe3\x90\xd7\x2c\x5c\xab\x77\x74\x18\x39\x73\x7a\xf7\x9c\x36\x11\xba\xef\xdc\xd2\xb7\xf9\xd7\xbe\xc8\x1f\xff\x86\x59\x39\x09\xd5\x51\xb6\xe9\xe3\x79\x9b\xe1\xf1\x4c\x8e\x19\xe1\x6a\x57\x6a\x00\xc7\x70\x4e\xde\x15\x79\x89\x74\x34\xde\x99\x6c\xdf\xa9\x82\xf1\x60\x3f\x7c\x07\x00\x00\xff\xff\x50\x2f\x43\xbe\x77\x01\x00\x00")

func templatesCommonMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonMenuGoHtml,
		"templates/common/menu.go.html",
	)
}

func templatesCommonMenuGoHtml() (*asset, error) {
	bytes, err := templatesCommonMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common/menu.go.html", size: 375, mode: os.FileMode(420), modTime: time.Unix(1490699267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMonitorIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\x4d\x6e\xc3\x20\x10\x85\xf7\x39\xc5\x08\x65\x59\xc7\x6a\x94\x55\xe5\xb0\x69\x36\xd9\xf5\x06\x15\x09\x93\x18\x89\x82\x05\xf4\x17\x71\xf7\x8e\x5d\x9b\xc4\x76\xd4\x6c\x6c\xcf\xbc\x6f\xcc\x7b\x40\x8c\x12\x4f\xca\x20\xb0\xa3\x35\x01\x4d\x60\x29\x2d\x2a\xa9\x3e\xe0\xa8\x85\xf7\x5b\xe6\xec\x27\xe3\x0b\x80\xeb\x5e\x8b\x0a\x1a\x72\x9d\x32\xd5\x74\xf1\x26\x8b\xc7\x75\xaf\x91\x5a\xaf\xf9\x8b\xb3\x47\xf4\xde\x3a\x5f\x95\x54\xf6\x4a\x8c\xcb\x83\xf0\xf8\xda\x88\x50\xc3\xd3\x16\x56\xb9\x22\x13\x03\xe2\x84\x39\x23\x2c\x95\x91\xf8\xf5\x00\xcb\x66\xf8\x53\x37\x90\x2b\x9f\x27\xee\xd9\x99\x12\x8d\x30\xa8\xa1\x7b\x16\xb4\x17\xe2\x5d\x87\x11\x7b\x83\x2e\x6a\x14\x52\x99\xf3\x84\x6b\xa3\x6e\xc6\x60\x50\x41\x23\xe3\x95\x80\xda\xe1\x69\xcb\xae\x13\xa7\x54\x66\xfb\x25\x09\x5d\xc2\x94\xd8\x65\xb3\xe0\xd2\xad\x4a\x31\x5b\xac\xac\x37\x7c\x31\x76\x5a\x92\xd5\x7b\xe6\x0f\x56\x7e\xcf\x9d\x4b\x3d\x50\x92\xf2\x59\xa7\x7e\xda\x43\xd6\x33\xb0\x45\x03\x7f\xd6\x8a\xee\x0a\xec\x77\xb4\x62\xb8\x85\x48\x4e\xde\x73\xbc\xd5\x1f\xbf\xdf\xb5\x41\x48\x9b\x25\x91\x9a\xff\x1f\x64\xd2\x18\x95\x31\xa2\x91\xfd\xf9\x67\xa1\xff\xe8\x5f\x03\xf2\x1b\x00\x00\xff\xff\x68\xfb\xe0\xd6\xef\x02\x00\x00")

func templatesMonitorIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMonitorIndexGoHtml,
		"templates/monitor/index.go.html",
	)
}

func templatesMonitorIndexGoHtml() (*asset, error) {
	bytes, err := templatesMonitorIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/monitor/index.go.html", size: 751, mode: os.FileMode(436), modTime: time.Unix(1490699267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMonitorMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8e\x41\xaa\xc3\x30\x0c\x44\xf7\xff\x14\x42\x64\xf9\x53\xef\x8b\xed\x4d\x0f\x12\x94\x58\x69\x0d\x89\x0a\x76\x52\x0a\xc2\x77\xaf\x68\x5a\xba\x12\x0c\xf3\x9e\x46\x35\xf1\x9c\x85\x01\x57\x96\x7d\x98\xee\xb2\xb1\x6c\xd8\xda\x9f\x6a\x21\xb9\x32\x74\x59\x12\x3f\xff\xa1\x1b\xe0\x1c\xe0\x64\x8d\xba\xaf\x5c\xaa\x55\x00\xfc\x92\xa3\x27\x98\x16\xaa\x35\xa0\xd0\x63\xa4\xd2\x8f\x06\x26\x84\x5b\xe1\x39\xa0\xfb\x02\x4e\xf5\x50\xb5\x86\xf1\xf2\x09\xfb\x5f\xe8\x1d\x45\xef\xcc\x67\x9f\x59\xd2\x7b\xc1\x71\x5f\x01\x00\x00\xff\xff\xf2\xc3\x81\x2e\xa4\x00\x00\x00")

func templatesMonitorMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMonitorMenuGoHtml,
		"templates/monitor/menu.go.html",
	)
}

func templatesMonitorMenuGoHtml() (*asset, error) {
	bytes, err := templatesMonitorMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/monitor/menu.go.html", size: 164, mode: os.FileMode(420), modTime: time.Unix(1490699267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMonitorProcessorGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x59\x5f\x93\x9b\x38\x12\x7f\xf7\xa7\xd0\x71\x95\x18\x62\x1b\xec\x49\x2e\x0f\xf3\x2f\x75\x77\xbb\x5b\x9b\xaa\xcd\x9f\x4a\xb2\xd9\x87\xa9\x94\x4b\x03\xb2\x4d\x02\xc8\x85\x84\x3d\x53\x2e\xbe\xfb\x76\x4b\x80\x05\x16\x63\xcf\xee\x26\x79\x70\x40\xfd\xd3\xaf\x5b\xdd\xad\x96\x9a\xd9\xed\x22\xb6\x88\x33\x46\x9c\x90\x67\x92\x65\xd2\x29\xcb\xc1\x65\x14\x6f\x48\x98\x50\x21\xae\x9c\x9c\x6f\x9d\xeb\x01\x21\xe6\x18\x42\x29\x4c\xca\x27\x8b\xa4\x88\x23\x25\xef\x22\x92\x49\x1a\x4d\x66\x67\x16\xd9\x9a\x66\x2c\x21\xea\x77\x02\xca\x69\x91\xc8\x0a\xd5\x8f\x5b\x31\x1a\xc5\xd9\xb2\xc1\x01\x72\x35\xbb\xde\xed\xfc\x75\xce\x43\x26\x04\xcf\xdf\xd2\x94\x95\xe5\x65\x00\xc3\x35\x57\x00\x64\xbd\xc4\x93\x5b\x1e\xdd\x9b\x7c\x92\xde\x26\xac\x86\xe8\x17\xf5\x3b\x11\x32\x8f\xd7\x2c\x32\xb0\x88\x46\x8b\xcc\x11\x1c\xcb\xdb\x03\x0a\x76\xfd\x09\x49\x2e\x03\x78\xb2\x08\xdf\xd3\x5c\xc6\x32\xe6\x59\x1f\xe0\xa3\xa4\xb2\x10\x7d\xd2\x0f\x54\x32\xe2\xa6\x62\x19\x08\xef\x01\x86\x5c\xf6\x09\xff\x5f\xe4\x39\x04\xbd\x4f\xfc\xeb\x1f\x6f\xfa\x44\x3f\x7f\xfa\x40\xdc\x07\xb4\x66\x74\x0d\xd6\xc5\x9c\xb8\xc5\x3a\x02\x33\x45\x00\x7a\xf2\xfb\xde\x09\xef\x73\xbe\xcc\x21\x92\x87\x72\x18\x69\x39\x16\x11\x1d\xe7\x5f\x4a\x8c\x26\x89\x23\x0c\x6e\xe5\xd1\xcf\x31\xdb\x3a\x9d\x79\x88\x32\x42\x1e\xa8\x00\x77\xf2\x45\xe7\xab\x08\x21\xe8\x92\xc8\xfb\x35\x83\x6c\x60\x77\x32\xf8\x4a\x37\x54\x8f\x36\xac\x1b\x9a\x13\x70\x5f\xc4\xf2\x26\x8e\x82\x5c\x91\x45\x91\x85\xf8\xec\x36\xb6\x08\x6f\x37\x68\xd4\xe2\x2c\x94\xfc\x16\x0b\x09\xe8\xb9\xbf\x48\xa8\x84\xad\xe7\xce\xfd\x94\xae\x8d\x49\xe3\x3d\xd3\x86\x26\x05\x1b\x93\x18\x74\xdd\x01\xd7\x7e\x51\xf1\xc2\xfd\x97\x1a\xf5\xe3\x2c\x4c\x8a\x88\x09\xd7\x59\xf2\x6f\xd4\xf1\x5a\x30\x02\x76\xca\x22\xcf\xc8\xcd\x97\x0b\x63\xb8\x34\x9e\x95\x8a\x9b\x61\xca\x20\xdd\xc3\x0c\x76\xd3\xf0\x0b\x58\xa7\xb9\xc5\x3a\x89\xa5\xeb\xf8\x8e\xe7\x8b\x24\x0e\x99\x7b\xe6\xf9\x5f\x79\x9c\xa9\xa1\x8b\xc1\xe0\x50\x8d\x62\x33\x74\x95\xde\x98\xc8\xbc\x60\x26\x3a\x08\x08\x54\x12\xc1\x13\xe6\x27\x7c\xe9\xaa\xd0\x91\x04\xbc\x82\x9c\x7b\x90\x89\xa9\xfd\x66\xd2\xd4\xfe\xfc\x94\x33\xa6\xfc\x89\x5e\x9c\xfb\xcb\x9c\x17\xeb\xff\xdd\x37\x53\xc6\xed\xb8\xb4\xdc\x53\x59\x8d\xe3\xfe\xde\x01\xbe\x28\x6e\x71\xf3\x67\x4b\x77\x3a\x3e\x10\x42\xa1\x90\xaf\xd1\x3b\xef\x16\xca\x0d\x5e\x7b\xb1\xc7\x43\x87\x76\x6b\x3e\x51\x9b\xfd\x19\xb1\xc2\x30\xbe\x9a\xbc\x27\xeb\x44\xb5\x31\x7d\x63\xb7\x7b\xd3\x67\xf1\xd0\x1f\x7a\xa3\x99\x77\xd1\x62\x6b\xd9\xad\x27\xf6\xe9\xd3\xd2\x9b\xe9\x97\x2e\x83\xf9\x5e\xad\xee\x66\xd8\x4e\x27\x2b\x44\x49\x4c\xcc\x61\x78\x2a\xac\xe9\xe7\x63\xd9\x24\x21\x27\xda\xd9\x44\xba\xe9\x84\x69\xd3\x4d\x27\x5d\xb2\x9a\x4d\xfd\x5e\x9d\x44\xb6\x8d\x6d\xb8\x07\xe7\xe9\x63\xe3\x8a\x34\x72\x5f\xc7\xa3\xbb\x7d\xa6\xe3\x59\xbd\x81\x4c\xeb\xea\x4c\x3e\x81\x61\x36\x3e\x33\x19\x5a\x14\x58\x1d\x81\xc2\xb9\x94\xd1\xb5\x33\x52\x36\x8d\x1c\x28\x76\xd1\xb5\x1e\x41\xee\x6a\xc0\x19\x74\x12\x52\xd5\x51\x63\xa5\x9d\xf8\x8b\x6d\x2c\xc3\xd5\x7e\xfd\xbe\x50\xe7\x93\xaf\xf2\xb4\x83\x0d\xa9\x60\x64\x7a\x8e\x3e\xe7\x6b\x96\xd9\xf3\x68\x88\x26\x5d\x0a\x38\x92\xeb\xa3\x37\xa1\xb7\xe0\x6c\xf5\xbb\xbf\x1e\xbc\x03\x82\xcb\x00\x61\xd7\xca\xee\xe1\xc5\xa1\xaa\x99\x52\x95\xb3\x90\x6f\x18\x26\xff\xa0\xa3\x11\x97\xb7\xc8\x79\xda\xf2\x6e\x05\x9f\x0b\x3c\x25\xe7\x7c\xb1\x10\x4c\x92\xa7\x4f\x8f\x20\xf4\x7a\xc9\xf5\xd5\x94\xbc\x3a\x0d\x7a\x4e\xa6\x17\x16\x7b\x72\x3c\xc0\xaf\x88\x7b\xc8\xa1\x24\x36\x72\x14\xdc\x0c\x67\xa9\x8f\x0f\xb0\x61\x80\xd9\xb3\x51\x4b\x6e\x5d\xe8\x6a\x9b\x5a\x69\x61\xfc\x61\x53\x43\x7d\x55\xb0\x92\x56\xb2\xda\x7f\x36\xfe\x36\xc4\x54\xd5\xd1\x85\x5b\x94\x26\x61\x91\xa0\x03\xd6\xd5\xad\xc0\x62\x4f\x2d\x42\xf7\xb9\xb5\x71\x13\x15\x61\xef\xd9\x6c\x3a\x25\x01\x71\x25\x9f\xa8\x77\xcf\x97\xfc\x97\xf8\x8e\x45\x70\x70\x75\xd7\x66\xf0\xc4\xe2\x2d\x7d\xeb\xd6\x03\xde\xab\xe9\x79\xfd\x6c\x73\x08\x54\x23\xa5\x1b\x1c\x3d\xa9\xbd\xe3\x81\x56\x8c\xcb\x5e\xe1\x61\x70\xcc\xcc\xff\xd0\x64\xab\x4e\x6b\x32\xea\x80\xf1\x9f\x42\x0e\x47\xc8\x6b\xac\x63\x34\x3c\x3e\x05\x17\x7f\x0a\xae\xb2\xfe\x14\xa8\xe4\xa7\xa0\xc0\x37\x35\xac\x17\x35\x39\x22\x6f\x5d\xd9\xab\x38\x38\x44\xc8\xfb\x04\xae\x64\x29\xcd\x97\x71\x06\x97\x78\x29\x79\x7a\x3e\x5d\xdf\x5d\x38\x76\x2a\x24\xb3\x10\x4d\x6e\x69\xee\x90\x9c\x23\x57\x3d\xa6\x86\x6a\x7e\x20\xdf\xc6\x91\x5c\x9d\x93\x17\x2c\xbd\x20\xfa\x19\xd6\xbc\xcf\x97\x11\x19\x3e\x41\xad\x56\x3f\xec\xf3\x6a\x34\x7c\xa2\x6f\x96\xf5\xaf\x2e\x5f\x87\xf5\xeb\xcc\xac\x5f\x60\x07\x4d\x12\x16\xfd\xad\xaa\x59\x65\x17\xf9\xa8\xb9\x8e\x15\xd0\xe7\xa6\x01\x1d\xd5\xdf\xad\x58\x7d\xd7\xd2\xd2\x55\x84\xc5\xcf\x54\xd2\x3c\xcd\x13\x4e\x23\x4b\x75\xb4\x01\x7a\x6b\x97\x6e\x07\x16\xe0\xbb\x2c\x44\x47\xbd\xa1\x72\xe5\x43\x22\xd5\xd5\x69\x5c\x8d\xd0\x3b\x17\x68\x26\xe8\x92\x67\x2f\xe1\x66\x39\xf5\x3a\x3e\x81\x4b\xbd\xb2\xf4\x0a\x44\xdd\xfb\x17\x2a\x49\xe8\x12\xe8\xa1\xc6\x75\x2e\x5f\x2c\x11\xac\x1f\xef\x36\xda\x8d\x72\xd9\xd8\xeb\xa1\x1d\x75\xdd\x54\xe6\x35\x92\xfe\xea\x59\x0e\x6c\x3d\x46\xab\xb4\x61\xda\xf5\x54\x8b\x47\xd7\x35\xb3\x6c\xc0\xf6\x0b\x02\x75\xd0\xf6\xf1\x1e\x2f\x6a\x15\x10\x56\xfb\x18\xa5\x58\xf8\x4f\x40\x89\xba\x03\xb6\x63\xff\x89\xea\xd6\x5b\xd9\x88\xf9\x32\x11\x45\x18\x2a\xf6\xbf\x52\xee\x30\x77\x1e\xa8\x74\x20\x86\x02\x07\x91\x87\xe3\x3a\xde\x30\x7c\x5f\xaa\xd3\xec\x94\x8a\xf7\x42\x17\x9c\x22\xcb\x0e\xae\x6b\x3f\xac\x2c\xb4\x4e\x63\x6d\xc9\xf1\xe3\xcd\x28\x81\xd8\x4e\x14\x29\x8b\xe6\x29\xb8\x94\x2e\x99\x38\x2c\x86\x76\x48\xb7\x2c\x3e\xea\x68\xdf\x67\x9a\x4a\x35\xeb\x36\x78\xfc\xe9\xde\x66\x85\x6d\x71\x1a\xd0\xb6\x1f\x6c\xae\xc2\x1d\x21\x56\x5c\xce\x75\x7b\x95\xab\xcf\x43\xa6\xa3\x6c\x80\x7d\xdc\x6c\x1e\xb2\x9c\x62\xff\x51\x49\x15\x26\x5c\x60\x52\xb5\x6d\xaa\x82\xed\x54\x52\xc7\x32\xfd\x65\x33\xbd\xef\xf0\x75\xb4\xb0\x33\xb9\x3a\x74\xcf\x7b\xe6\xfc\x9e\x7d\xcb\xf8\x36\xc3\x50\x49\xe6\x74\xab\xa8\xf1\xec\xb6\xbf\xa7\xe8\xae\x32\xc6\xee\xd3\x71\x6c\x4d\xb1\x6a\xf6\x46\xba\x6f\x1b\x29\xa8\x81\x2a\x5b\xed\xb1\xf6\xea\x98\x40\x2e\xc0\xa5\x80\x66\x11\x70\xa4\xb0\x55\x08\x0c\xd3\x56\x07\x89\xed\x63\xf4\xdc\x17\x2c\x61\x21\xb4\x9d\xff\x6e\x7f\x54\xf3\x2a\xc1\x7f\x93\x04\x5a\xd2\x46\x76\xcb\xef\x40\x84\x5c\x4d\x5f\x6d\x7c\x4e\x88\xbc\x5d\x6d\x72\xe4\xeb\x16\xdf\xfc\x58\x10\xf9\x2b\x99\x26\xae\xad\xf1\x36\x51\x3f\xc0\x26\x5f\xb9\xc7\xf5\x7c\xba\x86\xa6\x33\x72\x1d\x99\x03\x85\x2a\xb2\x90\x7a\x4e\x8b\xba\xfa\xa6\x75\x9a\xed\x3e\xbb\x83\x1e\xde\xf3\xb5\xd3\x5d\xaf\x8e\x4d\xd9\x04\x6c\x0b\x36\xf0\x2d\xac\x44\xbe\x46\x1b\x20\xf3\x5d\x5b\x23\x0e\x5e\xf8\x2a\x60\xc8\xd9\xed\xfc\x5b\xc8\xd9\xf9\x1a\x0e\xf6\xb2\x0c\x9a\x6f\xe1\xb8\xde\xc0\xfc\x38\xfe\xfa\xa7\xb2\x04\x63\xbb\xdf\x2a\x1b\xeb\xca\x31\x79\x31\x9d\x4e\xf7\xa9\xa7\x9b\xb0\x84\xc4\x92\xc4\x19\x80\xe1\xf9\x7e\xf0\x7d\xb4\xc3\xad\x54\x7d\x55\xad\xfe\x50\x50\x7f\xb4\xaf\x1e\xaa\xff\x76\x3b\x98\x5c\x96\x83\x3f\x03\x00\x00\xff\xff\x5c\x92\x44\xe9\xaf\x18\x00\x00")

func templatesMonitorProcessorGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMonitorProcessorGoHtml,
		"templates/monitor/processor.go.html",
	)
}

func templatesMonitorProcessorGoHtml() (*asset, error) {
	bytes, err := templatesMonitorProcessorGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/monitor/processor.go.html", size: 6319, mode: os.FileMode(436), modTime: time.Unix(1490699267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesQueryIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\xcb\x6e\xdb\x3a\x10\xdd\xe7\x2b\x26\x82\x71\xb3\xb8\x95\x85\x26\xbb\x46\x56\x37\xcd\xaa\x40\x81\xa2\x8b\x2e\x03\x5a\x1c\x4b\x44\x69\xd2\xe5\x23\x8e\x61\xf8\xdf\x3b\x12\x25\xc5\x96\x69\x25\x45\x51\xa0\x02\x2c\x7b\x34\x0f\x9e\x39\x3c\x1c\x6b\xbf\xe7\xb8\x12\x0a\x21\x29\xb5\x72\xa8\x5c\x72\x38\x5c\xe5\x5c\x3c\x41\x29\x99\xb5\x8b\xc4\xe8\x6d\x52\x5c\x01\x5d\xc7\x4f\x9b\x60\x46\x69\xa6\xf3\x8d\xfd\x2f\x59\x31\x6f\xa9\x65\xba\xe6\xe9\xfb\xdb\x51\x0c\x45\xd5\xb7\xc5\x57\x8f\x66\x07\x9f\x71\x67\xf3\x8c\xcc\x51\xc4\x7e\x2f\x56\x30\x47\x63\xb4\x21\xa4\xa3\xec\xa3\x35\x98\x44\xe3\xa0\xbd\xa7\x9c\xa9\x0a\x4d\x6f\x08\xbb\x16\xd6\x8a\xa5\xc4\x04\x8c\x96\xd8\xc5\x9e\x61\xa1\x7a\x4b\xef\x9c\x56\xe0\x76\x1b\x8a\x0a\x46\x32\x34\x21\xb5\xa5\x0a\x9c\x39\xd6\xd7\x1c\x2a\xe5\x76\xc3\x54\xf1\x9f\x13\x6b\xb4\xf7\x79\xd6\x5a\x79\x16\x0a\x44\x96\xb1\xce\x68\x55\x15\x0f\x4d\x53\xd7\x14\x1e\x4c\x6a\x35\xf4\x39\x7f\x88\x77\x9b\x51\xbb\xe7\xf4\xa0\xe2\x67\xa1\x81\xb4\x2d\x33\x4a\xa8\xea\xcd\xb4\x75\xf1\xff\x3c\x6f\xdf\x03\xce\x53\xe6\x3a\xf0\x7f\xce\x5d\xcf\x9e\xd5\xde\x94\x68\x23\xfe\x7c\xa5\xcd\x1a\xb4\xb2\x7e\xb9\x16\x6e\x91\x6c\x85\xe2\x7a\x3b\x97\xba\x64\x4e\x10\x0d\x0b\xb8\x21\x40\x4b\x66\xf1\x71\xc3\x5c\x7d\x38\x64\x64\x5a\x94\x58\x3a\xe4\x8f\xa1\x2e\x3d\xbc\x81\xff\xc1\x22\x33\x65\x3d\x7f\x62\xd2\xe3\x3d\x18\x74\xde\x28\x58\x31\x69\xf1\x3e\xc2\xf3\x78\xf3\x84\xda\x78\x97\x56\x46\xfb\x0d\x1c\xfd\x4e\x65\x75\x21\x79\xa2\x40\xba\x74\x6a\x22\xab\xcd\x9c\xda\x66\x4a\x07\xfa\xa4\x34\x5e\x98\x97\x0e\xb8\xd1\x1b\x62\x45\xa5\x4e\x57\x95\xec\x05\x10\x8c\x45\xd2\x7b\x93\x22\x46\x4d\x2b\x8b\x41\x3f\x8c\x68\x21\xa9\xbc\xa6\x8e\x13\xa4\x5e\xf6\xe9\x03\x8e\x35\x2a\xff\x4a\x83\xcd\xb5\xdf\xcf\x86\x9d\x83\x0f\x0b\x38\xde\xc7\x37\x24\x9b\x66\xf4\xc0\x8c\x14\x81\xcf\xef\x60\x16\x3a\x6a\xeb\x5c\xd6\xd3\x19\x7a\x29\x8a\x9c\x41\x6d\x70\xb5\x48\x8e\xf1\xb4\x4a\x9a\xf5\x2c\x35\xdc\x0d\x46\x9e\x31\xe2\x86\x12\xdf\x80\x31\xae\xfa\x13\x04\x99\x97\x13\x0a\x8a\x1d\xa6\x23\x77\xab\x2a\x10\x7c\x91\x04\x79\x27\x9d\x62\x1c\x3e\xbb\x41\x2f\xcd\x19\x4a\x9b\xbf\x15\x1a\x2d\x34\x5f\xf0\xa7\x17\x06\xf9\xdf\x97\x6d\x38\xb4\x97\x64\x9b\x14\xdf\x5a\xc8\xaf\xcb\x6c\x82\x84\x0b\xae\x3c\x6b\x7a\x3e\x7f\x4e\x1b\x42\xe7\x3d\x36\x67\xa6\xe7\xf4\x68\x2a\x7f\xd1\xd0\x49\x0c\x56\xda\x2b\x7e\x0d\x9f\x04\x87\x9d\xf6\x64\x9a\x0a\x1d\x38\x0d\xcc\x39\x56\xd6\xe0\x6a\x5c\x7f\xbc\x80\x32\x26\x8f\x51\xe8\xc8\x0c\xe3\xb2\x1d\x61\x47\x79\x79\x6d\xb2\xf8\xbb\x02\x1d\x63\x94\xd0\xde\x87\x69\x11\x2c\xeb\x4b\x82\x6f\x27\xde\x24\x42\x5c\x8d\x8c\x37\x04\x44\x38\xae\xef\x4e\x43\x9d\x70\x34\x7c\x9a\x29\xf3\x03\x77\xcd\x31\xa9\xef\x8a\xa9\xde\xe2\x0b\x2e\x35\xdf\xc5\x56\xdb\x18\x6c\x4a\x77\xbd\xe7\x59\x63\xff\x16\x73\x2f\x5c\x77\xae\xee\xab\xf7\xfc\x0a\x00\x00\xff\xff\x05\xda\xbe\x9a\xaf\x09\x00\x00")

func templatesQueryIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesQueryIndexGoHtml,
		"templates/query/index.go.html",
	)
}

func templatesQueryIndexGoHtml() (*asset, error) {
	bytes, err := templatesQueryIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/query/index.go.html", size: 2479, mode: os.FileMode(420), modTime: time.Unix(1490699267, 0)}
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
	"templates/common/base.go.html": templatesCommonBaseGoHtml,
	"templates/common/head.go.html": templatesCommonHeadGoHtml,
	"templates/common/menu.go.html": templatesCommonMenuGoHtml,
	"templates/monitor/index.go.html": templatesMonitorIndexGoHtml,
	"templates/monitor/menu.go.html": templatesMonitorMenuGoHtml,
	"templates/monitor/processor.go.html": templatesMonitorProcessorGoHtml,
	"templates/query/index.go.html": templatesQueryIndexGoHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"common": &bintree{nil, map[string]*bintree{
			"base.go.html": &bintree{templatesCommonBaseGoHtml, map[string]*bintree{}},
			"head.go.html": &bintree{templatesCommonHeadGoHtml, map[string]*bintree{}},
			"menu.go.html": &bintree{templatesCommonMenuGoHtml, map[string]*bintree{}},
		}},
		"monitor": &bintree{nil, map[string]*bintree{
			"index.go.html": &bintree{templatesMonitorIndexGoHtml, map[string]*bintree{}},
			"menu.go.html": &bintree{templatesMonitorMenuGoHtml, map[string]*bintree{}},
			"processor.go.html": &bintree{templatesMonitorProcessorGoHtml, map[string]*bintree{}},
		}},
		"query": &bintree{nil, map[string]*bintree{
			"index.go.html": &bintree{templatesQueryIndexGoHtml, map[string]*bintree{}},
		}},
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

