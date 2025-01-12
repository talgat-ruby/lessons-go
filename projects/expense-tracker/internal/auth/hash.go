package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
)

const (
	// Recommended minimum lengths for security
	SaltLength   = 16     // 16 bytes = 128 bits
	PepperLength = 32     // 32 bytes = 256 bits
	HashLength   = 64     // SHA-512 outputs 64 bytes
	Iterations   = 210000 // Recommended minimum iterations for PBKDF2
)

func HashPassword(password string, pepper string) (string, string, error) {
	// Generate random salt
	salt := make([]byte, SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", "", fmt.Errorf("error generating salt: %w", err)
	}

	// Hash the password
	hash := hashWithSaltAndPepper([]byte(password), salt, []byte(pepper))

	return base64.StdEncoding.EncodeToString(hash), base64.StdEncoding.EncodeToString(salt), nil
}

func hashWithSaltAndPepper(password, salt, pepper []byte) []byte {
	// Combine password and pepper
	pepperedPassword := make([]byte, len(password)+len(pepper))
	copy(pepperedPassword, password)
	copy(pepperedPassword[len(password):], pepper)

	// Initialize HMAC-SHA512
	hash := hmac.New(sha512.New, salt)

	// Perform multiple iterations of hashing
	result := pepperedPassword
	for i := 0; i < Iterations; i++ {
		hash.Reset()
		hash.Write(result)
		result = hash.Sum(nil)
	}

	return result
}

// VerifyPassword checks if a password matches its hash
func VerifyPassword(password, pepper, hash, salt string) (bool, error) {
	// Decode salt and hash from base64
	decodedSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return false, fmt.Errorf("error decoding salt: %w", err)
	}

	decodedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false, fmt.Errorf("error decoding hash: %w", err)
	}

	// Hash the provided password with the same salt
	newHash := hashWithSaltAndPepper([]byte(password), decodedSalt, []byte(pepper))

	// Compare hashes in constant time to prevent timing attacks
	return subtle.ConstantTimeCompare(newHash, decodedHash) == 1, nil
}
