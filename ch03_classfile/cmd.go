package main

import (
	"flag" // 命令行工具包
	"fmt"  // 标准输入输出流包
	"os"
)

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
