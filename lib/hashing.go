package lib

func BinHash(str string) uint32 {
	hash := uint32(0xFFFFFFFF)

	for _, c := range str {
		hash = hash*33 + uint32(c)
	}

	return hash
}

func VltHash(str string) uint32 {
	initVal := uint32(0xABCDEF00)
	strOffset := 0
	strLength := len(str)
	firstFactor := uint32(0x9e3779b9)
	secondFactor := firstFactor
	thirdFactor := initVal

	for strLength >= 12 {
		firstFactor += uint32(str[strOffset]) + uint32(str[strOffset+1]<<8) + uint32(str[strOffset+2]<<16) + uint32(str[strOffset+3]<<24)
		secondFactor += uint32(str[strOffset+4]) + uint32(str[strOffset+5]<<8) + uint32(str[strOffset+6]<<16) + uint32(str[strOffset+7]<<24)
		thirdFactor += uint32(str[strOffset+8]) + uint32(str[strOffset+9]<<8) + uint32(str[strOffset+10]<<16) + uint32(str[strOffset+11]<<24)

		//                 a -= b; a -= c; a ^= (c >> 13);
		firstFactor -= secondFactor
		firstFactor -= thirdFactor
		firstFactor ^= thirdFactor >> 13

		//                 b -= c; b -= a; b ^= (a << 8);
		secondFactor -= thirdFactor
		secondFactor -= firstFactor
		secondFactor ^= firstFactor << 8

		//                 c -= a; c -= b; c ^= (b >> 13);
		thirdFactor -= firstFactor
		thirdFactor -= secondFactor
		thirdFactor ^= secondFactor >> 13

		//                a -= b; a -= c; a ^= (c >> 12);
		firstFactor -= secondFactor
		firstFactor -= thirdFactor
		firstFactor ^= thirdFactor >> 12
		//                b -= c; b -= a; b ^= (a << 16);
		secondFactor -= thirdFactor
		secondFactor -= firstFactor
		secondFactor ^= firstFactor << 16
		//                c -= a; c -= b; c ^= (b >> 5);
		thirdFactor -= firstFactor
		thirdFactor -= secondFactor
		thirdFactor ^= secondFactor >> 5
		//                a -= b; a -= c; a ^= (c >> 3);
		firstFactor -= secondFactor
		firstFactor -= thirdFactor
		firstFactor ^= thirdFactor >> 3
		//                b -= c; b -= a; b ^= (a << 10);
		secondFactor -= thirdFactor
		secondFactor -= firstFactor
		secondFactor ^= firstFactor << 10
		//                c -= a; c -= b; c ^= (b >> 15);
		thirdFactor -= firstFactor
		thirdFactor -= secondFactor
		thirdFactor ^= secondFactor >> 15

		strOffset += 12
		strLength -= 12
	}

	thirdFactor += uint32(len(str))

	switch strLength {
	case 11:
		thirdFactor += uint32(str[10+strOffset]) << 24
		fallthrough
	case 10:
		thirdFactor += uint32(str[9+strOffset]) << 16
		fallthrough
	case 9:
		thirdFactor += uint32(str[8+strOffset]) << 8
		fallthrough
	case 8:
		secondFactor += uint32(str[7+strOffset]) << 24
		fallthrough
	case 7:
		secondFactor += uint32(str[6+strOffset]) << 16
		fallthrough
	case 6:
		secondFactor += uint32(str[5+strOffset]) << 8
		fallthrough
	case 5:
		secondFactor += uint32(str[4+strOffset])
		fallthrough
	case 4:
		firstFactor += uint32(str[3+strOffset]) << 24
		fallthrough
	case 3:
		firstFactor += uint32(str[2+strOffset]) << 16
		fallthrough
	case 2:
		firstFactor += uint32(str[1+strOffset]) << 8
		fallthrough
	case 1:
		firstFactor += uint32(str[strOffset])
		break
	}

	//                 a -= b; a -= c; a ^= (c >> 13);
	firstFactor -= secondFactor
	firstFactor -= thirdFactor
	firstFactor ^= thirdFactor >> 13

	//                 b -= c; b -= a; b ^= (a << 8);
	secondFactor -= thirdFactor
	secondFactor -= firstFactor
	secondFactor ^= firstFactor << 8

	//                 c -= a; c -= b; c ^= (b >> 13);
	thirdFactor -= firstFactor
	thirdFactor -= secondFactor
	thirdFactor ^= secondFactor >> 13

	//                a -= b; a -= c; a ^= (c >> 12);
	firstFactor -= secondFactor
	firstFactor -= thirdFactor
	firstFactor ^= thirdFactor >> 12
	//                b -= c; b -= a; b ^= (a << 16);
	secondFactor -= thirdFactor
	secondFactor -= firstFactor
	secondFactor ^= firstFactor << 16
	//                c -= a; c -= b; c ^= (b >> 5);
	thirdFactor -= firstFactor
	thirdFactor -= secondFactor
	thirdFactor ^= secondFactor >> 5
	//                a -= b; a -= c; a ^= (c >> 3);
	firstFactor -= secondFactor
	firstFactor -= thirdFactor
	firstFactor ^= thirdFactor >> 3
	//                b -= c; b -= a; b ^= (a << 10);
	secondFactor -= thirdFactor
	secondFactor -= firstFactor
	secondFactor ^= firstFactor << 10
	//                c -= a; c -= b; c ^= (b >> 15);
	thirdFactor -= firstFactor
	thirdFactor -= secondFactor
	thirdFactor ^= secondFactor >> 15

	return thirdFactor
}
