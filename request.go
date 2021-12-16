package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func request(name, tag, num string) {
	url := fmt.Sprintf("https://registry.hub.docker.com/v2/repositories/library/%s/tags/?name=%s&page_size=%s", name, tag, num)
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

	// fmt.Println(url)
	// fmt.Printf("data json:\n%s\n", respBody)
	parse(name, respBody)
}
