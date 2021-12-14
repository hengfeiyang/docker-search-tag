package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

func parse(name string, respBody []byte) {
	respData := new(response)
	err := json.Unmarshal(respBody, &respData)
	if err != nil {
		fmt.Printf("json.Unmarshal data err: %v\n", err)
		return
	}
	if len(respData.Results) == 0 {
		fmt.Printf("there is no tag data for name: %s\n", name)
		return
	}

	// resort results
	sort.Sort(respData.Results)

	maxLengthForName := 10
	for _, v := range respData.Results {
		if len(v.Name) > maxLengthForName {
			maxLengthForName = len(v.Name)
		}
	}

	fmt.Printf("Tags for %s:\n\n", name)
	fmt.Printf("%s%s%10s   %s\n", "TAG", strings.Repeat(" ", maxLengthForName), "SIZE(MB)", "LASTPUSH")
	for _, v := range respData.Results {
		fmt.Printf("%s%s   %10d   %s\n",
			v.Name, strings.Repeat(" ", max(0, maxLengthForName-len(v.Name))),
			v.FullSize/1024/1024,
			v.TagLastPushed.Format(time.RFC3339),
		)
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
