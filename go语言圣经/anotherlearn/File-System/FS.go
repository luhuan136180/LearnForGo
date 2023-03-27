package File_System

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

type FS struct {
	rootDir *dir
}

func NewFS() *FS {
	return &FS{
		rootDir: &dir{
			children: make(map[string]fs.DirEntry),
		},
	}
}

func (fsys *FS) Open(name string) (fs.File, error) {
	//效验 name
	if !fs.ValidPath(name) {
		return nil, &fs.PathError{
			Op:   "open",
			Path: name,
			Err:  fs.ErrInvalid,
		}
	}

	//根目录处理
	if name == "." || name == "" {
		// 重置目录的遍历
		fsys.rootDir.idx = 0
		return fsys.rootDir, nil
	}

	// 3、根据 name 在目录树中进行查找
	cur := fsys.rootDir
	parts := strings.Split(name, "/")
	for i, part := range parts {
		// 不存在返回错误
		child := cur.children[part]
		if child == nil {
			return nil, &fs.PathError{
				Op:   "open",
				Path: name,
				Err:  fs.ErrNotExist,
			}
		}
		//判断是否为文件
		f, ok := child.(*file)
		if ok {
			//文件名是最后一项
			if i == len(parts)-1 {
				return f, nil
			}
			return nil, &fs.PathError{
				Op:   "open",
				Path: name,
				Err:  fs.ErrNotExist,
			}
		}
		// 是否是目录
		d, ok := child.(*dir)
		if !ok {
			return nil, &fs.PathError{
				Op:   "open",
				Path: name,
				Err:  errors.New("not a directory"),
			}
		}
		// 重置，避免遍历问题
		d.idx = 0
		cur = d

	}
	return cur, nil
}

// MkdirAll 这不是 io/fs 的要求，但一个文件系统目录树需要可以构建
// 这个方法就是用来创建目录
func (fsys *FS) MkdirAll(path string) error {
	if !fs.ValidPath(path) {
		return errors.New("Invalid path")
	}
	if path == "." {
		return nil
	}

	cur := fsys.rootDir
	parts := strings.Split(path, "/")
	for _, part := range parts {
		child := cur.children[part]
		if child == nil {
			childDir := &dir{
				name:     part,
				modTime:  time.Now(),
				children: make(map[string]fs.DirEntry),
			}
			cur.children[part] = childDir
			cur = childDir
		} else {
			childDir, ok := child.(*dir)
			if !ok {
				return fmt.Errorf("%s is not directory", part)
			}
			cur = childDir
		}
	}
	return nil
}

//WriteFile 方法就是生成一个 file 然后存入 files 中。
func (fsys *FS) WriteFile(name, content string) error {
	if !fs.ValidPath(name) {
		return &fs.PathError{
			Op:   "write",
			Path: name,
			Err:  fs.ErrInvalid,
		}
	}

	var err error
	dir := fsys.rootDir

	path := filepath.Dir(name)
	if path != "." {
		dir, err = fsys.getDir(path)
		if err != nil {
			return err
		}
	}
	filename := filepath.Base(name)

	dir.children[filename] = &file{
		name:    filename,
		content: bytes.NewBufferString(content),
		modTime: time.Now(),
	}
	return nil
}

// getDir 通过一个路径获取其 dir 类型实例
func (fsys *FS) getDir(path string) (*dir, error) {
	parts := strings.Split(path, "/")

	cur := fsys.rootDir
	for _, part := range parts {
		child := cur.children[part]
		if child == nil {
			return nil, fmt.Errorf("%s is not exist", path)
		}

		childDir, ok := child.(*dir)
		if !ok {
			return nil, fmt.Errorf("%s is not directory", path)
		}
		cur = childDir
	}
	return cur, nil
}
