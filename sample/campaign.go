package sample

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ReadGroups() {
	url := "https://ads.line.me/api/v3/groups/G08916310298/link-request"
	method := "POST"

	payload := strings.NewReader(`{"name":"test","campaignObjective":"VISIT_MY_WEBSITE"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Date", "Wed, 22 Dec 2021 00:00:00 GMT")
	req.Header.Add("Authorization", "Bearer "+GetJWS())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
