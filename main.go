package main

import lineads "line-ads/builder"

func main() {
	client := lineads.NewLineAdsService("accessKey", "secretKey")

	client.CreateCampaign("accountId")
}
