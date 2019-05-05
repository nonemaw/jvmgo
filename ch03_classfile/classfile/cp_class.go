package classfile

// CONSTANT_Class_info 常量表示类或者接口符号的引用，其结构如下
// CONSTANT_Class_info {
// 	   u1 tag;
// 	   u2 name_index;
// }
//
// 它也通过常量池索引来保存信息，故代码和 CONSTANT_Stirng_info 类似
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
