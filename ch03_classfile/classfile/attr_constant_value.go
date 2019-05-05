package classfile

// ConstantValue 是定长属性，只会出现在 field_info 结构中，用于表示常量表达式值，其结构为：
// ConstantValue_attribute {
// 	   u2 attribute_name_index;
// 	   u4 attribute_length;
// 	   u2 constantvalue_index;
// }
//
// attribute_length 值永为 2，constantvalue_index 是常量池索引，但具体指向的常量因
// 字段类型而异，如 CONSTANT_Long_info，CONSTANT_String_info 等等
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
