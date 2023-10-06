package hash

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func MD5SumFromFile(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Open %s err: %s", filepath, err.Error()))
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err = io.Copy(md5hash, f); err != nil {
		return "", errors.New(fmt.Sprintf("IO Copy err, err: %s", err.Error()))
	}

	return hex.EncodeToString(md5hash.Sum(nil)), nil
}

func GenerateRandomString(length int) (string, error) {
	randomString := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}

		randomString[i] = charset[randomIndex.Int64()]
	}

	return string(randomString), nil
}
