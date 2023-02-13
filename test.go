package main

import (
	"regexp"
)

func main() {
	text := []byte(" nat-destination-address=\"42.1.66.19\" nat-destination-port=\"80\"")
	re := regexp.MustCompile(`nat-destination-address="([^"]+)"`)
	match := re.FindSubmatch(text)
	if match != nil {
		println("nat-destination-address:", string(match[1]))
	}

}
