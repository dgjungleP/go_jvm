package classfile

type ConstantUtf8Info struct {
	str string
}

func (constat *ConstantUtf8Info) readInfo(reader *ClassReader) {
	lenght := uint32(reader.readUint16())
	bytes := reader.readBytes(lenght)
	constant.str = decodeMUTF8(bytes)
}
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
