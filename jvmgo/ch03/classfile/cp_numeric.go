package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (constant *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	constant.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (constant *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	constant.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (constant *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	constant.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (constant *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	constant.val = math.Float64frombits(bytes)
}
