package main

import (
	"flag" // 命令行工具包
	"fmt"  // 标准输入输出流包
	"os"
)

// 在这里我们需要实现三种 java classpath 类路径（Oracle JVM）：
// bootstrap classpath 路径为 jre/lib
// extension classpath 路径为 jre/lib/ext
// user classpath      路径自定义
// 在需要的情况下，可以通过 JVM -Xbootclasspath 选项来修改启动与扩展路径
//
// user classpath 默认为当前路径，CLASSPATH 变量值，或启动选项 -classpath/-cp 直接指定
// -classpath/-cp 选项可以指定一个或多个目录/zip/jar 文件，或通过 `*` 加载目录下所有 jar 文件，如：
// `java -cp path\to\lib1.jar`
// `java -cp path\to\classes;lib\a.jar;lib\b.jar;lib\c.zip`
// `java -cp classes;lib\*`
// （分隔符因系统而定，Win 为 `;`，类 UNIX 为 `:`）

// 在此首先需要让 Cmd 能够加载标准库的类，因此需要通过某种方式来指定 jre 目录未知
// 故我们新增一个自定义选项：-Xjre，并在 Cmd 中添加该字段
type Cmd struct {
	helpFlag    bool
	versionFlag bool

	cpOption   string
	XjreOption string // -Xjre 选项

	class string   // java 主类名
	args  []string // 主类参数
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage

	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")           // -help
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")              // -?
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit") // -version

	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath") // -classpath
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")        // -cp
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")  // -Xjre

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0] // 第一个参数为主类名
		cmd.args = args[1:] // 随后为主类的参数
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
