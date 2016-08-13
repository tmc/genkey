// Program genkey generates RSA private keys
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
)

var flagBits = flag.Int("bits", 2048, "bit size")

func main() {
	flag.Parse()

	key, err := rsa.GenerateKey(rand.Reader, *flagBits)
	if err != nil {
		log.Fatalln(err)
	}
	pem := string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}))
	fmt.Println(pem)
}
