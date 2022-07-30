package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

func SignatureWithRsa() {
	alicePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		os.Exit(1)
	}

	alicePublicKey := alicePrivateKey.PublicKey

	fmt.Println(alicePublicKey)

	secretMessage := "asdasdasd1238971237h1i23h1k2jh3k123h123l123k HALLOO WORLD"
	fmt.Println("Plain text => ", secretMessage)

	chiperText := EncryptOAEP(secretMessage, alicePublicKey)

	fmt.Println("ChiperText =>", chiperText)

	decryptText := DecryptOAEP(chiperText, *alicePrivateKey)
	fmt.Println("DecryptText => ", decryptText)
}

func EncryptOAEP(secretMessage string, publicKey rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")

	fmt.Println(label)
	rng := rand.Reader
	fmt.Println(rng)

	chiperText, err := rsa.EncryptOAEP(sha256.New(), rng, &publicKey, []byte(secretMessage), label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return "Error from encryption"
	}

	fmt.Println("Chiper text ASLI => ", chiperText)
	return base64.StdEncoding.EncodeToString(chiperText)
}

func DecryptOAEP(cipherText string, privKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")

	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return "Error from Decryption"
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))

	return string(plaintext)
}
