package utility

import (
	"crypto/sha1"
	"fmt"
)

func calculateSHA1(input []byte) string {
	sha1Hash := sha1.New()
	sha1Hash.Write(input)
	hashBytes := sha1Hash.Sum(nil)
	sha1String := fmt.Sprintf("%x", hashBytes)
	return sha1String
}
func getHexValue(input []byte) string {
	return fmt.Sprintf("%x", input)
}
