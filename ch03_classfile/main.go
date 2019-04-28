package main

import (
	"fmt"
	"jvmgo/ch03_classfile/classpath"
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
