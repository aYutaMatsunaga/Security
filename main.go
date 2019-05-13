package main

import (
	"fmt"
	openssl "gopkg.in/Luzifer/go-openssl.v3"
  )
  
func main() {
	plaintext := "Hello World!"
	passphrase := "z4yH36a6zerhfE5427ZV"

	o := openssl.New()

	enc, err := o.EncryptBytes(passphrase, []byte(plaintext), openssl.DigestSHA256Sum)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}
	opensslEncrypted := string(enc)
	fmt.Printf("Encrypted text: %s\n", opensslEncrypted)

	dec, err := o.DecryptBytes(passphrase, []byte(opensslEncrypted), openssl.DigestSHA256Sum)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}

  	fmt.Printf("Decrypted text: %s\n", string(dec))
}