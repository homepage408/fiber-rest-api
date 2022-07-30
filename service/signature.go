package service

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	// "xti-gateway-go/pkg/model"

	"github.com/joho/godotenv"
)

func GetPublicKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	// publicKey := private.PublicKey

	fmt.Println(privateKey.PublicKey)

	return privateKey, err
}

func Signature() (string, []byte) {

	privateKey, _ := GetPublicKey()

	excryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&privateKey.PublicKey,
		[]byte("HAHAH HIHI"),
		nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted => ", excryptedBytes)
	return string(excryptedBytes), excryptedBytes
}

func SignatureDecrypt() string {
	privateKey, _ := GetPublicKey()
	_, encryptedBytes := Signature()

	fmt.Println("Dec => ", encryptedBytes)

	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	fmt.Println("err => ", err)
	fmt.Println("Decrypted => ", decryptedBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decryptedBytes))
	return string(decryptedBytes)
}

func GenerateSiganturePayment(requestSignature string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// buf := new(bytes.Buffer)

	secret := os.Getenv("SECRET_KEY_SIGNATURE")
	paymentData := requestSignature
	fmt.Println("Secret key: ", secret)
	fmt.Println("Payment Data : ", paymentData)

	paymentDataJson, err := json.Marshal(paymentData)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Println(string(paymentDataJson))

	// error := binary.Write(buf, binary.BigEndian, &paymentDataJson)
	// if error != nil {
	// 	fmt.Println("binary.Write failed:", error)
	// }

	// fmt.Println(result)
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(paymentDataJson))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)
}

// func ValidMAC(message model.DataTestPayment, messageMAC string) bool {

// 	secret := os.Getenv("SECRET_KEY_SIGNATURE")
// 	fmt.Println("Secret key: ", secret)

// 	mac := hmac.New(sha256.New, []byte(secret))

// 	paymentData := message
// 	fmt.Println("Secret key: ", secret)
// 	fmt.Println("Payment Data : ", paymentData)

// 	paymentDataJson, err := json.Marshal(paymentData)
// 	if err != nil {
// 		fmt.Printf("Error: %s", err)
// 	}

// 	fmt.Println("Payment Json : ", string(paymentDataJson))
// 	mac.Write(paymentDataJson)
// 	expectedMAC := mac.Sum(nil)

// 	data := hmac.Equal([]byte(messageMAC), []byte(expectedMAC))
// 	fmt.Println("False / True => ", data)
// 	return data
// }
