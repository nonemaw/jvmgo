package classfile

// CONSTANT_Fieldref_info 表示字段符号引用
// CONSTANT_Methodref_info 表示非接口方法引用
// CONSTANT_InterfaceMethodref_info 表示接口方法引用
// 这三种常量结构一模一样，以 CONSTANT_Fieldref_info 为例：
//
// CONSTANT_Fieldref_info {
// 	   u1 tag;
// 	   u2 class_index;         指向 CONSTANT_Class_info
// 	   u2 name_and_type_index; 指向 CONSTANT_NameAndType_info
// }

// 这里先规范一个『基类』结构体
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// golang 没有继承的概念，但是可以通过架构提嵌套来模拟
// 这里我们根据之前创建的『基类』结构体来衍生出字段符号、非接口方法、接口方法引用
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

// 还有一些其它的常量没有包括在里面：
// CONSTANT_MethodType_info
// CONSTANT_MethodHandle_info
// CONSTANT_InvokeDynamic_info
// 他们是 Java7 之后才支持的常量类型，用于支持 invokeDynamic 指令，这里不讨论这一块
