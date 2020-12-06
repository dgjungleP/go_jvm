package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}
type ConstantFiledrefInfo struct{ ConstantMemberrefInfo }
type ConstanMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

func (constant *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	constant.classIndex = reader.readUint16()
	constant.nameAndTypeIndex = reader.readUint16()
}
func (constant *ConstantMemberrefInfo) String() string {
	return constant.cp.getClassName(constant.classIndex)
}

func (constant *ConstantMemberrefInfo) NameANdTypeDescriptor() (string, string) {
	return constant.cp.getNameType(constant.nameAndTypeIndex)
}
