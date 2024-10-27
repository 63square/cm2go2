package builder

import (
	"strconv"
)

func itoa99(b []byte, num byte) []byte {
	asciiLookup := "0001020304050607080910111213141516171819" +
		"2021222324252627282930313233343536373839" +
		"4041424344454647484950515253545556575859" +
		"6061626364656667686970717273747576777879" +
		"8081828384858687888990919293949596979899"

	if num < 10 {
		b = append(b, num+48)
	} else {
		start := num * 2
		b = append(b, asciiLookup[start:start+2]...)
	}

	return b
}

func MakeBlock(buf []byte, blockType byte, x float64, y float64, z float64, properties []float64) []byte {
	buf = itoa99(buf, blockType)
	buf = append(buf, ',', ',')

	buf = strconv.AppendFloat(buf, x, 'G', 2, 32)
	buf = append(buf, ',')
	buf = strconv.AppendFloat(buf, y, 'G', 2, 32)
	buf = append(buf, ',')
	buf = strconv.AppendFloat(buf, z, 'G', 2, 32)
	buf = append(buf, ',')

	for idx, val := range properties {
		buf = strconv.AppendFloat(buf, val, 'G', 2, 32)

		if idx != len(properties)-1 {
			buf = append(buf, '+')
		}
	}

	return append(buf, ';')
}

func ConnectBlock(buf []byte, a uint64, b uint64) []byte {
	buf = strconv.AppendUint(buf, a, 10)
	buf = append(buf, ',')
	buf = strconv.AppendUint(buf, b, 10)

	return append(buf, ';')
}
