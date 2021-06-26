// Package env exports environment variables that are necessary for the bot to work
package env

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"os"
)

// The application's public key. It can be found on your application in the Developer Portal.
var PubKey ed25519.PublicKey

func init() {
	pubKeyStr, ok := os.LookupEnv("APP_PUBKEY")
	if !ok {
		fmt.Fprintln(os.Stderr, `"APP_PUBKEY" variable not found`)
		os.Exit(1)
	}

	if pubKeyStr == "" {
		fmt.Fprintln(os.Stderr, `"APP_PUBKEY" variable is empty`)
		os.Exit(1)
	}

	pubKeyBytes, err := hex.DecodeString(pubKeyStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Failed to decode "APP_PUBKEY":`, err.Error())
		os.Exit(1)
	}

	PubKey = pubKeyBytes
}
