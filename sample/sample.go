package sample

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

func calcSHA256Digest(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func encodeWithBase64(value []byte) string {
	return base64.URLEncoding.EncodeToString(value)
}

func Run() {
	accessKey := "NtrT8kVWRaSjbXlP"
	secretKey := "0LUGh35uBen6d7E5AKUyOHFxy1ebP9zv"
	method := "POST"
	canonicalURL := "/api/v3/groups/G1/children"
	urlParameters := ""
	requestBody := map[string]string{"name": "test"}
	hasRequestBody := requestBody != nil

	endpoint := "https://ads.line.me" + canonicalURL + urlParameters
	var requestBodyJSON []byte
	var err error
	if hasRequestBody {
		requestBodyJSON, err = json.Marshal(requestBody)
		if err != nil {
			fmt.Println("Failed to encode request body:", err)
			return
		}
	}

	contentType := ""
	if hasRequestBody {
		contentType = "application/json"
	}

	jwsHeader := encodeWithBase64([]byte(fmt.Sprintf(
		`{"alg":"HS256","kid":"%s","typ":"text/plain"}`,
		accessKey,
	)))

	hexDigest := calcSHA256Digest(string(requestBodyJSON))
	payloadDate := time.Now().UTC().Format("20060102")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", hexDigest, contentType, payloadDate, canonicalURL)
	jwsPayload := encodeWithBase64([]byte(payload))

	signingInput := fmt.Sprintf("%s.%s", jwsHeader, jwsPayload)
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(signingInput))
	signature := h.Sum(nil)
	encodedSignature := encodeWithBase64(signature)
	token := fmt.Sprintf("%s.%s.%s", jwsHeader, jwsPayload, encodedSignature)

	httpHeaders := map[string]string{
		"Date":          time.Now().UTC().Format(http.TimeFormat),
		"Authorization": "Bearer " + token,
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		fmt.Println("Failed to create HTTP request:", err)
		return
	}

	for key, value := range httpHeaders {
		req.Header.Set(key, value)
	}

	if hasRequestBody {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Failed to decode response:", err)
		return
	}

	fmt.Println(result)
}
