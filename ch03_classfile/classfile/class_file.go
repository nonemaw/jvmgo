package classfile

import (
	"fmt"
)

// ClassFile 结构体反映了 JVM 规范定义的 class 文件格式信息
type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse() 函数把读取的 class 文件字节数据流解析为 ClassFile 结构体
// 这里使用了 defer - panic - recover 来预防异常，具体可以看：https://www.jianshu.com/p/f76b9ce083c4
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// read() 方法绑定至了 ClassFile 结构体，是解析 class 文件的入口方法
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// 下面几个是类似 getter 的方法，绑定至了 ClassFile 结构体用于让其它包共享数据
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fileds() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

// 魔法数字：JVM 规定某些文件（如 class 文件）必须以固定字节开头
// 0xCAFEBABE 是所有 class 文件的开头字节。当 JVM 遇到非法的 class 开头字节时会抛出 java.lang.ClassFormatError 异常
// 这里先不做错误处理，只用 panic 抛出异常信息
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 魔法数字是文件开头，之后便是版本号
// 版本号：class 文件都有一个主版本号 M 和次版本号 m，都是双字节 uint16 类型，完整版本号为 M.m
// 目前次版本号已经不再使用，都为 0
// 主版本号从 Java1 的 45 开始，在每一个 Java 版本发布时都会 +1，故 Java8 版本号为 52（0x34）
// 通常情况下 JVM 能够向后兼容旧版本的 class，如果版本号不能支持则会抛出 java.lang.UnsupportedClassVersionError 异常
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// 版本号之后便是常量池
// 常量池：这里先不讲
// 常量池之后是 class 访问标志
// 访问标志：一个 16 位的 bitmask，用于标明这个 class 文件是类还是接口，以它的权限如 public / private 等
// 这里先不去关心它的完整信息，只做初步解析

// 访问标志之后便是两个 uint16 类型的常量池索引
// 常量池索引：用于指明当前类名 thisClass 和父类名 superClass。class 文件会完整存储完整类名，只是将 "." 换成了 "/"
// 除了 java.lang.Object 外，其它所有 Java 类都有父类，故只有 Object 的 superClass 是 0
// 其它所有的 class 必须有一个合法的 thisClass 和 superClass 常量池索引
// 从常量池中查找继承的接口名
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

// 当前类和父类索引之后是接口索引，其中保存的也是常量池索引，大小为 uint16
// 从常量池中查找继承的父类名
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

// 接口索引之后便是字段表和方法表，分别存储字段和方法信息
// 字段和方法的基本结构大致相同，差别仅在于属性表，下面是一个 JVM 标准字段结构定义：
// field_into {
//     u2              access_flags;
// 	   u2              name_index;
//     u2              descriptor_index;
// 	   u2              attributes_count;
// 	   attribute_info  attributes[attributes_count];
// }
// 参考 member_info.go 代码
