/*
Go has built-in support for base64 encoding
*/ 

package main	

// This syntaz imports the encoding/base64 package with
// b64 name instead of default base64 ie. "import base64 as b64"
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// encode/decode the following string
	data := "abc123!?$*&()'-=@~"

	// We can use both standard and URL-compatible base64
	// The encoder requires a []byte
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Check if Decoding returns an error, as it may do
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// To encode/decode useing a URL-compatible base64 format
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}