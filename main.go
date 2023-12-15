package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
)

func main() {
	var secretkey string
	secretkey = "VIG4GSHBRAS526HM777COAXVUF3J7QH5" // YOUR SECRET KEY

	secret, _ := base32.StdEncoding.DecodeString(secretkey) // DECODE
	epochSeconds := time.Now().Unix()

	otp := oneTimePassword(secret, epochSeconds/30) // CRAFT OTP
	fmt.Println(otp)                                // SEND OTP
}

func oneTimePassword(key []byte, value int64) uint32 {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(value))

	// HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(data)
	hash := hmacSha1.Sum(nil)

	// INDEX
	offset := hash[len(hash)-1] & 0x0F
	truncatedHash := hash[offset : offset+4]

	truncatedHash[0] = truncatedHash[0] & 0x7F
	number := binary.BigEndian.Uint32(truncatedHash)

	// CRAFT OTP CODE
	return number % 1000000
}
