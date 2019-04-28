package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry 是由一系列继承自 Entry 接口的结构体实例组成的数组
// 这里通过 type 定义了一个新的数据结构：Entry 数组
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{} // 先创建一个存储 Entry 接口类型的数组
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path) // 切割 pathList 并遍历每一个 path，通过 path 建立继承自 Entry 接口的结构体实例
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// CompositeEntry 结构体实现 Entry 接口 readClass() 方法
// 依次调用每一个子路径（ZipEntry/DirEntry）的 readClass() 方法
// 如果成功匹配到 className 则读取 class 数据，返回数据，如果收到错误信息，则 continue
// 如果遍历完所有的子路径还没有找到 class 文件，则返回错误
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
