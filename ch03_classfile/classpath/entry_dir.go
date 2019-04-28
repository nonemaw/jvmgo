package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry 结构体，只有一个字段，用于存放 classpath 绝对路径
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	dir, err := filepath.Abs(path) // 将相对路径转换为绝对路径
	if err != nil {                // 通过多值返回捕获可能的异常
		panic(err) // 有异常则进行 panic() 中断执行
	}
	return &DirEntry{absDir: dir}
}

// DirEntry 结构体实现 Entry 接口 readClass() 方法
// 根据 className 与提供的 dir 信息，读取 class 文件并返回文件数据，结构体实例和错误信息
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// DirEntry 结构体实现 Entry 接口 String() 方法
// 至此结构体 DirEntry 已经实现了 Entry 接口的所有方法，DirEntry 成为了 Entry 接口的实现
func (self *DirEntry) String() string {
	return self.absDir
}
