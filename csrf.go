package csrf

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Secret is the key to use when creating the sha256 hash
var Secret string

// MaxTokenAge is the time.Duration of how long the CSRF token should be valid
var MaxTokenAge time.Duration


// CreateToken generates the CSRF token based on sessionID
func CreateToken(sessionID string) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	data := sessionID + timestamp

	sha := generateHMAC(data)

	return sha + "$" + timestamp
}


// ValidToken returns true if the token makes the sessionID and is within MaxTokenAge
func ValidToken(token, sessionID string) bool {
	splitToken := strings.Split(token, "$")
	if len(splitToken) != 2 {
		fmt.Println("CSRF Token Missing Timestamp")
		return false
	}

	sha := splitToken[0]
	timestamp := splitToken[1]

	data := sessionID + timestamp

	newSha := generateHMAC(data)

	if (newSha != sha) {
		fmt.Println("CSRF Token Mismatch")
		return false
	}

	milsec, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fullTime := time.Unix(milsec, 0)

	if fullTime.Sub(time.Now()) > MaxTokenAge {
		fmt.Println("CSRF Token Expired")
		return false
	}

	return true
}


// generateHMAC returns hash
func generateHMAC(data string) string {
	h := hmac.New(sha256.New, []byte(Secret))
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}
