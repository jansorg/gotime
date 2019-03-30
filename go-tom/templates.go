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
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x7b\x6f\xdb\x38\x12\xff\xbf\x9f\x82\xf0\x62\x81\xa6\x10\xe5\x3c" +
	"\xda\xa2\xd5\x66\x83\xeb\xa6\xed\x6d\x81\x0d\xee\x10\xef\x7e\x00\x4a\xa4\x25\x36\x7c\x08\xe4\x38\x8e\xd7\xf0\x77" +
	"\x3f\x50\x0f\x5b\xa2\x28\xd9\xc9\x2d\xda\x02\xb5\x44\x0e\x67\x7e\xbf\xe1\x3c\x28\x76\xbb\xa5\x6c\xc9\x15\x43\x33" +
	"\xa9\x15\xdb\xfc\xc1\x2d\xcc\x76\xbb\x57\x08\x21\xb4\xdd\x1a\xa2\x72\x86\xe2\xe6\xbd\x1e\x5b\x6a\x23\x09\xdc\x39" +
	"\x61\x37\x73\x9d\x9a\x9b\x46\x9a\x29\xba\xdb\xbd\x6a\x7f\x5f\x1d\x34\x83\x21\xd9\x03\xa3\xdf\xd4\x52\x1f\x74\xaf" +
	"\x39\x14\x9e\x6a\x8c\xe6\x6f\x72\x0d\x9b\x92\x25\x28\xe7\x50\xac\xd2\x38\xd3\x72\xfe\x9d\x28\xab\x4d\x3e\x07\x2d" +
	"\xe7\xb9\xc6\xee\x87\x12\x60\x7f\x72\xc9\x62\xf7\xcf\x17\x05\x66\xb3\x60\x86\x33\xfb\x66\x8e\x70\x4f\xe5\xfc\xcd" +
	"\x76\x2b\xb9\xfa\xbc\x32\x04\xb8\x56\x28\xbe\xe3\x6a\xb7\x7b\x33\x3f\x22\x44\x9e\x06\x42\x3d\x81\x4f\x8f\xf9\x9e" +
	"\xc8\x28\xed\xc5\x4a\x4a\x62\x36\x07\xca\x47\xf8\xe5\x1a\xb8\x64\xed\x8f\x61\xa5\x36\x10\xdf\x33\xbb\x12\xf0\xdb" +
	"\x2a\x7b\x60\xd0\xb0\xab\xb4\x5d\x03\x49\x05\x43\x99\x20\xd6\xfe\x3a\xb3\x8d\xa9\x9b\x3d\xe0\x6b\x48\x35\xdd\xdc" +
	"\x74\x08\xf0\x65\xfc\xcd\xfe\xd7\xe8\xef\x2c\x6b\xf4\x75\xf8\xd5\x4b\xcc\x4d\x6f\xa0\x1e\xa4\x37\xdb\x2d\xbf\xf8" +
	"\xa0\xd0\xac\x59\x3c\xdb\xed\xae\xe7\x40\x83\xb2\x2d\x20\xe0\x20\xd8\xec\x66\xbb\x8d\xff\x74\x4f\xa1\x15\xd7\xf3" +
	"\xae\xbd\xd6\x7f\x5d\xbc\x48\x69\x40\xb5\x1f\xfe\x53\x3a\xc7\xdb\xf8\xbe\xf6\xca\x67\x02\xec\x2b\x17\xc0\xcc\xbd" +
	"\x0b\xd1\xf8\x8b\x2c\x61\x83\x9e\xcb\xc7\x05\x0f\xaa\x62\x3c\x39\x8d\x93\x64\xb8\x12\x77\xc4\x4e\xc2\x75\xc7\x15" +
	"\x97\x44\x2c\xc0\x70\x95\x9f\xe4\x04\x61\x19\x6a\x98\x57\xda\x3a\xfc\x7e\x28\xbd\x8e\xf1\xe7\x93\xe8\xed\xe4\x00" +
	"\x67\x0f\x63\x5d\x1a\x90\x33\x1e\x46\xe9\x21\x74\xd8\x7a\xa9\xd8\x3e\xc5\xff\x76\x01\xdd\x5f\xef\x03\xe3\x4b\x2f" +
	"\x9c\x16\x85\x5e\x7f\x79\x22\x19\xb4\x5a\xec\xb3\x83\xa8\x5a\x8e\xe0\x28\x8f\xe7\x72\xa9\xf4\xbe\x30\x6f\x88\xa2" +
	"\x01\xa2\xad\xaf\xe3\xdf\x89\xfd\x4c\xb8\xd8\x34\x03\xcf\x65\x5c\xad\x7d\x29\xe3\x81\x44\x0d\x1a\x98\x2c\x05\x01" +
	"\xaf\x59\x20\xe7\x88\x09\xa8\x8d\x3b\xfe\x29\x07\xfd\xa5\x06\x2e\xda\x0f\xbd\xcc\x49\x2b\xf5\x63\xdd\x34\x06\xf7" +
	"\x25\x8e\xba\x9e\x77\x5a\xc8\xf5\xbc\xea\x37\x37\xa1\x3e\x97\x69\x29\xb5\xba\x5d\x2c\xda\x4e\x77\x6d\x61\xe3\x64" +
	"\x5b\x4d\x89\xd1\x1a\xd0\xb6\x67\x1b\xe3\xa5\x56\x80\x2d\xff\x9b\x25\xe8\xe2\xb2\x84\x5f\x42\xd3\x4b\x22\xb9\xd8" +
	"\x24\x68\x66\x37\x16\x98\xc4\x2b\x3e\x8b\x10\x26\x65\x29\x18\xae\x87\x22\xf4\x9b\xe0\xea\xe1\x8e\x64\x8b\xea\xfd" +
	"\xab\x56\x10\xa1\xd9\x82\xe5\x9a\xa1\xbf\xbe\xcd\x22\x74\xaf\x53\x0d\x3a\x42\xbf\x33\xf1\xc8\x80\x67\x24\x42\x9f" +
	"\x0c\x27\x22\x42\x96\x28\x8b\x2d\x33\x7c\x19\xa1\xd9\x27\xa7\x14\xdd\x6a\xa1\x0d\xfa\x22\xf5\x77\x67\x69\xaf\x26" +
	"\x30\xb2\xd8\xc8\x54\x8b\x99\x0f\xbb\xaa\xa2\x7d\xec\x7f\xac\x32\x4e\x09\xba\xd5\xca\x6a\xc1\x66\x11\xba\xd3\x8a" +
	"\x64\x3a\x42\x52\x2b\x6d\x4b\x92\xb1\x71\x25\x6b\xc6\xf3\x02\x12\xa4\xdc\xa9\x4b\xfc\xf2\xca\x13\xcc\x1c\xdc\x04" +
	"\xa5\x82\x64\x0f\xbe\x92\x34\x6f\xa7\xd7\x05\x87\x80\x0d\x10\xac\x95\xf8\xe9\xea\xe2\xfd\xbb\xf4\xad\x2f\x63\xf4" +
	"\x1a\x6b\x4a\xf7\x52\xac\xfa\xe3\x4b\x55\xb3\x78\xa5\x56\x96\xd1\x04\xfd\xf4\xe1\xa3\xfb\x3b\x40\xa3\x0d\x65\xa6" +
	"\x91\x15\x35\xa9\x47\x62\x5e\xf7\x97\x9f\x0d\x18\x56\x61\x87\xd7\x9c\x42\x91\xa0\x8b\xf3\xf3\x9f\x07\x3c\x2a\x01" +
	"\x4a\x80\x1c\x97\x82\xa2\xa5\xd2\x31\x7d\x36\x2a\x9b\xe6\x3d\xe9\xd6\xa1\xe3\x0b\xda\xdd\x4a\xb5\xa0\xa3\x42\x6e" +
	"\x5f\xeb\xb0\x3f\x8f\x3f\x18\x26\x07\x94\x9b\xc3\x5c\xcf\x65\x2d\x86\xde\x8e\x74\x80\xec\x0e\x4a\x0a\x90\xc2\x4b" +
	"\xb6\x4e\xaa\xd5\x6a\xf6\x03\x1e\x97\xa3\xde\x49\x49\xf6\x90\x1b\xbd\x52\x14\x9f\xe2\x9b\x5e\x22\x74\x4c\xd7\x43" +
	"\x61\xfc\xf5\xc1\xb6\x4f\xa0\xd9\xd9\x5a\x43\x27\x24\x7c\x74\x7b\x8f\x09\x52\x5a\x96\xa0\xf6\xa9\x67\xe8\x60\xa9" +
	"\x88\x10\x50\xcf\x94\xe0\x8a\xe1\xa2\xd9\xc8\x8b\xf8\xf2\x5d\xb5\x45\x5d\x89\x92\x50\xca\x55\xee\xf6\xaf\x9a\x45" +
	"\x17\x03\x91\x6e\x6d\x0b\x4f\xfa\x89\x1d\x72\x44\x11\x57\x49\xea\x40\xd6\x4f\x1e\xd4\x01\x90\xf3\x11\x3d\xbe\x37" +
	"\x5d\x3d\xc0\x55\xe5\x71\x00\xd6\x86\x94\x7d\x84\xc0\x9e\x00\x13\xc1\x73\x95\x20\xc1\x96\x5e\x6d\x7e\x64\xc6\x95" +
	"\x51\xd1\x4a\xa4\x1a\x40\xcb\xb0\x69\xdf\xbb\xfe\x5a\xd0\xe5\x08\x66\x46\xe8\x10\xf9\x58\xf8\xf9\x29\x3b\x15\xd6" +
	"\xfd\x52\x10\x0a\xd8\x6e\xa6\x0c\xd2\x36\xb4\xa0\xdd\x4c\x6f\x49\x3d\x1c\x8e\xf2\xf8\x50\xb7\x8e\xc7\xfa\xa1\xba" +
	"\x4d\x2a\xd3\x94\x22\x30\x89\x72\xdc\x0a\x2e\xe8\xeb\x4b\x85\x2f\xce\x10\xd0\xa8\x27\xe1\xba\x7a\x50\xee\x54\x6f" +
	"\x9f\x50\x84\x1a\x7b\x92\x80\xe1\x4f\x08\x68\xa2\x34\xbc\x4e\x04\xb1\x50\x9b\x3c\x8b\x7c\x91\x62\x20\x12\x8e\x76" +
	"\x6c\x9a\xd4\xec\xe5\x55\xd7\x74\xad\x11\x87\x32\x66\xb4\x43\x54\x11\xaf\x1f\x99\x59\x0a\xbd\x4e\x10\x13\x82\x97" +
	"\x96\xdb\xbe\x50\x3b\x8f\x9f\x12\x54\x70\x4a\x99\x9a\x4e\x0c\xc9\x29\x15\x6c\xc4\x3f\x85\xdb\xd0\x01\xbe\x7e\x66" +
	"\xf6\x4b\x43\x87\xc0\xc7\x8f\x3f\x87\xd5\x56\x96\xb1\xcb\x59\x4f\xf3\x78\x4a\x0f\x97\x57\x1e\x9e\x58\x5f\xcd\x8f" +
	"\xd0\x0a\x38\xbd\x9f\x7d\x87\x63\xc7\x54\x26\x0d\x3b\x68\xb7\xa4\xc6\x57\xa3\x9b\x9f\x56\x37\x1b\x16\xdd\x84\xb1" +
	"\xf4\xb4\xbc\x1f\xd5\xd2\xb4\xe0\x70\xf4\x90\x15\xe8\x50\xdf\x49\xd0\x55\xf9\x84\xac\x16\x9c\x36\x64\x43\x9d\xdc" +
	"\x63\x2d\x89\xc9\xb9\x4a\xd0\x39\x3a\x47\x57\xe3\x35\x7c\x8f\x68\x50\x51\xf7\x4d\x20\x7e\x37\x41\x27\x05\x0d\xc4" +
	"\x3f\x16\xd4\xb6\x31\xe8\x72\x2a\x9d\x5c\xb6\xef\x15\xb8\x96\xe9\x8d\xf8\x80\x1a\xb2\xb5\x56\xcf\x21\xc3\x53\xe0" +
	"\x48\xf9\xc8\xb4\xd8\x9b\x88\xfa\xaf\x61\x73\x2e\xa6\xff\x0f\x7b\xd5\xe5\xa7\xab\x49\x5c\x06\x43\xa6\x7f\x86\xf1" +
	"\x0f\xf9\x27\xf4\x04\xef\x44\x7f\x36\xde\x70\xbd\xec\x42\xc7\x1a\xf6\x80\x06\x76\xbd\x93\x99\x86\x4d\xf3\x76\x7a" +
	"\x3a\x1f\x35\x38\xa0\xd8\x4f\xd6\x7e\x39\x90\x6c\x6c\xe3\x4e\x28\x0b\x03\x55\x2f\xd4\xf3\x0c\xc4\x75\xfd\x08\x26" +
	"\x4a\x9d\xa4\x97\x53\x49\x0a\xc4\x00\xa6\xc4\x45\x52\xfd\xec\x50\x57\x2f\xba\xc4\x81\xd8\x92\x5c\xb5\x5f\x2d\xef" +
	"\x07\xe7\xc4\x93\x77\x5d\x69\x60\xd6\xff\xbe\x96\x16\x17\x9b\xb2\x60\xca\x86\x2a\x16\x5e\xb3\xf4\x81\xc3\x94\xc8" +
	"\xc8\x54\xd7\x30\x6d\xee\xab\xf0\xdf\xcc\xe8\xa9\x6d\xf1\xbe\xf1\x0e\xba\x50\x75\xb7\xd0\xdc\x17\xd4\x6f\xd5\x0b" +
	"\x92\x8c\x72\xf2\xeb\xac\x34\x5c\x41\xe7\x3a\xe4\x5f\x25\xc9\x07\x4e\x6c\x76\xe7\xf2\x5c\x86\x4b\x58\x75\xe0\x39" +
	"\xf9\xfa\xa1\xfb\x29\x75\x31\xd1\x3e\x3e\x8c\x2d\xba\x8c\x50\x71\x15\xa1\xe2\xed\xc4\xe2\xb7\x23\x8b\x27\x3e\x7e" +
	"\x86\x87\x16\xe7\x0b\x9c\x1a\x46\x1e\x30\x57\x96\x53\x96\x20\xf2\xa8\x79\x38\xb2\xc1\x2f\x01\xc7\x57\x7b\x52\x64" +
	"\x09\xae\xc9\x8d\x86\x43\x7d\x66\x1f\x35\x92\xb2\xa5\x36\x27\x1b\x99\x14\x9a\xc0\x4b\xb9\x2d\x05\xd9\x24\xb5\x27" +
	"\x9b\xe2\x87\xdd\x29\x76\xe4\x43\x23\x10\x1c\xff\x10\xe8\xae\x95\xe5\xf0\x06\xec\xd9\x56\x4e\x67\xed\x1a\xf4\x80" +
	"\x32\xea\xa6\x5a\xe8\x1a\x6f\x65\x41\xcb\xce\x35\x5e\xf3\x7f\x74\xfd\x7b\xd2\xdb\x56\xaa\x7b\x53\xe8\xdd\xf7\xd5" +
	"\x8b\xe3\xde\x5d\x62\x47\x64\x6f\xfb\xb8\x91\xaf\x5c\xb0\xa3\x86\xb8\x72\xdf\xec\xb7\x8b\x05\x3a\x6a\xb2\xf9\xfd" +
	"\x5f\x00\x00\x00\xff\xff\x61\x23\xc0\xb7\xf5\x1c\x00\x00")

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
		size: 7413,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1553936565, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataReportsHtmlDefaultgohtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5b\x8f\x9b\x38\x14\x7e\xcf\xaf\x38\x8b\xe6\x61\x5a\x2d\xa0\x79" +
	"\xab\x2a\x26\x0f\xed\x6c\x57\x95\xf6\x22\xcd\x4c\xf7\x61\xdf\x1c\xf0\x04\x6f\x01\x67\xb1\xe9\x34\x8b\xf2\xdf\x57" +
	"\xbe\x81\x01\x9b\x90\x4c\xaa\x36\x2f\x09\x60\x9f\xcb\x77\x2e\xfe\x0e\x69\xdb\x10\xe2\xd7\x5b\xca\xf7\x3b\xfc\x16" +
	"\xb6\x84\xe7\xcd\x26\x4a\x69\x19\xff\x83\x2a\x46\xeb\x6d\xcc\x69\x19\x6f\x69\x28\xbe\x6a\xbc\xa3\x35\x8f\xee\x31" +
	"\x6b\x0a\xfe\xae\x49\x3f\x63\xfe\x3a\x86\xf0\x70\x58\xad\xda\x36\xc3\x4f\xa4\xc2\x10\x3c\xa2\x4d\x81\xd5\xc3\xe0" +
	"\x70\x58\x01\x00\x5c\x46\x87\x92\x74\xb5\x91\xf7\xe0\xed\x2d\x44\xfd\x4d\x96\xd3\xe7\x5f\xbe\xa2\x54\xde\x57\x22" +
	"\xfe\xdc\x71\x42\x2b\x16\x3d\x98\x47\x77\x4d\x8d\xe4\x2d\x10\x06\x9f\x6a\x57\xce\xcb\x42\x09\x8e\xb5\xe4\xa1\x59" +
	"\x74\xc7\xd9\x44\x79\xa7\x29\xe1\x02\x15\x48\x0b\xc4\xd8\x6d\x20\x2f\x42\x9a\x65\xc1\x5a\x3e\x55\x2b\x72\x8c\x32" +
	"\xfb\xba\xee\x2f\xf4\x82\x6e\x7f\x1e\x3e\x93\x0c\x07\xeb\x84\xed\x50\xd5\xdd\x25\xbc\xc0\xc1\xba\x6d\xa3\x47\xf1" +
	"\xeb\x70\x48\x62\xf1\x78\x9d\xc4\x3c\x1f\x8a\x6a\x5b\xf2\x04\xd2\x62\x89\xce\x03\x2a\x30\xd3\x9e\x78\x34\x96\xb4" +
	"\xc2\xfb\x50\x58\x88\x6b\xa1\x82\xdc\xbc\xa9\x20\x90\x1b\x03\xa1\x88\xe7\x6b\x68\x5b\x5c\x65\x23\x31\x23\x4d\x9f" +
	"\xaa\xc7\x1a\xa5\x9f\xf1\x78\xd9\xd8\x3f\x52\xe2\x89\xb2\x3b\x44\x8a\x3d\x34\x55\xc8\x95\x08\xa3\x77\xa4\xef\xa8" +
	"\x0d\x2f\xb5\xe0\x54\xf5\x47\xa5\xea\xb4\xf4\x08\x14\xb6\x77\xd9\xed\x30\x7a\x61\x2c\x17\xc5\x53\x55\x10\xb3\xa3" +
	"\xea\xd0\x37\x75\x71\x89\x9b\x4a\x76\x36\xef\xec\x50\x74\x12\xdb\x35\x20\xd6\x8f\x2a\x64\x43\xb3\x7d\x7f\xdd\xb6" +
	"\x35\xaa\xb6\x18\x74\x87\x88\xde\xe7\xa4\xc8\x54\x07\x19\x23\x32\x29\x2e\x75\x33\x1b\x14\x0f\xcf\x5c\xde\x2f\xac" +
	"\x9c\x6c\x80\xb4\x80\x81\xe3\x72\x57\x20\x8e\x41\xdd\xfa\x8d\x30\x1e\x40\x24\x25\x44\xf7\xb4\xa9\x32\x91\x93\x3e" +
	"\xad\x6e\xcc\x4f\xaa\xae\xcc\x0e\xcf\xd0\x20\x9d\xd2\x1f\xab\x27\x1a\x40\xf4\x2b\xe6\x32\xd3\x2d\x71\x2f\x32\xeb" +
	"\xa2\x46\x9d\x69\x92\x43\x53\x49\x2a\x53\x7a\x10\x99\x5f\x42\xcf\x7c\xec\xe7\x6a\xd1\xe1\xfc\x7c\x3d\x9e\x96\x25" +
	"\x5a\xad\xb4\xce\xe7\xa8\xd3\x59\xe7\x2a\x65\xac\x17\x85\x39\x1f\x97\x63\x3f\x2c\x61\x47\x81\x0f\x4b\x38\xe1\x4f" +
	"\x94\xf2\xc1\x21\x68\x3c\xa9\xe9\x73\xc8\x9a\x0d\xa7\x1c\x15\xc1\xe4\x64\xec\xda\xcc\xa3\x7c\xee\x6f\xa5\xb3\x81" +
	"\x19\x37\x48\x18\xb7\x49\x67\x74\x4c\xbf\x99\x96\x72\xee\x0c\x93\xaf\x6a\x61\xf1\xc1\xe1\x29\x12\x63\x88\xbb\x80" +
	"\xcf\x3e\x29\x2f\x6e\xd7\x89\x56\x8d\xf4\xc1\x48\xa9\x9d\xc1\x46\xd3\xa4\x9c\xbf\xd5\xb1\x7a\x89\x8c\xe9\xcb\xda" +
	"\x9d\x2f\x17\xc0\xc0\x56\x71\xda\x91\xdb\xd7\x63\x12\x4b\xda\xba\x5e\x99\x3d\x16\xe7\xff\x1d\xf1\x9a\x7c\xfd\xe1" +
	"\x49\xbf\xbd\xaf\xdc\xf1\xfd\x74\xdf\xbd\x32\xe4\xc1\xac\x98\xe3\xef\x19\xe2\x08\x3a\x2a\xaf\x7f\x95\x12\x89\x33" +
	"\x79\xbd\xda\x1c\x6a\x22\x7f\x36\xb9\x7f\x26\x3c\xef\xd2\xe0\x03\xa9\x19\xff\x83\x56\xd2\x1d\x49\x89\x9c\xd9\x7e" +
	"\xa9\x58\x0d\xa5\x2a\x42\x36\xc7\xc4\x3c\x09\x6e\xe5\xb6\x45\xc8\x16\xd2\xd1\x73\x58\xb8\xfb\xdc\xf8\x76\x0c\x54" +
	"\xf6\x17\x5a\xc3\x35\xaa\x32\xb0\x32\x32\x92\x5f\xaf\xe0\xba\xa2\xdc\x5c\x38\xfb\x81\x83\xc0\xc2\x4c\x36\x81\xfc" +
	"\xba\x0d\x7a\x3c\xa7\xd8\x7a\xf8\xcc\xf2\x20\xc2\x49\xe4\xa3\x83\xa1\x6f\x55\x1f\x99\x3e\x3e\xff\xc6\x35\x9d\x51" +
	"\xd2\x29\xb3\x4b\xc4\xcc\x16\xe1\x7f\xb8\xa6\x47\x19\x9e\xac\x9e\xa3\x0a\x44\x61\x0c\x8e\x0a\x67\xa2\x3b\x4d\xdb" +
	"\x78\x22\xf4\x1d\x7c\xf0\x53\x45\x6b\x4d\xc1\xf0\x52\xc4\x7f\x04\x6c\xa5\x21\x8b\x56\x9b\x8f\x86\xb8\x6d\x41\x95" +
	"\xbc\x4e\xb5\x7e\x26\x05\x55\x14\xc7\x28\xf1\x85\x71\x9f\x5d\xe3\x66\xda\xf6\xce\x63\x5d\x20\xa5\x45\xc7\x9b\xa1" +
	"\xe7\x0e\x1d\x95\x76\x35\x86\x73\xe6\x85\x19\x27\x4e\x0f\x74\x17\xe0\x45\x41\xbb\x48\xd4\x66\x07\x2a\xe7\xd1\x33" +
	"\x3c\x1c\x5c\x32\xbe\xdd\xbc\xe3\x66\x0c\x4b\x86\x20\x7d\x42\x91\x9f\xaf\x52\xd1\xd0\x05\x05\xf2\x53\x85\x63\x4d" +
	"\xbf\x6d\xaf\x4c\x20\x84\x20\xd6\x94\x72\xc3\x5f\xa8\x68\x30\x33\x72\xe1\x8a\x2c\x65\xb4\x1e\x87\xbb\x3d\xba\xf9" +
	"\x5c\x0f\xf8\x6e\x66\x05\xfb\xd5\x7c\xa0\x4f\xcb\xc4\x17\x66\xe1\xc0\xb0\x05\x59\xe8\xe5\xfe\x2f\x78\xcd\x08\x9e" +
	"\xf2\xf7\xa1\xbc\x64\x9c\x72\xb2\xc7\x65\xa8\xbe\x14\xd1\xb9\xd1\xc6\x0d\xac\xe7\xa5\xc4\x1c\xc7\x5b\x36\xf2\x7c" +
	"\x97\x61\xe7\xc4\xa1\x25\x23\x5f\x0c\xd0\x4a\x66\x60\x73\x55\xf2\x04\x82\x7c\x4e\x07\x27\x35\xcd\xc9\xff\x71\x18" +
	"\x5c\x13\xa6\xae\xbb\x72\x96\xd4\xb4\xb7\xe7\xd5\x98\x9f\xda\x43\xef\x60\x30\xec\x24\x0c\x5b\x64\xc1\xb0\x48\x1f" +
	"\xfc\x6f\x1f\x5f\xbc\xe3\x39\xdc\x8c\x5f\x3c\xd8\x82\xed\x7f\x99\xcc\xbe\x89\x58\x17\xe7\x8e\x0c\x44\xd3\x4c\xd1" +
	"\x80\x5c\xcf\xa3\x6b\xb8\xf9\x1d\xe2\xf8\x5e\x34\x53\x3f\x4b\x1f\x07\x41\xb6\x69\x90\x1d\x58\xb2\xef\x4e\x84\xc8" +
	"\xdf\x8c\x7c\x59\x3c\xde\xb8\xb9\xda\x34\xde\x6c\xae\x91\x1e\x1b\x2d\x2f\x32\x0e\xd8\x31\x33\xe1\x8a\xbc\xaf\x4e" +
	"\x7d\x0d\x70\x82\xcd\xdc\x49\xab\x97\x77\x35\x9b\xfc\x94\xd1\x54\x14\x26\xe4\xbc\x2c\xd6\xab\x44\x7c\x41\x81\xaa" +
	"\xad\xec\x3f\xe2\xc7\x3b\xc4\xb0\x64\x7f\xab\xa4\x9f\xf0\x92\x12\x73\x04\x69\x8e\x6a\x86\xf9\x6d\xf0\xe9\xf1\x43" +
	"\xf8\x46\xe3\xa9\x67\xec\x61\xa2\xbc\x6f\x18\xa7\xa5\xc6\xd0\x3a\xd2\xc5\x8d\x75\xdb\x42\x04\xf2\x44\x96\x97\x2b" +
	"\x18\x53\x38\x1b\xa8\x94\x96\x25\xad\xde\x3f\x3c\xf4\xed\xc5\x7a\x28\xd5\xe8\x87\x49\xac\xec\x4d\x14\xaf\x58\x5d" +
	"\xe4\x1f\xcd\xd9\x7f\x33\xcd\xeb\x05\xf9\x6a\xcc\x72\xd9\x14\x6b\x92\xdf\x88\x54\x12\xae\xe6\x37\x83\xce\x39\xd9" +
	"\x77\x87\x59\x5a\x13\x29\xbb\xdb\xbd\xeb\x26\xa1\xfe\x61\x60\x04\xee\x06\xf2\x86\x2f\xe8\x9a\xb2\x44\xf5\x1e\x1c" +
	"\x80\xe9\x47\x01\xe8\x96\x2b\x3d\x54\x52\x9c\xd9\xd9\xad\x4a\x62\x85\x6a\x22\xa1\x5a\xff\x1f\x00\x00\xff\xff\xf2" +
	"\x7b\x91\x24\x2a\x1f\x00\x00")

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
		size: 7978,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1550089362, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataReportsHtmlTimeloggohtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xcd\x8e\xe3\x36\x0c\xbe\xfb\x29\x58\x63\x0f\xed\x00\xb6\x31\x28" +
	"\x0a\x2c\x16\x4e\x0e\xf3\x57\xf4\xd2\x16\x9b\xf4\x01\x14\x8b\x63\xab\x6b\x4b\x86\xc4\xec\x34\x30\xfc\xee\x85\x24" +
	"\xff\x24\x19\xdb\x93\x64\x73\x88\x2c\x91\xfc\xf8\x2f\xb1\x69\x22\x48\xee\x72\x45\x87\x1a\xbf\x40\x2e\xa8\xd8\xef" +
	"\xe2\x4c\x55\xc9\xbf\x4c\x1a\xa5\xf3\x24\x57\x24\x2a\xec\x17\x8d\xb5\xd2\x14\x3f\xec\xb3\x6f\x48\x5f\xdd\xe6\x2e" +
	"\x81\xa8\x6d\x83\xa0\x69\x38\xbe\x0a\x89\x10\x7a\x6a\xd8\xb6\x01\x00\xc0\x87\x1a\x48\x55\x49\xae\x22\xbb\x74\xf0" +
	"\x5f\xd1\xec\x4b\xf2\x30\x1d\xbc\x47\xfa\xb4\x73\x67\xf0\x65\x05\xf1\x78\x68\x0a\xf5\xf6\x5c\xd5\x74\xb0\xe7\x1e" +
	"\xe2\xaf\x9a\x84\x92\x26\xf6\x26\xc6\x9b\x9e\xc3\x1a\x6a\xa5\x52\x2e\xbe\x43\x56\x32\x63\x56\xa1\xc7\x0c\xd7\x8e" +
	"\xe0\x21\xc5\x2b\x74\xaa\xe2\x5e\x0c\x8e\x7e\x8e\x41\x69\xf8\x59\x2a\x3a\x65\xdc\xa8\xbd\xce\xf0\x17\x18\x6d\x3a" +
	"\x13\x75\xca\x89\xed\x4a\xec\xd5\xbb\x4d\xc4\x19\x31\xf0\x9f\x8a\xf3\x23\x63\x4e\x05\x0b\x64\x7c\x8e\xa6\xa7\x09" +
	"\xde\xe0\xe4\x0e\xab\x1d\x72\x8e\x1c\x48\x50\x89\x40\x0a\xbe\x21\xd6\xdd\x8e\x49\xee\xb5\x1b\x20\x95\x23\x15\xa8" +
	"\x41\x48\xa8\xb5\x90\x84\x1c\xfe\x7e\x7a\x31\x77\xc9\x84\x2b\x47\x96\x41\xa6\x4a\x53\x33\xb9\x0a\x7f\x0b\x07\xdf" +
	"\x2c\x78\xb8\x1e\x12\x17\x6f\xed\x41\xdb\x8e\xa9\x8c\x37\x75\x29\xe8\xe1\xb0\x3d\xd4\xd8\xb6\x69\x42\xc5\x8c\x7b" +
	"\xc9\x9c\x7f\x8b\x8e\x3b\xbb\xbc\x2d\x9c\x11\x46\x36\x7e\xa8\xad\x45\xe2\xfe\xb3\x84\xf0\x89\x11\x86\x4b\x7a\x2f" +
	"\x01\xd9\x10\xd3\xf4\xc3\x28\xcf\x92\x5f\x83\x61\xdb\xf1\xbd\x3b\x7b\xcd\x6c\xe1\x5f\x05\x54\x44\x6f\x82\xe3\x08" +
	"\xf2\xa7\x22\x34\x8b\x08\x0b\xc9\x48\x16\x4b\x74\xa7\xf8\x61\x9a\xd6\x34\x9a\xc9\x1c\x87\x6e\x7a\xd1\xac\x42\xd3" +
	"\x2d\x8b\x85\xb7\x90\x7d\xcf\xc0\x7b\x4f\x8d\xcd\x53\xc4\xd9\xc1\xfa\xfa\xaa\x74\xc5\xc8\xe6\x1f\x62\x97\x40\xe7" +
	"\xf0\x8c\xe5\xb3\x60\x36\x0b\x23\xda\x56\x54\x37\xa3\xa9\xba\x03\x5b\x14\x81\xfe\xf6\x89\xff\x30\x1b\x52\x75\x8d" +
	"\x7c\x21\x38\x93\x52\x42\xe6\x25\x3e\xb1\xa9\x8b\x69\x5e\xf6\xd8\xbf\x67\x79\xb9\x4e\x2c\x0d\xde\xa0\xc8\xa6\xe5" +
	"\x06\x65\x17\xf1\x5e\xc2\x77\x55\xee\xfa\x1a\xa8\x84\xec\xfb\x0f\xe2\xfe\xeb\xca\x3a\x90\xae\xf7\xd6\x4d\x13\xbb" +
	"\x2e\xfc\x48\x7a\xbe\x11\x97\xbc\x4c\x93\x85\x46\x4c\xe9\x55\x29\x1a\xea\x72\xbf\x23\x45\xac\x9c\x7d\x90\x74\xcf" +
	"\xa9\xd5\x5b\xf4\x01\x37\x9c\x3f\x14\xbf\x8e\x97\xce\xd6\xc9\x5d\x71\x6d\x65\xaa\x1c\xf4\x81\x4d\x81\xfb\x3b\x32" +
	"\x61\x26\x21\xf1\xef\x48\xcf\xff\xb1\x8c\x2e\xd0\xb5\xbe\xf1\x0a\xb4\x11\x7c\x4f\x4c\x13\xf7\xc2\xae\xcf\x46\x89" +
	"\xd3\x3c\x4d\xb6\xcc\xf1\xb4\x32\xf7\xa4\xa6\x09\x17\xdf\xcf\xb1\x4f\xef\xd4\xc7\x42\x94\xdc\x4f\x55\x53\x57\x6a" +
	"\xd3\x10\x56\x75\x69\x6f\xc4\x7e\x84\x1b\xa6\xac\x05\x73\x87\x7d\x67\x41\x7f\x14\x5c\x33\xf8\x15\x54\x95\x7e\x72" +
	"\x4b\xba\xd1\xad\x9b\xfb\x9a\xe6\x93\xaa\xc9\xbc\x9b\xec\xc0\x6a\x48\x7f\xe2\x2a\xb3\xf0\x60\xe5\xd7\x41\x6a\x17" +
	"\x28\x99\xcc\x57\x61\xd3\xb8\x8f\x07\x66\x10\xda\x36\xb4\xc4\xe1\x61\x4a\x2b\x24\x06\x59\xc1\xb4\x41\x5a\x85\xff" +
	"\x6c\x5f\xa2\xcf\x5d\xcd\x36\xcd\x9b\xa0\x02\x9c\xd2\xf8\x71\x6f\x48\x55\x5d\x80\x83\xb1\x30\xec\x81\xed\x51\x57" +
	"\x42\x6e\x13\x1c\x07\x23\x38\x8f\x66\xa6\xaa\x4a\xc9\xc7\xcd\x66\x9c\x89\x8f\x88\x4e\x49\x47\x4c\x13\x6f\x65\xea" +
	"\x1b\x34\x98\x33\x07\xfa\x98\x17\xf7\xbd\x21\xc5\xfd\x49\xec\xdf\xc9\x3d\xa1\xc9\xb4\x70\xd1\x1b\xa4\xeb\x61\x22" +
	"\x19\x89\x61\x0f\x58\x9f\xe0\xd9\x91\xd8\xa1\xd9\x59\x7a\xb3\xaf\x2a\xa6\x0f\x30\xe1\x4e\x47\x0a\xa1\x9b\xe2\x5d" +
	"\x0e\x07\x94\x89\x0a\x1b\xd8\xd2\xc4\x3b\x9d\xba\x6a\x58\xff\x1f\x00\x00\xff\xff\x42\x08\x63\x62\x9a\x0c\x00\x00" +
	"")

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
		size: 3226,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1549139854, 0),
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
