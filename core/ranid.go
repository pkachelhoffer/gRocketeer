package core

import (
	"crypto/rand"
	"encoding/binary"
	"strconv"
)

func RandomID() string {
	// Generate random bytes
	var buf [8]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		panic(err)
	}

	// Convert bytes to int64
	id := int64(binary.BigEndian.Uint64(buf[:]))

	// Ensure it's positive (since we're dealing with signed integers)
	id = id & (1<<63 - 1)

	return strconv.FormatInt(id, 10)
}
