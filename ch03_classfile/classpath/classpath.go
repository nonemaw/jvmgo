package classpath

import (
	"os"
	"path/filepath"
)

// 用户使用 -Xjre 选项配置启动类和扩展类路径，通过 -classpath/-cp 选项配置用户类路径
// ClassPath 结构体需包含全部三个字段
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAntExtClasspath(jreOption) // 解析 -Xjre 选项配置的 classpath
	cp.parseUserClasspath(cpOption)        // 解析 -cp 选项配置的用户 classpath
	return cp
}

// Classpath 的 ReadClass 方法按照 boot -> ext -> user 的顺序搜索提供的 class 文件名
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

func (self *Classpath) parseBootAntExtClasspath(jreOption string) {
	// 获取 jre 路径，为 bootClasspath 与 extClasspath 服务
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath) // 建立 bootClasspath
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath) // 建立 extClasspath
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "." // 如果用户未通过 -cp，则默认使用当前路径为 userclasspath
	}
	self.userClasspath = newEntry(cpOption)
}

// 根据配置值尝试建立 jre 路径，为 bootClasspath 与 extClasspath 服务
// 优先使用 -Xjre 选项配置的路径作为 classpath，若无则使用 JAVA_HOME
func getJreDir(jreOption string) string {
	// 如果输入的路径存在，则立刻返回
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 如果输入路径无效，则尝试在当前目录下寻找 jre 目录
	if exists("./jre") {
		return "./jre"
	}
	// 如果 jre 目录不存在，则尝试寻找环境变量
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Cannot find jre folder!")
}

// 判断一个目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
