package sarama

import "unsafe"

func MurmurHash2(key []byte) uint32 {
	var (
		seed    = uint32(0x9747b28c)
		m       = uint32(0x5bd1e995)
		r       = uint32(24)
		dataLen = uint32(len(key))

		h    = uint32(seed ^ dataLen)
		data = key
		i    int
		k    uint32
	)

	for {
		if dataLen < 4 {
			break
		}

		k = *(*uint32)(unsafe.Pointer(&data[i*4]))

		k *= m
		k ^= k >> r
		k *= m

		h *= m
		h ^= k

		i++
		dataLen -= 4
	}

	switch dataLen {
	case 3:
		h ^= uint32(data[2]) << 16
		fallthrough

	case 2:
		h ^= uint32(data[1]) << 8
		fallthrough

	case 1:
		h ^= uint32(data[0])
		h *= m

	default:
	}

	h ^= h >> 13
	h *= m
	h ^= h >> 15

	return h
}

func toPositive(number uint32) uint32 {
	return number & 0x7fffffff
}
