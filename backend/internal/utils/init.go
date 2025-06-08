package utils

import (
    "crypto/rand"
    "encoding/hex"
    "log"
    "os"
    "time"
)

func StartAdminKeyRotator() {
    currentKey := generateRandomKey()
    os.Setenv("ADMIN_KEY", currentKey)
    log.Println("Initial ADMIN_KEY set:", currentKey)

    go func() {
        for {
            time.Sleep(1 * time.Hour)
            newKey := generateRandomKey()
            os.Setenv("ADMIN_KEY", newKey)
            log.Println("ADMIN_KEY rotated:", newKey)
        }
    }()
}

func generateRandomKey() string {
    key := make([]byte, 32)
    if _, err := rand.Read(key); err != nil {
        panic("Failed to generate random key: " + err.Error())
    }
    return hex.EncodeToString(key)
}