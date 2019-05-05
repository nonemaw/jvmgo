package classfile

import (
	"math"
)

// CONSTANT_Integer_info 使用 1 个字节存储 tag，4 个字节存储整数常量，其结构定义为
//
// CONSTANT_Integer_info {
// 	   u1 tag;
// 	   u4 bytes;
// }
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

// CONSTANT_Float_info 使用 1 个字节存储 tag，4 个字节存储浮点常量，其结构定义为
//
// CONSTANT_Float_info {
// 	   u1 tag;
// 	   u4 bytes;
// }
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes) // 将 4 个字节转换为浮点
}

// CONSTANT_Double_info 使用 1 个字节存储 tag，8 个字节存储双精度浮点常量，其结构定义为
//
// CONSTANT_Double_info {
// 	   u1 tag;
// 	   u4 high_bytes;
//     u4 low_bytes;
// }
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

// CONSTANT_Long_info 使用 1 个字节存储 tag，8 个字节存储整数常量，其结构定义为
//
// CONSTANT_Long_info {
// 	   u1 tag;
// 	   u4 high_bytes;
//     u4 low_bytes;
// }
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}
