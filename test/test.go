package test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Function to calculate SHA-256 digest
func calcSHA256Digest(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Function to encode value to Base64 URL-safe encoding
func encodeWithBase64(value []byte) string {
	return base64.RawURLEncoding.EncodeToString(value)
}

func CreateChildGroup() {
	// Setting parameters for your request
	accessKey := "NtrT8kVWRaSjbXlP"
	secretKey := "0LUGh35uBen6d7E5AKUyOHFxy1ebP9zv"
	method := "POST"
	canonicalURL := "/api/v3/groups/G08916310298/children"
	urlParameters := ""
	requestBody := map[string]string{
		"name": "test",
	}
	hasRequestBody := requestBody != nil

	endpoint := "https://ads.line.me" + canonicalURL + urlParameters
	var requestBodyJSON string
	var contentType string

	// Prepare request body and content type
	if hasRequestBody {
		bodyJSON, _ := json.Marshal(requestBody)
		requestBodyJSON = string(bodyJSON)
		contentType = "application/json"
	}

	// Create JWS header and encode with Base64
	jwsHeaderJSON, _ := json.Marshal(map[string]string{
		"alg":  "HS256",
		"kid":  accessKey,
		"typ":  "text/plain",
		"date": time.Now().UTC().Format("20060102"),
	})
	jwsHeader := encodeWithBase64(jwsHeaderJSON)

	// Create SHA-256 digest of the request body
	hexDigest := calcSHA256Digest(requestBodyJSON)

	// Prepare payload
	payloadDate := time.Now().UTC().Format("20060102")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", hexDigest, contentType, payloadDate, canonicalURL)
	jwsPayload := encodeWithBase64([]byte(payload))

	// Create signing input (header + payload)
	signingInput := fmt.Sprintf("%s.%s", jwsHeader, jwsPayload)

	// Create the HMAC-SHA-256 signature using the secret key
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(signingInput))
	signature := h.Sum(nil)
	encodedSignature := encodeWithBase64(signature)

	// Generate the final token
	token := fmt.Sprintf("%s.%s.%s", jwsHeader, jwsPayload, encodedSignature)

	// Prepare HTTP headers
	httpHeaders := map[string]string{
		"Date":          time.Now().UTC().Format(time.RFC1123),
		"Authorization": "Bearer " + token,
	}

	// Set content type header if there's a request body
	if hasRequestBody {
		httpHeaders["Content-Type"] = contentType
	}

	// Create the HTTP request
	var req *http.Request
	var err error
	if hasRequestBody {
		req, err = http.NewRequest(method, endpoint, bytes.NewBuffer([]byte(requestBodyJSON)))
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers to the request
	for key, value := range httpHeaders {
		req.Header.Add(key, value)
	}

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
