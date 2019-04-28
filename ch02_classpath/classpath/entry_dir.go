package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry 结构体，只有一个字段，用于存放 classpath 绝对路径
type DirEntry struct {
	absDir string
}

// 构建 DirEntry 结构体实例
// golang 没有构造函数，所以这里可以统一用 newXxxx() 方法来创建结构体实例
//
// golang 的异常
// golang 中没有 try... catch... 的异常处理方式，通常直接使用多值返回 err 来进行异常处理
// （如下方 readClass() 方法的 err 值）
//
// 只有在非不得已的情况下才需要异常处理：defer, panic, recover
// 简单来说即为：golang 中抛出一个 panic 异常，然后在 defer 中通过 recover 捕获这个异常进行后续处理
// http://www.cnblogs.com/ghj1976/archive/2013/02/11/2910114.html
//
// golang 接口的实现
// 和 Java/Python 等语言不同, golang 的结构体在实现接口时无需显式构建，只要方法匹配即可
// 只要一个结构体含有接口类型中的所有方法，那么这个结构体就实现这个接口
// 而如果它只实现了某些个方法，则这个结构体没有实现这个接口
// 如果接口没有实现，则这些定义的方法会成为结构体的方法（即方法与结构体实例绑定）
//
// golang 的函数与方法
// 由于 golang 并不是面向对象的语言，故它无法实现类。但和 C 一样，它可以实现方法（即将函数与一个实例绑定）
//
// golang 方法的实现
// 通过将一个函数与实例的绑定，我们可以创建 golang 中的方法，其语法为：
// `func (t T) funcName() [return type] { ... }`
// 其中 "(t T)" 为该函数的【接收者】，即函数的绑定者，它可以是一个值 "(t T)" 或一个指针 "(t *T)"
//
// ##### golang 值接收者 Value Receiver #####
// 声明：func (u User) notify() { ... }
// 调用：
//
// ##### golang 指针接收者 Value Receiver #####
// 声明：func (u *User) notify() { ... }
// 调用：
//

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
