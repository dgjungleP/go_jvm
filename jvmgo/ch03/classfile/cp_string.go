package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (constant *ConstantStringInfo) readInfo(reader *ClassReader) {
	constant.stringIndex = reader.readUint16()
}
func (constant *ConstantStringInfo) String() string {
	return constant.cp.getUtf8(constant.stringIndex)
}
