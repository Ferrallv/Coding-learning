/*
Sha1 hashes are used to compute short identities for binary
text or text blobs. git revision control uses SHA1 a lot.
*/ 

package main

// Go has several hash functions in crypto/*
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// Pattern for generating a hash:
	// sha1.New(), sha1.Write(bytes), sha1.Sum([]byte{})
	h := sha1.New()

	// Write expects bytes. You can use []byte{s}
	// to coerce a string to bytes
	h.Write([]byte(s))

	// Get a finalized has result as a byte slice
	bs := h.Sum(nil)

	// SHA1 vals are often printed in hex. Use the %x format
	// verb to convert a hash to a hex string
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}