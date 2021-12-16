// Package docker-search-tag provider a tool list tags for docker search repository results

package main

import (
	"fmt"
	"os"
)

func main() {
	var name, tag, num string
	if len(os.Args) < 2 || len(os.Args[1]) == 0 {
		fmt.Println("Usage: docker-search-tag NAME [TAG] [NUM]")
		return
	}

	name = os.Args[1]
	num = "100"
	if len(os.Args) >= 3 && len(os.Args[2]) > 0 {
		tag = os.Args[2]
	}
	if len(os.Args) == 4 && len(os.Args[3]) > 0 {
		num = os.Args[3]
	}

	request(name, tag, num)
}
