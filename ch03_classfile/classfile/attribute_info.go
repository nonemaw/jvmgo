package classfile

// 属性表能够存储各种信息
// 和常量池类似，各种属性的表达信息也各不相同，因此无法使用统一的结构来定义。不同之处在于，
// JVM 规范严格定义了 14 种属性，且它们可以进行扩展，使得不同的 JVM 可以实现自定义的属性类型
//
// 也因为自定义属性的允许，使得 JVM 规范中对对属性的定义中不包含 tag 信息，而是通过属性名来区分属性
// 且属性数据存放在属性名之后，这样允许 JVM 跳过无法处理的属性。一个典型的属性结构定义如下：
// attribute_info {
// 	   u2 attribute_name_index;
// 	   u2 attirbute_length;
// 	   u1 info[attribute_length]
// }
//
// 注意，属性名 attribute_name_index 并不是编码后的字符串，而是常量池的索引，指向一个存放属性名的
// CONSTANT_Utf8_info 常量

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// readAttributes() 挨个读取属性信息，并返回一个 AttributeInfo 接口实例组成的数组
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// readAttribute() 读取单个属性信息，并返回一个 AttributeInfo 接口实例
// 先读取属性名索引，然后从常量池根据索引获取属性名，然后传递给 newAttributeInfo() 创建具体实例
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	return attrInfo
}

// newAttributeInfo() 根据属性名创建 AttributeInfo 接口实例
// JVM 规范制定了 23 种属性，这里先解析其中的 8 种
//
// 按照 23 预定义属性，其可以分成三组：
// - （必选）第一组是实现 JVM 的必须属性，共有 5 种
// - （必选）第二组是 Java 类库所必须的属性，共有 12 种
// - （可选）第三组是主要提供给工具使用的属性，共有 6 种，可选意味着其不必出现在 class 文件中，JVM 本身或类库
// 中也能够实现它们
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		// Code 是变长属性，只存在 method_info 结构中，用于存放字节码等相关信息
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		// ConstantValue 是定长属性，只会出现在 field_info 结构中，用于表示常量表达式值
		return &ConstantValueAttribute{}
	case "Deprecated":
		// Deprecated 是最简单的属性，仅起到标志作用，不包含任何数据
		return &DeprecatedAttribute{}
	case "Exceptions":
		// Exception 是变长属性，记录方法抛出的异常表
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		// LineNumberTable 存放方法的行号信息，它属于可选的调试信息，不是运行时的必要信息
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		// LocalVariableTable 存放方法的局部变量信息，它属于可选的调试信息，不是运行时的必要信息
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		// SourceFile 属性是可选长属性，只会出现在 ClassFile 结构中，用于指出源文件名，它属于可选的调试信息，不是运行时的必要信息
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		// Synthetic 是最贱的属性，仅乞讨标志作用，不包含任何数据
		return &SyntheticAttribute{}
	default:
		// 未能处理的属性类型
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
