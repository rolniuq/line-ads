package main

import "line-ads/sample"

func main() {
	sample.ReadGroups()

	l := NewLineClient()
	l.Auth()
}
