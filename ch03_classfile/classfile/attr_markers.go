package classfile

// Deprecated 和 Synthetic 是最简单的两种属性，不包含任何数据，仅起到标志作用
// Deprecated 不做赘述，用于指出类、方法、字段、接口等不建议使用
// Synthetic 用于标记源文件中不存在，但由编译器自动生成的类成员，用于支持嵌套类和嵌套接口
//
// 它们可以出现在 ClassFile、filed_info 和 method_info 结构中，其结构定义为：
// Deprecated_attribute {
// 	   u2 attribute_name_index;
// 	   u4 attribute_length;
// }
// Synthetic_attribute {
// 	   u2 attribute_name_index;
// 	   u4 attribute_length;
// }
//
// 由于不包含数据，所以它们的 attribute_length 值永为 0
type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {}
