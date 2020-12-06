package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (constant *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	constant.nameIndex = reader.readUint16()
	constant.descriptorIndex = reader.readUint16()
}
