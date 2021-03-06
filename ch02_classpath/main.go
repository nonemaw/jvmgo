package main

import (
	"fmt"
	"jvmgo/ch02_classpath/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

// 使用范例
// ./ch02_classpath.exe -Xjre "D:\Java\jdk1.8.0_171\jre" java.lang.Object
// 用户通过 -Xjre 配置了 boot 和 ext classpath
// cpOption 为空，未配置，则 user classpath 默认为当前路径
// 主类文件名为 java.lang.Object
// 文件内容由 ZipEntry 解析，路径为 D:\Java\jdk1.8.0_171\jre\lib\rt.jar -> java/lang/Object.class
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption) // 制作 classpath
	fmt.Printf("classpath:%v mainclass:%v args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1) // 根据主类名制作 main class 路径
	classData, _, err := cp.ReadClass(className)          // 读取主类文件
	if err != nil {
		fmt.Printf("Cannot find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
