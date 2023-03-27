package File_System

import (
	"bytes"
	"errors"
	"io/fs"
	"time"
)

type file struct {
	name string
	//存放文件内容
	content *bytes.Buffer
	modTime time.Time
	closed  bool
}

func (f *file) Read(p []byte) (int, error) {
	if f.closed {
		return 0, errors.New("file closed")
	}
	return f.content.Read(p)
}

func (f *file) Stat() (fs.FileInfo, error) {
	if f.closed {
		return nil, errors.New("file closed")
	}
	return f, nil
}

// Close 关闭文件，可以调用多次。
func (f *file) Close() error {
	f.closed = true
	return nil
}

// 实现 fs.FileInfo
func (f *file) Name() string {
	return f.name
}

//
func (f *file) Size() int64 {
	return int64(f.content.Len())
}

func (f *file) Mode() fs.FileMode {
	return 0444
}

func (f *file) ModTime() time.Time {
	return f.modTime
}

func (f *file) IsDir() bool {
	return false
}

func (f *file) Sys() interface{} {
	return nil
}

// 文件也是某个目录下的目录项，因此需要实现 fs.DirEntry 接口
func (f *file) Type() fs.FileMode {
	return f.Mode()
}

func (f *file) Info() (fs.FileInfo, error) {
	return f, nil
}
