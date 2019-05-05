package classfile

// CONSTANT_String_info 常量表示 java.lang.Sintrg，其结构体如下
//
// CONSTANT_String_info {
// 	   u1 tag;
// 	   u2 string_index;
// }
//
// 可以看到 CONSTANT_String_info 本身是不存放字符串数据的，它只存放了常量池的索引，而这个索引指向了
// 一个 CONSTANT_Utf8_info 常量
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// readInfo() 读取常量池索引
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// String() 方法从常量池中根据索引查找字符串
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
