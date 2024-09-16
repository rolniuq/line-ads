package sample

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// Create HMAC-SHA-256 signature
func createHMACSignature(secretKey, message string) string {
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64Encode(h.Sum(nil))
}

func GetJWS() string {
	header := map[string]string{
		"alg": "HS256",
		"kid": "LINEADSAMPLE",
		"typ": "text/plain",
	}

	payload := map[string]string{
		"Digest-SHA-256": secretKey,
		"Content-Type":   "application/json",
		"Date":           "20240905",
		"CanonicalURI":   "/v3/groups/G08916310298/link-request",
	}

	headerJSON, _ := json.Marshal(header)
	payloadJSON, _ := json.Marshal(payload)

	// Step 3: Base64 encode the header and payload
	encodedHeader := base64Encode(headerJSON)
	encodedPayload := base64Encode(payloadJSON)

	// Step 4: Concatenate the encoded header and payload
	inputValue := encodedHeader + "." + encodedPayload
	fmt.Println("InputValue:", inputValue)

	// Step 6: Create the HMAC-SHA-256 signature
	signature := createHMACSignature(secretKey, inputValue)
	fmt.Println("Signature:", signature)

	// Step 7: Concatenate InputValue and Signature
	finalJWS := inputValue + "." + signature
	fmt.Println("Final JWS:", finalJWS)

	return finalJWS
}
