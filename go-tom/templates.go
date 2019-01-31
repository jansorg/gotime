// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/reports/html/commons.gohtml
// templates/reports/html/default.gohtml
// templates/reports/html/timelog.gohtml

package tom


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

var _bindataReportsHtmlCommonsgohtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\xdb\x6e\xdb\x38\x10\x7d\xcf\x57\x10\x5a\x14\x68\x03\x53\x8e\xd3" +
	"\x4d\xb7\x55\x53\x63\x7b\xc5\xf6\xa1\x8b\x45\xdc\xfd\x00\x4a\xa4\x65\xd6\xbc\x08\xe4\x38\xa9\x57\xf0\xbf\x2f\xa8" +
	"\x8b\x2d\x51\x94\xe2\xa2\xcd\x4b\x24\x71\x38\x73\x66\x78\xe6\x90\x74\x59\x52\xb6\xe6\x8a\xa1\x68\xb5\x93\x92\x98" +
	"\x7d\x74\x38\x5c\x20\x84\x50\x59\x62\x34\xbf\xcc\x35\xec\x0b\x96\xa0\x9c\xc3\x66\x97\xc6\x99\x96\xf3\x6f\x44\x59" +
	"\x6d\xf2\x79\xae\x81\x4b\xd6\xfe\x33\xac\xd0\x06\xe2\x3b\x66\x77\x02\xde\xed\xb2\x2d\x83\xcb\x39\xc2\x87\xc3\x45" +
	"\xe5\xed\x16\x48\x2a\x18\xca\x04\xb1\xf6\x4d\x64\x9b\x50\xcb\x6a\xac\x1e\x4f\x35\xdd\x9f\xde\xcb\x92\xaf\xe3\xcf" +
	"\xf6\x1f\xa3\xbf\xb1\xac\xf1\xd7\x00\x3b\x4d\x31\xcb\xde\x87\xfa\x23\x5d\x96\x25\x5f\xbc\x54\x28\x6a\x26\x47\x87" +
	"\xc3\xed\x1c\x68\xd0\xb6\x05\x04\x1c\x04\x8b\x96\x65\x19\x7f\x75\x4f\xa1\x19\xb7\xf3\x6e\xbc\xb2\x64\x8a\x76\x00" +
	"\x39\xbc\x48\x69\x40\xf1\x07\x02\xec\x8e\xa8\x9c\xc5\x1f\x65\x01\xfb\x1f\x05\xfd\x95\x4b\x86\x8c\x9b\x9f\x9c\x07" +
	"\x5c\x32\x5c\x99\x57\xe8\x4f\xc1\xbf\x70\xc5\x25\x11\x2b\x30\x5c\xe5\x3f\x9e\xcf\x00\x67\x0f\xa3\x21\xd9\x96\x51" +
	"\xe4\x82\x87\x51\xfa\x08\x05\x5b\x83\x03\x28\xb9\xfa\xb0\x33\x04\xb8\x56\x28\x6e\x9f\x7c\x07\x3e\xb2\x8a\x09\x77" +
	"\x7a\xa7\x28\xa3\x3f\x5a\xce\x8f\xdf\x49\x06\x08\x1e\x05\x7c\x36\xe8\xca\xe1\x18\xf2\x9f\xac\xeb\x27\x43\x24\xb3" +
	"\xa3\x15\x5d\x96\xe5\x5a\x1b\x49\xe0\xef\x9d\x4c\x99\x41\x71\x65\xff\x5e\xef\x14\x4c\x97\xf0\x76\xde\xe9\xae\xdb" +
	"\x79\xd5\x8a\xcb\x8b\x16\xda\xc5\x49\x02\x32\x2d\xa5\x56\xef\x57\xab\x56\x04\x6e\x2d\xec\x05\x43\x4e\x03\xde\x44" +
	"\xc0\xbe\xc3\x3c\xb3\xb6\xd3\xb6\x89\xd1\x1a\x50\xd9\x43\x8a\xf1\x5a\x2b\xc0\x96\xff\xc7\x12\xb4\xb8\x2e\xe0\x75" +
	"\x68\x78\x4d\x24\x17\xfb\x04\x45\x76\x6f\x81\x49\xbc\xe3\xd1\x0c\x61\x52\x14\x82\xe1\xfa\xd3\x0c\xbd\x13\x5c\x6d" +
	"\xbf\x90\x6c\x55\xbd\x7f\xd2\x0a\x66\x28\x5a\xb1\x5c\x33\xf4\xef\xe7\x68\x86\xee\x74\xaa\x41\xcf\xd0\x5f\x4c\xdc" +
	"\x33\xe0\x19\x99\xa1\xb7\x86\x13\x31\x43\x96\x28\x8b\x2d\x33\x7c\x3d\x43\xd1\x5b\xe7\x14\xbd\xd7\x42\x1b\xf4\x51" +
	"\xea\x6f\x2e\xd2\xd1\x4d\xe0\xcb\x6a\x2f\x53\x2d\x22\x1f\x76\xc5\x89\x1e\x76\xa9\x95\xb6\x05\xc9\xd8\xb8\xe9\x03" +
	"\xe3\xf9\x06\x12\xa4\xdc\xc2\x89\xd7\x17\x9e\x61\xe6\x40\x25\x28\x15\x24\xdb\xfa\x4e\xd2\xbc\x1d\x7e\xd8\x70\x08" +
	"\xc4\x00\xc1\x5a\x8b\xdf\x9e\x2f\x5e\xdc\xa4\xbf\xfb\x36\x46\x3f\x60\x4d\xe9\xd1\x8a\x55\x7f\x03\x14\x15\x21\xf0" +
	"\x03\xa7\xb0\x49\xd0\xe2\xea\xea\xc9\x20\x56\x65\x40\x09\x90\xd6\xea\xd5\xcd\x88\x11\x6c\xda\x68\xf7\xc4\x3c\x6d" +
	"\x32\x7c\x36\x6a\x9b\xe6\x3d\xeb\x36\xe7\xf1\x09\x6d\x41\x53\x2d\xe8\xa8\x91\x2b\x7d\xcd\xbf\xab\xf8\xa5\x61\x72" +
	"\x90\x71\xb3\x01\xe1\x54\x1b\xca\x4c\x1f\x71\xaf\x68\x1d\x20\x87\x93\x93\x0d\x48\xe1\xb1\xbe\xc3\xf9\xda\xcd\xf1" +
	"\x83\x97\xcb\xa3\xd5\x49\x49\xb6\xcd\x8d\x93\x3b\x7c\x4e\x6d\x7a\x8c\xec\x84\xae\x3f\x85\xf1\xd7\x9b\x71\x3f\x81" +
	"\x66\x61\x6b\x0f\x1d\x46\xf8\xe8\x8e\x15\x13\xa4\xb0\x2c\x41\xed\x53\x2f\xd0\x29\xd2\x66\x86\x80\x7a\xa1\x04\x57" +
	"\x0c\x6f\x9a\x85\x5c\xc4\xd7\x37\xd5\x12\x75\x2d\x0a\x42\x29\x57\xb9\x5b\xbf\x6a\x14\x2d\x7a\x26\x87\x8e\xff\xb8" +
	"\x6a\x04\x17\xa5\x7e\xf2\x62\x0d\x3c\x5d\x85\x2b\xb2\xf1\xcb\xe1\x7a\x0e\x57\xdd\xed\xba\xf7\xc1\x90\xa2\x0f\xd1" +
	"\xa9\x21\x26\x82\xe7\x2a\x41\x6e\xa7\xe8\x8f\xde\x33\xe3\x04\x49\xb4\x16\xa9\x06\xd0\x32\x1c\xda\x2f\x8f\x3f\x17" +
	"\x74\x31\x82\x99\x11\x3a\x44\x3e\xc6\x1f\xbf\xe7\xa6\x78\xd9\xef\xe5\x10\xe3\xba\x54\x1f\xf4\x5d\x68\x42\xdb\xb8" +
	"\xde\x94\xfa\x73\x98\xa6\xf1\x49\x77\x1e\x27\xeb\x49\x9d\x26\x9d\x69\x4a\x11\x98\x44\xb9\xdc\x36\x5c\xd0\xa7\xd7" +
	"\x0a\x2f\x9e\x0d\x17\x61\xac\x8a\x67\xa8\x43\xec\xd2\xe2\x74\xd0\x60\x7d\x46\xd5\xfb\x41\x20\xa9\x57\xaf\x9e\x84" +
	"\xdd\x56\x74\xc0\xc6\xd5\xcb\x73\xdd\xe5\x62\x35\x3e\x82\x2b\xd0\x1f\xfd\x65\x3f\xed\x29\x53\x4b\x38\xd4\xde\xee" +
	"\x8e\x1f\x8f\xf5\x6a\x9c\x56\xc7\x78\x8b\x96\x61\x28\x3d\x27\x2f\x3c\x27\x27\x2f\x8d\x76\x87\x19\x41\x76\xa0\x43" +
	"\x82\x95\xa0\xe7\xc5\x77\x64\xb5\xe0\xb4\xc9\x35\xb4\x05\x78\x49\x4b\x62\x72\xae\x12\x74\x85\xae\xd0\xf3\x71\xed" +
	"\x38\x22\x1a\x90\xe8\x28\x3e\xf1\xcd\x44\x3a\x29\x68\x20\xfe\x7e\x52\xc7\xc6\xa0\x8b\x64\x20\x7e\x81\xc9\x40\x67" +
	"\xdd\x37\x5f\x13\x1e\x59\xbe\xa6\x06\x55\xb0\xeb\x63\x9d\xbc\x53\x49\x4f\x79\x62\x77\xc4\x39\x9f\x85\xe8\x31\x45" +
	"\xed\xe9\x61\xc8\x79\x60\x9b\xf3\x0f\x64\x67\xa8\x8e\x77\x2e\x7b\x36\x2e\xe9\x3f\x93\xc0\xda\x1d\x8a\x7f\x4d\x12" +
	"\xbf\x08\x52\xd3\x7a\xfe\x7f\xeb\xe1\x9b\x5f\x36\xb4\x73\x1b\x5a\x52\x71\xfe\xf5\xe5\x7c\xca\x61\x88\xb4\x75\xc3" +
	"\x5c\x4f\x35\x0c\x10\x03\x98\x92\xfd\xac\x7d\x76\x65\xa8\x5e\x74\x81\x03\x85\x3b\x3b\x4f\xa5\x81\xf9\x59\x61\x69" +
	"\xf1\x66\x5f\x6c\x98\xb2\x21\x81\xc0\x0f\x2c\xdd\x72\x98\x32\x19\x19\x6a\xee\x48\xf3\xea\x92\xb4\xbc\x98\xb8\x31" +
	"\x21\xc9\x28\x27\x6f\xa2\xc2\x70\x05\x9d\xfb\xd3\x9f\x05\xc9\xfd\x5c\x6b\x05\x1c\x82\x68\x8b\xfb\x47\x7c\x23\xc3" +
	"\x72\xe0\x2e\x79\x93\x97\xb1\xab\x22\xbc\x33\x4c\x9c\x07\x87\xd7\x01\x07\x19\xa7\x86\x91\x2d\xe6\xca\x72\xea\xc0" +
	"\xde\x6b\x4e\xc3\x9e\xcd\x40\x12\x1f\x9b\xed\x59\x91\x35\x38\xf9\xf6\x2b\xef\x9d\x82\x46\x83\xa4\x6c\xad\xcd\xd9" +
	"\x41\x26\x8d\x26\xf0\x52\x6e\x0b\x41\xf6\x49\x5d\x49\xec\x20\x31\x83\xdd\xf9\x61\x44\x22\x02\x4b\xf5\x8b\x40\x0f" +
	"\x84\xe8\x27\xa3\x9c\x9f\xb5\x3b\x1a\x0d\x52\x46\xdd\x0e\x09\xfd\xe4\xb0\xb3\xa0\x65\xfd\x93\x43\x33\xfc\x7f\x00" +
	"\x00\x00\xff\xff\xc6\x23\x02\x63\x93\x14\x00\x00")

func bindataReportsHtmlCommonsgohtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataReportsHtmlCommonsgohtml,
		"reports/html/commons.gohtml",
	)
}



func bindataReportsHtmlCommonsgohtml() (*asset, error) {
	bytes, err := bindataReportsHtmlCommonsgohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "reports/html/commons.gohtml",
		size: 5267,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1548880419, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataReportsHtmlDefaultgohtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4b\x8f\xa3\x38\x10\xbe\xf3\x2b\xbc\xd6\x9e\x5a\x0d\x6c\xdf\x5a" +
	"\x2d\x13\x69\xfb\xb1\xda\xc3\x9c\xfa\x21\xcd\xd5\xc1\x95\xe0\x19\x6c\x67\x70\x91\x74\x84\xf8\xef\x23\x9b\x04\x48" +
	"\x62\xba\x39\xf4\xe4\x02\xd4\xe3\x2b\xe7\xfb\xca\x65\x37\x4d\x4c\xd2\xab\xb5\xc1\xfd\x06\xee\xc8\x5a\x62\x51\x2f" +
	"\x93\xdc\xa8\xf4\x07\xd7\xd6\x54\xeb\x14\x8d\x4a\xd7\x26\x76\x8f\x0a\x36\xa6\xc2\xe4\x19\x6c\x5d\xe2\x7d\x9d\xff" +
	"\x04\xbc\x4a\x49\xdc\xb6\x51\xd4\x34\x02\x56\x52\x03\xa1\x9d\x9d\xb6\x6d\x44\x08\x21\x5f\x03\xdf\x21\xfd\xbd\xf4" +
	"\x36\x72\x97\x91\x64\x30\x96\xc0\x57\xd6\xdb\xbe\x01\x5f\x3d\x14\xb2\x14\x15\xe8\xc1\x6f\x0b\xb3\x7b\x7a\xe7\x79" +
	"\x97\xf7\x3f\xb7\xcf\xa6\xd6\x02\xc4\x28\xd2\x87\x32\x21\xb7\x24\x2f\xb9\xb5\x19\xed\x0a\xd1\x85\x77\x74\x38\x72" +
	"\x45\xba\x52\x07\xe4\xe3\x8f\x21\x5f\x96\x70\x4c\xf4\x1f\xb1\x11\x62\x94\x3b\x44\x16\xc0\x45\xc8\x5e\x5d\x1a\x0f" +
	"\x09\x3d\x6e\x11\xef\xa4\x00\xba\x60\x76\xc3\x75\x6f\x95\x58\x02\x5d\x34\x4d\xf2\xea\xde\xda\x96\xa5\xce\xbd\x60" +
	"\x29\x16\x9f\x43\x4a\xe5\x73\xe5\xcd\xad\x26\xf4\xb1\xae\x38\x4a\xa3\xa9\x03\x99\xca\xee\x58\xe8\x09\x3d\x63\x62" +
	"\x56\x9d\x4e\x09\x31\xaf\x1a\x68\x11\xa8\xc1\xd2\x10\x61\x0e\x67\x82\xde\xa5\x11\xfb\x4b\x7b\xd3\x54\x5c\xaf\x21" +
	"\xac\xea\x90\x3c\xa1\x4d\xe7\x14\x27\xdc\x63\xa0\xfa\x28\xf6\x9c\x10\x25\xf5\x91\x75\x92\x1c\xdf\x3e\xc3\x99\x2f" +
	"\xc1\x9c\xaa\x1e\x63\x7e\xe9\xb0\x1e\x64\x52\x93\x0f\x14\x0c\x6b\xc2\x70\x65\x0c\x1e\x97\x6c\xeb\x25\x1a\xe4\x65" +
	"\x70\x2f\x4d\xef\x99\xbe\xd7\x5e\x7d\xf2\x47\x2d\x16\x68\xd3\x31\x3f\x87\x81\x73\x2a\xce\x9f\xd8\x1b\xa1\xa2\x97" +
	"\xda\x7c\xd5\x3e\x71\x1c\x9f\x3a\x58\xea\x07\xd7\x78\xde\x41\x69\xe1\x0c\xd4\xff\xc5\xe4\x49\x6d\x70\x1f\x28\xe7" +
	"\xbd\xda\x20\x49\x1e\x39\xc2\xb3\xdb\x5b\x93\xb1\xe4\x6c\xde\xfa\x41\x46\xfc\x7e\xf4\xe3\xac\x47\x70\xff\x5b\xc8" +
	"\xed\xbc\xd6\x0a\x2e\xfa\xbc\x54\xc7\xae\x0d\xf4\x94\x8f\x9c\x33\x5e\x27\x54\xe8\xc6\xc9\xe1\xf8\x9a\x1a\x28\xa7" +
	"\x4c\x7d\xc4\xcf\x10\x8d\xa0\x36\x25\xc7\xe1\x6c\xed\x8f\xbf\x70\xfc\xf4\x36\x9d\x6e\x95\x0b\x92\xcf\x43\xc7\xdf" +
	"\x87\xf0\xa3\x29\x62\x7f\x09\x93\xbb\x13\x9e\x14\xa8\xca\x45\xc4\xdc\x83\x94\x5c\xaf\x33\xda\x34\xfe\xe5\x9e\x5b" +
	"\x20\x6d\x4b\x9d\xb3\x1f\xd3\x4c\x01\x72\x92\x17\xbc\xb2\x80\x19\x7d\x7b\xfd\x2f\xbe\xa5\x63\x97\xe6\x0a\x32\xba" +
	"\x95\xb0\x73\x37\x03\x3a\x5a\x60\x6e\x34\x82\xc6\x8c\xee\xa4\xc0\x22\x13\xb0\x95\x39\xc4\xfe\xe3\x9a\xd4\x16\xaa" +
	"\xd8\xe6\xbc\x74\x4d\x9d\x69\x73\x4d\xa4\x96\x28\x79\xe9\x8d\x90\xdd\x24\xff\x5c\x13\xc5\xdf\xa5\xaa\xd5\x89\x49" +
	"\xea\x53\xd3\xc9\x62\x0a\xc4\x4d\x0c\xbf\x6a\xb9\xcd\xe8\xf7\xf8\xed\xdf\xf8\xc1\xa8\x0d\x47\xb9\x2c\x81\x0e\xeb" +
	"\x91\x90\x81\x70\x6d\x7c\xb8\x55\xf8\x26\x72\x27\xb2\x7f\x46\xd1\xb9\xa4\xb9\x51\xca\xe8\x87\x97\x97\xe1\xc6\x34" +
	"\x72\xd6\x16\x8d\x3a\x38\x59\xda\x31\xc7\xba\xd9\x39\x8e\x7b\xa9\x95\xe2\xd5\x9e\x1e\xbb\xaf\x6d\xa3\x60\xdb\xf4" +
	"\x5e\x96\x76\x20\x2c\xf5\x8a\xfd\x0e\x00\x00\xff\xff\x3e\x86\xdf\xec\x05\x0a\x00\x00")

func bindataReportsHtmlDefaultgohtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataReportsHtmlDefaultgohtml,
		"reports/html/default.gohtml",
	)
}



func bindataReportsHtmlDefaultgohtml() (*asset, error) {
	bytes, err := bindataReportsHtmlDefaultgohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "reports/html/default.gohtml",
		size: 2565,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1548880291, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataReportsHtmlTimeloggohtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\xdd\x6b\xe4\x36\x10\x7f\xdf\xbf\x62\x2a\xfa\x14\x62\xbb\x47\x29" +
	"\x1c\x87\x1d\x68\xf3\x41\x0b\xa5\x94\x6c\x0e\xfa\xaa\xb5\x66\x6d\xf5\xf4\xe1\x5a\xe3\xcd\x2d\x66\xff\xf7\x22\xc9" +
	"\xde\xf5\xe6\xec\x24\x7b\x7e\x91\x35\xa3\xf9\xcd\x97\x66\x46\x7d\x9f\x40\x76\x55\x59\xda\x37\xf8\x09\x2a\x49\x75" +
	"\xb7\x49\x4b\xab\xb3\x7f\xb9\x71\xb6\xad\xb2\xca\x92\xd4\x38\x2e\x2d\x36\xb6\xa5\xf4\xb7\xae\xfc\x82\xf4\x18\x36" +
	"\x57\x19\x24\x87\xc3\x6a\xd5\xf7\x02\xb7\xd2\x20\xb0\xc8\x65\x87\xc3\x0a\x00\xe0\x4d\x0d\x64\x75\x56\xd9\xc4\x2f" +
	"\x03\xfc\x23\xba\x4e\x51\x84\x19\xe0\x23\xd2\x8f\x9b\x40\x83\x4f\x05\xa4\x27\xa2\x42\xbe\x75\x81\xf6\x27\xf2\xed" +
	"\x6d\x2d\x95\x68\xd1\x9c\xf8\xae\xb6\xcf\xf7\x5f\x79\x19\xe5\x7e\xe7\xee\xd1\x76\x46\xa0\x98\x9c\x0c\x47\x73\x21" +
	"\x77\x50\x2a\xee\x5c\xc1\xa2\x22\x76\x13\x18\x11\x47\x6e\x21\xbd\xd7\x0d\xed\x07\xe4\xf1\xcb\x89\x6f\x14\x8e\x82" +
	"\x61\x93\x08\x4e\x1c\xe2\xaf\x15\x62\x02\x73\x12\xaa\x91\x8b\x39\x7a\xfb\x2d\x31\xea\xcf\xae\x50\x6f\x50\x08\x14" +
	"\x40\x92\x14\x02\x59\xf8\x82\xd8\x0c\x3b\x6e\x44\xd4\xe8\x80\x6c\x85\x54\x63\x0b\xd2\x40\xd3\x4a\x43\x28\xe0\xef" +
	"\xbb\x07\x77\x95\xbd\x30\x7d\x62\x0d\x94\x56\xb9\x86\x9b\x82\xfd\xc2\x8e\xbe\x78\x60\x76\xd3\xf7\xe9\x93\xff\x3b" +
	"\x1c\xf2\x8c\xea\x19\x9b\xb3\x39\xa3\x17\x3d\x09\xca\xa2\x02\xc1\x29\xe0\xcb\x0f\x1f\x0d\xb0\x3b\xbf\x5b\x52\xf2" +
	"\xaa\xe4\x9a\x78\x4b\xdf\x27\x7a\x6f\xc4\x7b\x05\x7d\x05\x4c\xac\xed\x5a\x4e\xd2\x9a\x77\x4b\xd7\xc9\xb3\x14\x13" +
	"\x80\xbf\x2c\xa1\x5b\x94\x5e\x08\x6a\xb6\x78\x6f\x36\x56\xec\xbf\xa5\xf7\x7d\xcb\x4d\x85\x90\x3e\xb4\x5c\xa3\x1b" +
	"\x96\xc5\x6b\xb0\x90\xb2\xc8\x14\xa3\x2b\xce\x87\x3b\x11\x7c\xef\x9d\xd9\xda\x56\x73\xf2\xb9\x83\x34\xe4\x21\x78" +
	"\x34\x63\xe2\x22\xd0\x18\xd7\x88\xf4\x24\xf5\x77\x21\xd9\x66\x00\x5a\x3c\x0e\xc7\x32\xfe\xc3\xad\xc9\x36\x0d\x8a" +
	"\x85\x40\xcc\x4a\x48\x53\x29\xbc\xe3\x2f\xcb\x7f\x59\x6e\xea\xcf\xbd\x79\x9f\x2e\x54\x0e\x2f\x54\xe0\x43\x7f\xa1" +
	"\x92\x37\xcf\xbd\x75\xe6\xdd\x79\x19\x73\xab\xa5\x19\x0b\x06\xd2\xf1\xef\x82\xfc\x9a\x50\x2c\xbe\x15\x85\xb2\x79" +
	"\x4d\x72\xbe\x72\x96\x3c\xca\xb3\x85\xca\xc9\x69\x6b\x2d\x1d\xef\x57\xb7\x21\x4b\x5c\xcd\xb6\xf2\xd7\x1a\xdd\xd8" +
	"\x55\x7f\x3e\x15\xfe\x53\x00\xba\xb0\xe9\x2c\x07\x70\x19\xe3\xe6\xc2\xd6\xe2\x1d\x3e\x67\xe4\x59\x18\x29\xd3\x31" +
	"\x38\x73\x43\xa7\x83\x73\x66\x66\x08\xb9\x3b\x47\x3d\x36\xa5\x38\xe8\xe7\xda\x51\xdf\x13\xea\x46\xf9\xae\x32\xbe" +
	"\x27\x8e\x23\x7f\x62\xcb\x59\x4a\xa7\xfb\x41\xeb\x48\x5a\xe5\x3f\x08\x5b\xfa\x27\x08\xd4\xa4\xd5\xcd\x2a\xf7\x0b" +
	"\x28\x6e\xaa\x82\xa1\x61\x9e\x70\xec\xab\xb9\x46\xe2\x50\xd6\xbc\x75\x48\x05\xfb\xfc\xf4\x90\x7c\x64\x53\x96\xe1" +
	"\x1a\x0b\xb6\x93\xf8\xec\x9f\x2b\x6c\x62\x54\x69\x0d\xa1\xa1\x82\x3d\x4b\x41\x75\x21\x70\x27\x4b\x4c\xc2\xe6\x1a" +
	"\x3a\x87\x6d\xe2\x4a\xae\x7c\x48\x0b\x63\xaf\x41\x1a\x49\x92\xab\x40\xc4\xe2\x43\xfa\xd3\x35\x68\xfe\x55\xea\x4e" +
	"\x9f\x91\xa4\x39\x27\x9d\x19\x53\x13\x35\x09\xfe\xd7\xc9\x5d\xc1\xfe\x49\x3e\xff\x9a\xdc\x5a\xdd\x70\x92\x1b\x85" +
	"\xec\x64\x8f\xc4\x02\x45\x35\x76\xc7\x3c\x64\xc9\x5f\x8f\xb0\xae\x56\x2f\x43\x5e\x5a\xad\xad\xb9\x5d\xaf\x4f\xaf" +
	"\xb8\x09\xb3\x73\x64\xf5\xc0\xcc\xb3\x18\xb8\x3c\x56\xd1\xf4\xdc\xba\xd3\x9a\xb7\x7b\x36\xe6\xf9\x70\x58\xcd\xa6" +
	"\xf5\xc8\xcd\xb3\x08\x92\x67\x21\x49\xff\x07\x00\x00\xff\xff\x8c\xe9\x88\x64\x9c\x0a\x00\x00")

func bindataReportsHtmlTimeloggohtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataReportsHtmlTimeloggohtml,
		"reports/html/timelog.gohtml",
	)
}



func bindataReportsHtmlTimeloggohtml() (*asset, error) {
	bytes, err := bindataReportsHtmlTimeloggohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "reports/html/timelog.gohtml",
		size: 2716,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1547324009, 0),
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
	"reports/html/commons.gohtml": bindataReportsHtmlCommonsgohtml,
	"reports/html/default.gohtml": bindataReportsHtmlDefaultgohtml,
	"reports/html/timelog.gohtml": bindataReportsHtmlTimeloggohtml,
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
	"reports": {Func: nil, Children: map[string]*bintree{
		"html": {Func: nil, Children: map[string]*bintree{
			"commons.gohtml": {Func: bindataReportsHtmlCommonsgohtml, Children: map[string]*bintree{}},
			"default.gohtml": {Func: bindataReportsHtmlDefaultgohtml, Children: map[string]*bintree{}},
			"timelog.gohtml": {Func: bindataReportsHtmlTimeloggohtml, Children: map[string]*bintree{}},
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
