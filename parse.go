package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func parse(name string, respBody []byte) {
	respData := new(response)
	err := json.Unmarshal(respBody, &respData)
	if err != nil {
		fmt.Printf("json.Unmarshal data err: %v\n", err)
		return
	}
	if len(respData.Results) == 0 {
		fmt.Printf("there is no tag data for name: %s", name)
		return
	}

	// resort results
	sort.Sort(respData.Results)

	fmt.Printf("Tags for %s:\n", name)
	fmt.Println("-------------------------------")
	fmt.Printf("%s %16s %s\n", "Tag", "", "Size(MB)")
	fmt.Println("--------------------------------")
	for _, v := range respData.Results {
		fmt.Printf("%s%s %d\n", v.Name, strings.Repeat(" ", max(0, 20-len(v.Name))), v.FullSize/1024/1024)
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
