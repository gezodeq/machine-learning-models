package models

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GenerateUUID() string {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	max := len(alphabet) - 1
	buf := make([]byte, 32)
	for i := range buf {
		buf[i] = alphabet[rand.Intn(max)]
	}
	return strings.ToUpper(string(buf))
}

func GenerateSHA256(input string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func GetElapsedTime(start time.Time) string {
	return fmt.Sprintf("%fs", time.Since(start).Seconds())
}

func TimeToInt(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func IntToTime(n int64) time.Time {
	return time.Unix(n/time.Second, 0).UTC()
}

func Int64ToTime(n int64) time.Time {
	return time.Unix(n, 0).UTC()
}

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}