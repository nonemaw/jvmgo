package classfile

import (
	"encoding/binary"
)

// golang        <->  Java 的基本类型对照表
// -----------------
// int8          <->  byte
// uint8 (byte)  <->  N/A
// int16         <->  short
// uint16        <->  char
// int32         <->  int
// uint32 (rune) <->  N/A
// int64         <->  long
// uint64        <->  N/A
// float32       <->  float
// float64       <->  double
//
// 解析 class 文件的第一步是读取数据，虽说我们可以把 class 文件当作字节流来处理，
// 但直接操作字节不切实际，我们先创建一个结构体用于协助读取数据
// 这里的 ClassReader 结构体只是一个 byte 数组的封装
type ClassReader struct {
	data []byte
}

// 从 data 中读取一个字节 u1，注意这里没有使用索引用于记录数据未知，只是使用了 golang 自带的分片语法
// TODO: 可以优化
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// 读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

// binary.BigEndian 用于读取多字节 u2
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// 读取 uint16 数组，数组大小由开头的 uint16 数据指出
func (self *ClassReader) readUint16s() []uint16 {
	size := self.readUint16()
	res := make([]uint16, size)
	for i := range res {
		res[i] = self.readUint16()
	}
	return res
}

// 读取 uint32 4 个字节
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// 每次读取 u8 8 个字节
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}
