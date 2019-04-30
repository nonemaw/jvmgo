package main

import "flag" // 命令行工具包
import "fmt"  // 标准输入输出流包
import "os"

// cmd.go 用于提供基础的 java 命令行功能，它可以接收多个常用标准参数选项
// java 命令应当能够接收三个参数：选项，main 方法类名（或 jar 文件名），main 方法参数（JVM 无法搜索主类，需要手动指定）
// java [-options] class [args]
// java [-options] -jar jarfile [args]
// javaw [-options] class [args]
// javaw [-options] -jar jarfile [args]
//
// 而选项本身又分为两类：标准选项和非标准选项，常见选项有 -version 和 -help
//
// 这里定义了命令行选项结构体
type Cmd struct {
	helpFlag    bool
	versionFlag bool

	cpOption string
	class    string   // java 主类名
	args     []string // 主类参数
}

// 关于 golang 的访问控制
// golang 中的访问控制只有 public 和 private 两种：
// 所有首字母大写的类型、结构体、变量、函数等都是 public 的，可以在任意包 import
// 所有首字母小写的都是 private，仅能在在包内使用

// 关于变量赋值，go 解释器允许隐藏类型声明，解释器会自动基于值的类型隐式为变量声明类型
// ":=" 被称为 “短式变量声明 Short variable declarations”
// a := "b" 等同于 var a string = "b" 或 var a ="b"

// 关于指针
// var p *int 为声明一个指向 int 类型的指针，此时 p 值为 `<nil>`（golang 的 null），若引用 *p 则会报错因为它是空指针
// "&" 术语 “dereferencing”，用于获取另一个变量的地址，如 i := 1, p := &i，此时 p 指向了 i 的地址
// "*" 术语 “indirecting”，用于获取指针指向的地址的值，如 fmt.Println(*p) 即会输出 1，又如 *p = 9，则此时 i 的值也是 9

// 关于创建结构体变量与指针
// cmd := &Cmd{} 完全展开等同于 var struct_cmd Cmd = Cmd{}, var cmd *Cmd= &struct_cmd
// 创建结构体变量：可以直接通过 a := Cmd{ helpFlag: true } 来创建一个结构体的变量，并顺带初始化其某些值
// 创建结构体指针：可以通过 &Cmd{} 或 `new` 关键字实现，new(Cmd) 会返回一个指向当前结构体/接口的指针
//     即：var cmd_p *Cmd = new(Cmd)

// 关于结构体指针
// cmd := &Cmd{}
// fmt.Println(cmd)  输出结构体的引用
// fmt.Println(*cmd) 输出结构体的本身值
// fmt.Println(&cmd) 输出指针变量本身的地址

// 关于通过结构体指针来引用值，假设我要调用结构体变量 helpFlag
// (*cmd).helpFlag   标准调用方式，即获得结构体本身然后调用其 helpFlag
// cmd.helpFlag      精简调用方式，等同于 (*cmd).helpFlag，解释器会自动补全
// &(cmd.helpFlag)   获结构体的 helpFlag 变量的地址（即匿名指针），等同于 &((*cmd)cmd.helpFlag)
// &cmd.helpFlag     等同于 &(cmd.helpFlag)，"." 的优先级大于 "&"
func parseCmd() *Cmd {
	cmd := &Cmd{}
	// go 可以直接将函数赋值给变量，这里的方法用于打印使用提示
	flag.Usage = printUsage
	// 设置 flag 的各种 Var() 方法用于选项的解析
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")           // -help
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")              // -?
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit") // -version
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")                // -classpath
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")                       // -cp
	// 解析输入选项，如果解析失败（即未知的 "-" 选项）则执行 flag.Usage 存储的方法
	flag.Parse()

	// 如果解析成功，则开始处理输入选项后的 args 参数，通过 flag.Args() 来捕获选项
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
