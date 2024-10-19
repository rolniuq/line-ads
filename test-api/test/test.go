package test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func CreateChildGroup() error {
	// Setting parameters for your request
	accessKey := "NtrT8kVWRaSjbXlP"
	secretKey := "0LUGh35uBen6d7E5AKUyOHFxy1ebP9zv"
	method := "POST"
	canonicalURL := "/api/v3/groups/G08916310298/children"
	urlParameters := ""
	requestBody := map[string]interface{}{"name": "test group"}
	// hasRequestBody := requestBody != nil

	endpoint := "https://ads.line.me" + canonicalURL + urlParameters
	requestBodyJSON, _ := json.Marshal(requestBody)
	contentType := "application/json"

	jwsHeader := encodeWithBase64([]byte(fmt.Sprintf(`{"alg":"HS256","kid":"%s","typ":"text/plain"}`, accessKey)))

	hexDigest := calcSHA256Digest(string(requestBodyJSON))
	payloadDate := time.Now().UTC().Format("20060102")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", hexDigest, contentType, payloadDate, canonicalURL)
	jwsPayload := encodeWithBase64([]byte(payload))

	signingInput := fmt.Sprintf("%s.%s", jwsHeader, jwsPayload)
	signature := hmac.New(sha256.New, []byte(secretKey))
	signature.Write([]byte(signingInput))
	encodedSignature := encodeWithBase64(signature.Sum(nil))
	token := fmt.Sprintf("%s.%s.%s", jwsHeader, jwsPayload, encodedSignature)

	httpHeaders := http.Header{
		"Date":          {time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")},
		"Authorization": {"Bearer " + token},
		"Content-Type":  {contentType},
	}

	req, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(requestBodyJSON))
	req.Header = httpHeaders

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return nil
}

func GetChildGroups() error {
	// Setting parameters for your request
	accessKey := "NtrT8kVWRaSjbXlP"
	secretKey := "0LUGh35uBen6d7E5AKUyOHFxy1ebP9zv"
	method := "GET"
	canonicalURL := "/api/v3/groups/G08916310298/children"
	urlParameters := ""
	requestBody := ""
	// hasRequestBody := requestBody != nil

	endpoint := "https://ads.line.me" + canonicalURL + urlParameters
	requestBodyJSON, _ := json.Marshal(requestBody)
	contentType := ""

	jwsHeader := encodeWithBase64([]byte(fmt.Sprintf(`{"alg":"HS256","kid":"%s","typ":"text/plain"}`, accessKey)))

	hexDigest := calcSHA256Digest(string(requestBodyJSON))
	payloadDate := time.Now().UTC().Format("20060102")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", hexDigest, contentType, payloadDate, canonicalURL)
	jwsPayload := encodeWithBase64([]byte(payload))

	signingInput := fmt.Sprintf("%s.%s", jwsHeader, jwsPayload)
	signature := hmac.New(sha256.New, []byte(secretKey))
	signature.Write([]byte(signingInput))
	encodedSignature := encodeWithBase64(signature.Sum(nil))
	token := fmt.Sprintf("%s.%s.%s", jwsHeader, jwsPayload, encodedSignature)

	httpHeaders := http.Header{
		"Date":          {time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")},
		"Authorization": {"Bearer " + token},
		"Content-Type":  {contentType},
	}

	req, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(requestBodyJSON))
	req.Header = httpHeaders

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status %s", resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return nil
}

func GetListAdsAccounts() error {
	accessKey := "NtrT8kVWRaSjbXlP"
	secretKey := "0LUGh35uBen6d7E5AKUyOHFxy1ebP9zv"
	method := "GET"
	canonicalURL := "/api/v3/groups/G08916310298/adaccounts"
	urlParameters := "?includeRemoved=false"
	requestBody := ""
	// hasRequestBody := requestBody != nil

	endpoint := "https://ads.line.me" + canonicalURL + urlParameters
	requestBodyJSON, _ := json.Marshal(requestBody)
	contentType := ""

	jwsHeader := encodeWithBase64([]byte(fmt.Sprintf(`{"alg":"HS256","kid":"%s","typ":"text/plain"}`, accessKey)))

	hexDigest := calcSHA256Digest(string(requestBodyJSON))
	payloadDate := time.Now().UTC().Format("20060102")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", hexDigest, contentType, payloadDate, canonicalURL)
	jwsPayload := encodeWithBase64([]byte(payload))

	signingInput := fmt.Sprintf("%s.%s", jwsHeader, jwsPayload)
	signature := hmac.New(sha256.New, []byte(secretKey))
	signature.Write([]byte(signingInput))
	encodedSignature := encodeWithBase64(signature.Sum(nil))
	token := fmt.Sprintf("%s.%s.%s", jwsHeader, jwsPayload, encodedSignature)

	httpHeaders := http.Header{
		"Date":          {time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")},
		"Authorization": {"Bearer " + token},
		"Content-Type":  {contentType},
	}

	req, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(requestBodyJSON))
	req.Header = httpHeaders

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status %s", resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	os.WriteFile("list-ads-accounts.json", body, 0644)
	fmt.Println(string(body))

	return nil
}

type Status string

const (
	LINKED Status = "LINKED"
)

type LinkRequest struct {
	Id                  int    `json:"id"`
	SourceGroupId       string `json:"sourceGroupId"`
	SourceGroupName     string `json:"sourceGroupName"`
	TargetAdaccountId   string `json:"targetAdaccountId"`
	TargetAdaccountName string `json:"targetAdaccountName"`
	Status              Status `json:"status"`
}

func SendLinkRequest() error {
	// Setting parameters for your request
	accessKey := "NtrT8kVWRaSjbXlP"
	secretKey := "0LUGh35uBen6d7E5AKUyOHFxy1ebP9zv"
	method := "POST"
	canonicalURL := "/api/v3/groups/G01679313261/link-request/adaccount"
	urlParameters := ""
	requestBody := &LinkRequest{
		Id:                  101,
		SourceGroupId:       "G01679313261",
		SourceGroupName:     "test group",
		TargetAdaccountId:   "A43859314796",
		TargetAdaccountName: "Quynh",
		Status:              LINKED,
	}

	endpoint := "https://ads.line.me" + canonicalURL + urlParameters
	requestBodyJSON, _ := json.Marshal(requestBody)
	contentType := "application/json"

	jwsHeader := encodeWithBase64([]byte(fmt.Sprintf(`{"alg":"HS256","kid":"%s","typ":"text/plain"}`, accessKey)))

	hexDigest := calcSHA256Digest(string(requestBodyJSON))
	payloadDate := time.Now().UTC().Format("20060102")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", hexDigest, contentType, payloadDate, canonicalURL)
	jwsPayload := encodeWithBase64([]byte(payload))

	signingInput := fmt.Sprintf("%s.%s", jwsHeader, jwsPayload)
	signature := hmac.New(sha256.New, []byte(secretKey))
	signature.Write([]byte(signingInput))
	encodedSignature := encodeWithBase64(signature.Sum(nil))
	token := fmt.Sprintf("%s.%s.%s", jwsHeader, jwsPayload, encodedSignature)

	httpHeaders := http.Header{
		"Date":          {time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")},
		"Authorization": {"Bearer " + token},
		"Content-Type":  {contentType},
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return err
	}
	req.Header = httpHeaders

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		return fmt.Errorf("status code error: %v", buf.String())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return nil
}

func GetLinkRequest() error {
	accessKey := "NtrT8kVWRaSjbXlP"
	secretKey := "0LUGh35uBen6d7E5AKUyOHFxy1ebP9zv"
	method := "GET"
	canonicalURL := "/api/v3/groups/G08916310298/link-request"
	urlParameters := ""
	requestBody := ""
	// hasRequestBody := requestBody != nil

	endpoint := "https://ads.line.me" + canonicalURL + urlParameters
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}
	contentType := "application/json"

	jwsHeader := encodeWithBase64([]byte(fmt.Sprintf(`{"alg":"HS256","kid":"%s","typ":"text/plain"}`, accessKey)))

	hexDigest := calcSHA256Digest(string(requestBodyJSON))
	payloadDate := time.Now().UTC().Format("20060102")
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", hexDigest, contentType, payloadDate, canonicalURL)
	jwsPayload := encodeWithBase64([]byte(payload))

	signingInput := fmt.Sprintf("%s.%s", jwsHeader, jwsPayload)
	signature := hmac.New(sha256.New, []byte(secretKey))
	signature.Write([]byte(signingInput))
	encodedSignature := encodeWithBase64(signature.Sum(nil))
	token := fmt.Sprintf("%s.%s.%s", jwsHeader, jwsPayload, encodedSignature)

	httpHeaders := http.Header{
		"Date":          {time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")},
		"Authorization": {"Bearer " + token},
		"Content-Type":  {contentType},
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header = httpHeaders

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send request: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	os.WriteFile("link-requests.json", body, 0644)
	fmt.Println(string(body))

	return nil
}
