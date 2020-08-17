package rand

import (
	"errors"
	"math"
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func StringByLetters(n int) (string, error) {
	return StringByBytes([]byte(letterBytes), n)
}

func StringByBytes(bytes []byte, n int) (string, error) {
	indexBitNum, indexMask, indexNum := getIndexData(len(bytes))
	if indexNum <= 0 {
		return "", errors.New("invalid bytes")
	}

	builder := strings.Builder{}
	builder.Grow(n)
	for i, cache, remain := n-1, rand.Int63(), indexNum; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), indexNum
		}
		if offset := int(cache & indexMask); offset < len(bytes) {
			builder.WriteByte(bytes[offset])
			i--
		}

		cache >>= indexBitNum
		remain--
	}

	return builder.String(), nil
}

func StringByRunes(characters []rune, n int) (string, error) {
	indexBitNum, indexMask, indexNum := getIndexData(len(characters))
	if indexNum <= 0 {
		return "", errors.New("invalid characters")
	}

	builder := strings.Builder{}
	builder.Grow(n)
	for i, cache, remain := n-1, rand.Int63(), indexNum; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), indexNum
		}
		if offset := int(cache & indexMask); offset < len(characters) {
			builder.WriteRune(characters[offset])
			i--
		}

		cache >>= indexBitNum
		remain--
	}

	return builder.String(), nil
}

func getIndexData(num int) (indexBitNum uint, indexMask int64, indexNum int) {
	indexBitNum = uint(math.Ceil(math.Sqrt(float64(num))))
	indexMask = 1<<uint(indexBitNum) - 1
	indexNum = 63 / int(indexBitNum)
	return
}
