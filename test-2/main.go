package main

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

func main() {
	accessKey := "NtrT8kVWRaSjbXlP"
	secretKey := "0LUGh35uBen6d7E5AKUyOHFxy1ebP9zv"
	method := "POST"
	canonicalURL := "/api/v3/groups/G08916310298/link-request/adaccount"
	urlParameters := ""
	requestBody := map[string]string{
		"targetAdaccountId": "A08655312340",
	}

	hasRequestBody := requestBody != nil

	endpoint := "https://ads.line.me" + canonicalURL + urlParameters
	var requestBodyJson []byte
	var err error
	if hasRequestBody {
		requestBodyJson, err = json.Marshal(requestBody)
		if err != nil {
			fmt.Println("Error marshalling request body:", err)
			return
		}
	}

	contentType := "application/json"
	jwsHeader := encodeWithBase64([]byte(fmt.Sprintf(`{
		"alg": "HS256",
		"kid": "%s",
		"typ": "text/plain"
	}`, accessKey)))

	hexDigest := calcSHA256Digest(string(requestBodyJson))
	payloadDate := time.Now().UTC().Format("20060102")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", hexDigest, contentType, payloadDate, canonicalURL)
	fmt.Println(payload)
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
	var req *http.Request
	if hasRequestBody {
		req, err = http.NewRequest(method, endpoint, bytes.NewBuffer(requestBodyJson))
		req.Header.Set("Content-Type", contentType)
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for key, value := range httpHeaders {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
