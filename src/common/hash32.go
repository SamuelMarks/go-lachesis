package common

import "hash/fnv"

func Hash32(data []byte) int64 {
	h := fnv.New32a()

	h.Write(data)

	return int64(h.Sum32())
}
