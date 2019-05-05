package classfile

// SourceFile 是可选属性，用于指出源文件名，其结构定义为：
// SourceFile_attribute {
// 	   u2 attribute_name_index;
// 	   u4 attribute_length;
// 	   u2 sourcefile_index;
// }
//
// 其 attribtue_length 值永为 2
// sourcefile_index 是常量池索引，指向一个 CONSTANT_Utf8_info 常量
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
