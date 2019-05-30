package main

import (
	"./types"
	sha "crypto/sha512"
	"encoding/base64"
	"fmt"
	"time"
	"strings"
)

// This type of authentication has weaknesses:
// * weak passwords can be brute forced (no salt, fast hash function)
// * attacker can replay messages for config.MAX_DELAY
//	* could be prevented by prohibitting same txs within a certain
//	  time or or by adding a unique number for each message (counter,
//	  unique random number)

func IsValid(auth types.Authentication, password string, maxDelay time.Duration) bool {
	if auth.Time.Add(maxDelay).Before(time.Now()) {
		return false
	}
	if strings.ToUpper(auth.From)[0] == auth.From[0] {
		println("someone tried to spend from a Project!")
		return false
	}
	toHash := fmt.Sprintf(
		"%s%s%v%v%s",
		auth.From,
		auth.To,
		auth.Amount,
		auth.Time.UnixNano()/1000000,
		password)
	hashed := sha.Sum384([]byte(toHash))
	digest := base64.StdEncoding.EncodeToString(hashed[:])
	println("Digest:", digest)
	return digest == auth.Digest
}
