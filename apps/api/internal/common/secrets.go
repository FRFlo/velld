package common

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Secrets struct {
	JWTSecret     string
	EncryptionKey string
}

var secretsFilePath = filepath.Join(".", ".env")

var once sync.Once
var instance *Secrets

func GetSecrets() *Secrets {
	once.Do(func() {
		instance = loadOrGenerateSecrets()
	})
	return instance
}

func loadOrGenerateSecrets() *Secrets {
	jwtSecret := getOrGenerateSecret("JWT_SECRET")
	encryptionKey := getOrGenerateSecret("ENCRYPTION_KEY")

	saveSecretsToFile(jwtSecret, encryptionKey)

	return &Secrets{
		JWTSecret:     jwtSecret,
		EncryptionKey: encryptionKey,
	}
}

func getOrGenerateSecret(envVar string) string {
	if secret := os.Getenv(envVar); secret != "" {
		return secret
	}

	var newSecret string
	if envVar == "ENCRYPTION_KEY" {
		newSecret = generateSecureHexKey()
	} else {
		newSecret = generateSecureBase64Key()
	}

	fmt.Printf("[INFO] %s not found in env. Generated new key.\n", envVar)
	return newSecret
}

// Generate a secure random hex key (for AES encryption)
func generateSecureHexKey() string {
	key := make([]byte, 32) // 32 bytes = 64 hex characters
	if _, err := rand.Read(key); err != nil {
		log.Fatal("Failed to generate encryption key:", err)
	}
	return hex.EncodeToString(key) // Convert to hex string
}


// Generate a secure random base64 key (for JWT)
func generateSecureBase64Key() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("Failed to generate secure key")
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func saveSecretsToFile(jwtSecret, encryptionKey string) {
	content := fmt.Sprintf("JWT_SECRET=%s\nENCRYPTION_KEY=%s\n", jwtSecret, encryptionKey)
	os.WriteFile(secretsFilePath, []byte(content), 0600) // Secure permissions
}
