package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (constantPool ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := constantPool[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}
func (constantPool ConstantPool) getNameType(index uint16) (string, string) {
	ntIfdo := constantPool.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := constantPool.getUtf8(ntInfo.nameIndex)
	_type := constantPool.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}
func (constantPool ConstantPool) getClassName(index uint16) string {
	classInfo := constantPool.getConstantInfo(index).(*ConstantClassInfo)
	return constantPool.getUtf8(classInfo.nameIndex)
}
func (constantPool ConstantPool) getUtf8(index uint16) string {
	utf8Info := constantPool.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
