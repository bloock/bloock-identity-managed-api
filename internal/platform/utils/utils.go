package utils

import (
	"crypto/rand"
	"encoding/binary"
	mathRand "math/rand"
	"time"
)

func RandInt64() (uint64, error) {
	var buf [8]byte
	_, err := rand.Read(buf[:4])

	return binary.LittleEndian.Uint64(buf[:]), err
}

func RandInt32() uint32 {
	mathRand.Seed(time.Now().UnixNano())

	return mathRand.Uint32()
}
