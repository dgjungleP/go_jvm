package classfile

type ConstantClassInfo struct {
	cp        ConstantInfo
	nameIndex uint16
}

func (constant *ConstantClassInfo) readInfo(reader *ClassReader) {
	constant.nameIndex = reader.readUint16()
}

func (constant *ConstantClassInfo) String() string {
	return constant.cp.getUtf8(constant.stringIndex)
}
