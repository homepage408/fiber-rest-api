package service

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func HashPassword(password string, defaultSalt, saltText string) (string, error) {
// 	salt := sha512hash(strings.ToLower(saltText), defaultSalt)
// 	return sha512hash(password, salt), nil
// }

// func sha512hash(text, salt string) string {
// 	if text == "" {
// 		return ""
// 	}

// 	clearTextBytes := []byte(text)
// 	saltBytes := []byte(salt)
// 	sha512Hasher := sha512.New()

// 	clearTextBytes = append(clearTextBytes, saltBytes...)
// 	sha512Hasher.Write(clearTextBytes)

// 	hashedPasswordBytes := sha512Hasher.Sum(nil)

// 	return base64.URLEncoding.EncodeToString(hashedPasswordBytes)
// }

// func CompareHash(hashClient, expectedHash string) bool {
// 	return hashClient == expectedHash
// }
