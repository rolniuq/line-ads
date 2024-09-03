package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	host         = "ads.line.me"
	basePath     = "api"
	scheme       = "https"
	clientId     = "NtrT8kVWRaSjbXlP"
	clientSecret = "0LUGh35uBen6d7E5AKUyOHFxy1ebP9zv"
	authCode     = "123"
	redirectUrl  = "http://localhost:3000/callback"
)

type WithScope struct{}

func (w *WithScope) WithScope() {}

type Line struct{}

func NewLineClient() *Line {
	return &Line{}
}

func (l *Line) Auth() {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", redirectUrl)
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", "https://api.line.me/oauth2/v2.1/token", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}

func (l *Line) SendInviteLink() {

}

func (l *Line) GetCampaigns() {

}

func (l *Line) GetAds() {

}

func (l *Line) CreateAd() {

}
