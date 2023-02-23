package main

import "fmt"

func encode(utf32 []rune) (b []byte) {
	b = make([]byte, 0, len(utf32)*4)
	for _, r := range utf32 {
		if r < 0x80 {
			b = append(b, byte(r))
		} else if r < 0x800 {
			b = append(b, 0xc0|(byte(r)>>6))
			b = append(b, 0x80|(byte(r)&0x3f))
		} else if r < 0x10000 {
			b = append(b, 0xe0|(byte(r)>>12))
			b = append(b, 0x80|(byte(r)>>6)&0x3f)
			b = append(b, 0x80|(byte(r)&0x3f))
		} else {
			b = append(b, 0xf0|(byte(r)>>18))
			b = append(b, 0x80|(byte(r)>>12)&0x3f)
			b = append(b, 0x80|(byte(r)>>6)&0x3f)
			b = append(b, 0x80|(byte(r)&0x3f))
		}
	}
	return
}

func decode(utf8 []byte) (runes []rune) {
	runes = make([]rune, 0, len(utf8))
	for i := 0; i < len(utf8); {
		r, size := utf8DecodeRune(utf8[i:])
		runes = append(runes, r)
		i += size
	}
	return
}
func utf8DecodeRune(b []byte) (rune, int) {
	if len(b) < 1 {
		return 0, 0
	}
	if b[0] < 0x80 {
		return rune(b[0]), 1
	}
	if b[0] < 0xc2 {
		return 0, 1
	}
	if b[0] < 0xe0 {
		if len(b) < 2 || b[1]&0xc0 != 0x80 {
			return 0, 1
		}
		r := rune(b[0]&0x1f)<<6 | rune(b[1]&0x3f)
		if r < 0x80 {
			return 0, 1
		}
		return r, 2
	}
	if b[0] < 0xf0 {
		if len(b) < 3 || b[1]&0xc0 != 0x80 || b[2]&0xc0 != 0x80 {
			return 0, 1
		}
		r := rune(b[0]&0xf)<<12 | rune(b[1]&0x3f)<<6 | rune(b[2]&0x3f)
		if r < 0x800 {
			return 0, 1
		}
		return r, 3
	}
	if b[0] < 0xf8 {
		if len(b) < 4 || b[1]&0xc0 != 0x80 || b[2]&0xc0 != 0x80 || b[3]&0xc0 != 0x80 {
			return 0, 1
		}
		r := rune(b[0]&0x7)<<18 | rune(b[1]&0x3f)<<12 | rune(b[2]&0x3f)<<6 | rune(b[3]&0x3f)
		if r < 0x10000 {
			return 0, 1
		}
		return r, 4
	}
	return 0, 1
}
func main() {
	utf32 := []rune{0x1005F, 0x10010, 0x10020}
	fmt.Printf("Исходная строка в UTF-32: %x\n", utf32)
	utf8 := encode(utf32)
	fmt.Printf("Кодированная строка в UTF-8: %x\n", utf8)
	decoded := decode(utf8)
	fmt.Printf("Декодированная строка из UTF-8: %x\n", decoded)
}
