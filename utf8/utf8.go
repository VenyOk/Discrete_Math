package main

func encode(utf32 []rune) []byte {
	utf8 := make([]byte, 0)

	for _, codePoint := range utf32 {
		if codePoint <= 0x7F {

			utf8 = append(utf8, byte(codePoint))
		} else if codePoint <= 0x7FF {

			utf8 = append(utf8, byte(0xC0|(codePoint>>6)))
			utf8 = append(utf8, byte(0x80|(codePoint&0x3F)))
		} else if codePoint <= 0xFFFF {

			utf8 = append(utf8, byte(0xE0|(codePoint>>12)))
			utf8 = append(utf8, byte(0x80|((codePoint>>6)&0x3F)))
			utf8 = append(utf8, byte(0x80|(codePoint&0x3F)))
		} else {

			utf8 = append(utf8, byte(0xF0|(codePoint>>18)))
			utf8 = append(utf8, byte(0x80|((codePoint>>12)&0x3F)))
			utf8 = append(utf8, byte(0x80|((codePoint>>6)&0x3F)))
			utf8 = append(utf8, byte(0x80|(codePoint&0x3F)))
		}
	}

	return utf8
}

func decode(utf8 []byte) []rune {
	utf32 := make([]rune, 0)

	for i := 0; i < len(utf8); {
		b1 := utf8[i]
		var codePoint rune

		if (b1 & 0x80) == 0 {

			codePoint = rune(b1)
			i++
		} else if (b1 & 0xE0) == 0xC0 {

			codePoint = rune(b1&0x1F) << 6
			codePoint |= rune(utf8[i+1] & 0x3F)
			i += 2
		} else if (b1 & 0xF0) == 0xE0 {

			codePoint = rune(b1&0x0F) << 12
			codePoint |= rune(utf8[i+1]&0x3F) << 6
			codePoint |= rune(utf8[i+2] & 0x3F)
			i += 3
		} else if (b1 & 0xF8) == 0xF0 {

			codePoint = rune(b1&0x07) << 18
			codePoint |= rune(utf8[i+1]&0x3F) << 12
			codePoint |= rune(utf8[i+2]&0x3F) << 6
			codePoint |= rune(utf8[i+3] & 0x3F)
            i += 4
		}
		utf32 = append(utf32, codePoint)
	}

	return utf32
}

func main(){
}
