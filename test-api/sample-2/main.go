package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"line-ads/utils"
	"net/http"
	"net/url"
	"os"
)

var (
	// clientID      = "2006159961"
	clientID = "2007829527"
	// clientSecret  = "735c548e1ebd9ba8e4ee17b531ca30d6"
	clientSecret  = "a97c609ff67ee2e14ee7e8227d4984d3"
	redirectURI   = "http://localhost:8080/callback"
	authEndpoint  = "https://access.line.me/oauth2/v2.1/authorize"
	state         = "YOUR_STATE"
	scope         = "profile openid email"
	tokenEndpoint = "https://api.line.me/oauth2/v2.1/token"
)

// https://help.line.me/line/win/?contentId=50001462&lang=en

func main() {
	http.HandleFunc("/", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	authURL := fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s&scope=%s", authEndpoint, clientID, url.QueryEscape(redirectURI), state, scope)
	http.Redirect(w, r, authURL, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		return
	}

	accessToken, err := exchangeCodeForToken(code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get access token: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Println(accessToken)

	profile, err := getProfile(accessToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user profile: %v", err), http.StatusInternalServerError)
		return
	}

	utils.WriteFile("profile.json", profile)

	getLineAds(accessToken, profile.UserId)

	fmt.Fprintf(w, "User Id: %s, Name: %s", profile.UserId, profile.DisplayName)
}

func exchangeCodeForToken(code string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	resp, err := http.PostForm(tokenEndpoint, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	os.WriteFile("result.json", body, 0644)

	if accessToken, ok := result["access_token"].(string); ok {
		return accessToken, nil
	}

	return "", fmt.Errorf("unable to fetch access token")
}

type Profile struct {
	UserId      string `json:"userId"`
	DisplayName string `json:"displayName"`
}

func getProfile(accessToken string) (*Profile, error) {
	req, _ := http.NewRequest("GET", "https://api.line.me/v2/profile", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received non-OK HTTP status %s\n", resp.Status)
		fmt.Printf("Response body: %s\n", string(bodyBytes))
		return nil, fmt.Errorf("received non-OK HTTP status %s", resp.Status)
	}

	str := string(bodyBytes)
	fmt.Println(str)

	var profile Profile
	err = json.Unmarshal(bodyBytes, &profile)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return nil, err
	}

	return &profile, nil
}

type Ad struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type AdsResponse struct {
	Ads []Ad `json:"ads"`
}

func getLineAds(accessToken string, accountId string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.line.me/v3/adaccounts/%s/ads", accountId), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received non-OK HTTP status %s\n", resp.Status)
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Response body: %s\n", string(bodyBytes))
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var adsResponse AdsResponse
	err = json.Unmarshal(bodyBytes, &adsResponse)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	fmt.Println("Ads retrieved:")
	for _, ad := range adsResponse.Ads {
		fmt.Printf("ID: %s, Name: %s, Description: %s, Status: %s\n", ad.ID, ad.Name, ad.Description, ad.Status)
	}
}
