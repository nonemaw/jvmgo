package classfile

// CONSTANT_NameAndType_info 给出了字段或方法的名称和描述符，其结构如下：
//
// CONSTANT_NameAndType_info {
// 	   u1 tag;
// 	   u2 name_index;
// 	   u2 descriptor_index;
// }
//
// CONSTANT_Class_info 和 CONSTANT_NameAndType_info 加在一起可以唯一确定一个字段或者方法：
// 1. 字段或方法名由 name_index 给出
// 2. 字段或方法的描述符由 descriptor_index 给出
// 二者都是常量池索引，指向 CONSTANT_Utf8_info 常量。字段和方法名就是代码中出现或编译器生成的字段或方法名
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}

// JVM 规范定义了一种简单的语法来描述字段或方法，并生成描述符 descriptor：
// A. 类型描述符
//   - 基本类型 byte、short、char、int、long、float 和 double 的描述符为单个字母，分别是
//             B    S      C     I    J    F        D
//   - 引用类型的描述符是 "L" + className + ";"
//   - 数组类型的描述符是 "[" + 数组元素类型描述符
// B. 字段描述符
//   - 字段描述符就是字段类型的描述符
// C. 方法描述符
//   - 方法描述符是 分号分隔的参数类型描述符 + 返回值类型描述符，如果返回值是 void 则以单个字母 "V" 表示

// 关于方法的重载，JVM 是如何根据参数的类型和数量来识别重载的方法的？
// 这是因为 CONSTANT_NameAndType_info 结构会同时包含名称和描述符的缘故。对于方法描述符而言，参数的不同也
// 就意味着方法描述符的不同，以此区分同名的重载方法
