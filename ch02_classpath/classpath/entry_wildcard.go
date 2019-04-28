package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// WildcardEntry 结构体没有实现接口 Entry
// WildcardEntry 实际就是 CompositeEntry，所以只需实现 newWildcardEntry() 并返回一个 CompositeEntry 即可
// WildcardEntry 的文件列表是根据通配符进行自动遍历的
// CompositeEntry 是手动指定多个 jar/zip 文件
//
// 对于带有通配符 `*` 的路径，首先需要去除末尾星号，然后通过 filepath.Walk() 对目录遍历
// filepath.Walk() 方法支持自定义遍历方法
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // 去除 `*`
	compositeEntry := []Entry{}

	// 自定义遍历方法：寻找 jar 文件包。自定义遍历方法的定义与参数为：
	// `type WalkFunc func(path string, info os.FileInfo, err error) error`
	findClassFiles := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir // 如果当前遍历文件为目录则跳过，因为通配符路径不能递归
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path) // 如果当前文件为 jar 文件，则为其建立 ZipEntry
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	// 通过自定义遍历方法对目录进行遍历
	filepath.Walk(baseDir, findClassFiles)
	return compositeEntry
}
