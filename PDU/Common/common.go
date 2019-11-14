package Common

func DecodeUnsigned(signed byte) uint16 {
	return uint16(signed)
}

func DecodeUnsignedFromInt16(signed uint16) int {
	return int(signed)
}

func EncodeUnsigned(positive uint16) byte {
	return byte(positive)
}

func EncodeUnsignedFromInt(positive int) uint16 {
	return uint16(positive)
}
