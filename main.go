// Package docker-search-tag provider a tool list tags for docker search repository results

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var name, num string
	if len(os.Args) < 2 || len(os.Args[1]) == 0 {
		fmt.Println("Usage: docker-search-tag NAME [NUM]")
		return
	}

	name = os.Args[1]
	num = "100"
	if len(os.Args) == 3 && len(os.Args[2]) > 0 {
		num = os.Args[2]
	}

	url := fmt.Sprintf("https://registry.hub.docker.com/v2/repositories/library/%s/tags/?page_size=%s", name, num)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Get information from registry.hub.docker.com err: %v\n", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read data err: %v\n", err)
		return
	}

	// fmt.Printf("data json:\n%s\n", respBody)
	parse(name, respBody)
}
