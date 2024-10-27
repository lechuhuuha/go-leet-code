package sequences

import "bytes"

func ThueMorseSequence(buffer *bytes.Buffer) {
	for b, currLength, currBytes := 0, buffer.Len(), buffer.Bytes(); b <
		currLength; b++ {
		if currBytes[b] == '1' {
			buffer.WriteByte('0')
		} else {
			buffer.WriteByte('1')
		}
	}
}
