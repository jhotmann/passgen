package main

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"unicode/utf16"

	"ekyu.moe/base91"
	"github.com/aead/skein"
)

func passgenV1(opts opts) string {
	encoded := utf16.Encode([]rune(opts.passphrase + opts.salt))
	var hashOut [64]byte
	skein.Sum512(&hashOut, convertUTF16ToLittleEndianBytes(encoded), nil)
	hashHex := hex.EncodeToString(hashOut[:])

	start := hexToInt(string(hashHex[0])) + hexToInt(string(hashHex[1])) + hexToInt(string(hashHex[2]))
	end := start + opts.length

	encoded91 := base91.EncodeToString([]byte(hashHex))

	for {
		if len(encoded91) > end {
			break
		}
		encoded91 = encoded91 + encoded91
	}

	return encoded91[start:end]
}

func convertUTF16ToLittleEndianBytes(u []uint16) []byte {
	b := make([]byte, 2*len(u))
	for index, value := range u {
		binary.LittleEndian.PutUint16(b[index*2:], value)
	}
	return b
}

func hexToInt(hexString string) int {
	i, _ := strconv.ParseInt(hexString, 16, 0)
	return int(i)
}
