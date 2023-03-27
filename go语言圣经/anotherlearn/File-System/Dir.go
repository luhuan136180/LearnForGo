package File_System

import (
	"errors"
	"io/fs"
	"time"
)

type dir struct {
	name    string
	modTime time.Time
	// 存放该目录下的子项，value 可能是 *dir 或 *file
	children map[string]fs.DirEntry
	// ReadDir 遍历用
	idx int
}

// dir 虽然是一个目录，但根据一切皆文件的思想，目录也是文件，因此需要实现 fs.File 接口
// 这样，fs.FS 的 Open 方法可以对目录起作用。
func (d *dir) Read(p []byte) (int, error) {
	return 0, &fs.PathError{
		Op:   "read",
		Path: d.name,
		Err:  errors.New("is directory"),
	}
}

func (d *dir) Stat() (fs.FileInfo, error) {
	return d, nil
}
func (d *dir) Close() error {
	return nil
}

// ReadDir 实现 fs.ReadDirFile 接口，方便遍历目录
func (d *dir) ReadDir(n int) ([]fs.DirEntry, error) {
	names := make([]string, 0, len(d.children))
	for name := range d.children {
		names = append(names, name)
	}

	totalEntry := len(names)
	if n <= 0 {
		n = totalEntry
	}

	dirEntries := make([]fs.DirEntry, 0, n)
	for i := d.idx; i < n && i < totalEntry; i++ {
		name := names[i]
		child := d.children[name]

		f, isFile := child.(*file)
		if isFile {
			dirEntries = append(dirEntries, f)
		} else {
			dirEntry := child.(*dir)
			dirEntries = append(dirEntries, dirEntry)
		}
		d.idx = i
	}
	return dirEntries, nil
}

// 因为 fs.Stat 对目录也是有效的，因此 dir 需要实现 fs.FileInfo 接口
func (d *dir) Name() string {
	return d.name
}

func (d *dir) Size() int64 {
	return 0
}

func (d *dir) Mode() fs.FileMode {
	return fs.ModeDir | 0444
}

func (d *dir) ModTime() time.Time {
	return d.modTime
}
func (d *dir) IsDir() bool {
	return true
}
func (d *dir) Sys() interface{} {
	return nil
}

// 因为 dir 是一个目录项，因此需要实现 fs.DirEntry 接口
func (d *dir) Type() fs.FileMode {
	return d.Mode()

}

func (d *dir) Info() (fs.FileInfo, error) {
	return d, nil
}
