package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: absPath}
}

// ZipEntry 结构体实现 Entry 接口 readClass() 方法
// 从 zip 文件进行遍历并提取与 class Name 同名的 class 文件
// 这里可以看到，目前每一次寻找 class 文件时都需要遍历，这里可以进行优化
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath) // 尝试打开 zip 文件，如果出错则直接返回
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	for _, f := range r.File { // 若 zip 文件打开成功，则遍历并寻找 class 文件
		if f.Name == className {
			rc, err := f.Open() // 若文件名为 className 则尝试打开当前遍历的文件，若打开失败则直接返回
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc) // 若文件打开成功，则尝试读取文件内容
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}
	return nil, nil, errors.New(" class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
