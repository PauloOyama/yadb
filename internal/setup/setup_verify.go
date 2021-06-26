package setup

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/agstrc/yadb/internal/setup/env"
	"github.com/gofiber/fiber/v2"
)

// verify is a middleware that checks for Discord's signature according to the API documentations:
// https://discord.com/developers/docs/interactions/slash-commands#security-and-authorization
func verify(c *fiber.Ctx) error {
	signatureStr := c.Get("X-Signature-Ed25519")
	timestamp := c.Get("X-Signature-Timestamp")

	if signatureStr == "" || timestamp == "" {
		c.Status(401)
		return nil
	}

	signatureBytes, err := hex.DecodeString(signatureStr)
	if err != nil {
		c.Status(401)
		return nil
	}

	if len(signatureBytes) != ed25519.SignatureSize || signatureBytes[63]&224 != 0 {
		c.Status(401)
		return nil
	}

	message := append([]byte(timestamp), c.Body()...)

	if !ed25519.Verify(env.PubKey, message, signatureBytes) {
		c.Status(401)
		return nil
	}

	return c.Next()
}
