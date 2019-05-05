package classfile

// CONSTANT_Utf8_info 用于存放 MUTF-8 编码的字符串，其结构定义如下
//
// CONSTANT_Utf8_info {
//     u1 tag;
// 	   u2 length;
// 	   u1 bytes[length];
// }
//
// Java 默认使用 MUTF-8（非标准 UTF-8）存储，原因暂时未知，二者编码非常类似但互不兼容
// http://stackoverflow.com/questions/15440584/why-does-java-usemodified-utf-8-instead-of-utf-8
// http://www.oracle.com/technetwork/articles/javase/supplementary-142654.html
type ConstantUtf8Info struct {
	str string
}

// 先读取出 class 文件常量字节 []byte，然后将其解码为 golang 的标准 UTF-8 字符串
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// TODO: 简化版，完成版查看项目源码
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
