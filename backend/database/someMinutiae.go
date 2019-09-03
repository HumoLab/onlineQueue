package database

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"time"
)

const (
	alphabet  = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm0123456789"
	soulLenth = 12 // 10^10 years!!! If YOU try Brute-Force attack:)
)

// GenerateSOUL - SOUL for passwords
func generateSOUL() string {
	rand.Seed(time.Now().UnixNano())
	soul := []byte{}
	for i := 0; i < soulLenth; i++ {
		soul = append(soul, alphabet[rand.Intn(len(alphabet))])
	}
	return string(soul)
}

// hashing - password and soul hashing
func hashing(pass string, soul string) string {
	var hash = sha256.New()
	hash.Write([]byte(pass + soul))
	return hex.EncodeToString(hash.Sum(nil))
}

// readFromFile
func readFromFile(path string) []byte {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte("error")
	}
	return fileContent
}
