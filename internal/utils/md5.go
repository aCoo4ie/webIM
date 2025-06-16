package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// HashPassword hashes a plaintext password with a salt using MD5.
// It returns the hashed password as a hexadecimal string.
func HashPassword(password, salt string) string {
	hasher := md5.New()
	// Combine the password and salt before hashing
	hasher.Write([]byte(password + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

// ValidatePassword checks if the provided plaintext password and salt,
// when hashed, match the stored hashed password.
// Returns true if they match, false otherwise.
func ValidatePassword(plainPassword, salt, hashedPassword string) bool {
	return hashedPassword == HashPassword(plainPassword, salt)
}

func HashToken(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
