package util

import (
	"bytes"
	"encoding/binary"
)

func InterfaceToBinaryBytes(vs ...interface{}) []byte {
	var b bytes.Buffer
	for _, v := range vs {
		binary.Write(&b, binary.LittleEndian, v)
	}
	return b.Bytes()
}

// TODO from?
