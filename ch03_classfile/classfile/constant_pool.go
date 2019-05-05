package classfile

// 常量池占据了 class 文件的很大一部分，里面存放着各种常量信息，包括数字常量，字符串常量，
// 类名，接口名，字段，方法等等
//
// 常量池实际也是一个表，有如下特点：
// 1. 表头给出的常量池大小会比实际大 1，为 n - 1
// 2. 常量池的有效索引范围是 0 ~ n - 1，0 是无效索引
// 3. CONSTANT_Long_info 和 CONSTANT_Double_info 各占两个字节，也就是说如果常量池存在这两种变量，
// 则常量池的大小会比 n - 1 还要小
type ConstantPool []ConstantInfo

//
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { // 索引从 1 开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo: // 如果是 long 或 double 则占两个位置
			i++
		}
	}
}

// 从常量池按照索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// 从常量池查找字段或方法名和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

// 从常量池查找 utf8 字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
