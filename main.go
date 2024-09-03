package main

import "line-ads/sample"

func main() {
	sample.Run()

	l := NewLineClient()
	l.Auth()
}
