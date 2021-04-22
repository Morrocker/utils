package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// Trimmer asfda
func Trimmer(hash string, headntail ...int) string {
	var head, tail string
	switch len(headntail) {
	case 0:
		if len(hash) < 13 {
			return "string too short to be trimmed"
		}
		head = hash[:6]
		tail = hash[len(hash)-6:]
	case 1:
		if len(hash) <= headntail[0] {
			return "string too short to be trimmed"
		}
		head = hash[:headntail[0]]
		tail = ""
	case 2:
		if len(hash) <= headntail[0]+headntail[1] {
			return "string too short to be trimmed"
		}
		head = hash[:headntail[0]]
		tail = ""
	default:
		return "too many arguments"
	}

	return fmt.Sprintf("%s...%s", head, tail)
}

func B2H(n int64) string {
	idx := 0
	b := float64(n)
	byteUnits := []string{"b", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	for b > 1024 {
		b /= 1024
		idx++
	}

	if idx == 0 {
		return fmt.Sprintf("%.0f b", b)
	}
	return fmt.Sprintf("%.1f %s", b, byteUnits[idx])
}

// GetFunctionName
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
