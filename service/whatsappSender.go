package service

import (
	"fmt"
	"math/rand"
	"time"
)

func GetOtp(phoneNumber string) (interface{}, error) {

	rand.Seed(time.Now().Unix())
	number := rand.Int31n(99999)
	fmt.Println(number)

	return "adsad", nil
}

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "1234567890ABCDEFGHIJabcdefg"
	fmt.Println("Len Letter => ", len(letterBytes))
	b := make([]byte, n)
	fmt.Println(b)
	for i := range b {
		fmt.Println(rand.Intn(len(letterBytes)))
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	fmt.Println(b)
	return string(b)
}
