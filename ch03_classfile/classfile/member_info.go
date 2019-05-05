package classfile

// 和类一样，字段和方法也有自己的访问标志
// 访问标志之后也是常量池索引，给出字段名和方法名，随后又是一个常量池索引，给出字段或方法的描述符
// 最后是属性表
//
// 为了避免重复性代码，这里公用一个结构体 MemberInfo 来统一标示字段和方法
type MemberInfo struct {
	cp              ConstantPool // 保存常量池指针
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// getter 方法
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

// 读取字段或方法表，返回 MemberInfo 类型数组
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取字段或方法的数据，返回一个 MemberInfo 实例
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

// 根据 nameIndex 从常量池获取字段或方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// 根据 descriptorIndex 从常量池获取字段或方法的描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
