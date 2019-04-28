package classpath

import (
	"os"
	"strings"
)

// 可以将类路径想象成一个大整体，由启动类路径+扩展类路径+用户类路径三个模块构成
// 这里就可以使用组合模式来设计类路径

// `string(os.PathListSeparator)` 既可自动获得系统分隔符
// （分隔符因系统而定，Win 为 `;`，类 UNIX 为 `:`）
const pathListSeparator = string(os.PathListSeparator)

// Entry 是一个接口，包含两个方法
type Entry interface {
	// 负责寻找和加载 .class 文件（相对路径），返回字节数据、Entry 实例和错误信息
	// golang 和 Python 类似，可以同时返回多个返回值
	readClass(className string) ([]byte, Entry, error) // 根据提供的 className 读取 class 字节码
	String() string                                    // 类似于 Java 的 toString() 作用
}

// 根据参数创建不同类型的 Entry 接口实例
// Entry 接口共有 4 个实现方式，分别是 DirEntry、ZipEntry、CompositeEntry 和 WildcardEntry
func newEntry(path string) Entry {
	// 若包含系统分隔符（即加载多个类和目录），则返回 CompositeEntry 实例
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	// 若包含 `*`（即加载目录下所有 jar 文件），则返回 WildcardEntry 实例
	if strings.Contains(path, "*") {
		return newWildcardEntry(path)
	}
	// 若包含 jar/zip 文件名，则返回 ZipEntry 实例
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	// 加载目录，返回 DirEntry
	return newDirEntry(path)
}
