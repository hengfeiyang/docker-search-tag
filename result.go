package main

import "time"

type response struct {
	Count   int                    `json:"count"`
	Next    string                 `json:"next"`
	Results responseResultSortable `json:"results"`
}

type responseResult struct {
	Name          string    `json:"name"`
	FullSize      int       `json:"full_size"`
	V2            bool      `json:"v2"`
	TagStatus     string    `json:"tag_status"`
	TagLastPulled time.Time `json:"tag_last_pulled"`
	TagLastPushed time.Time `json:"tag_last_pushed"`
}

type responseResultSortable []responseResult

func (r responseResultSortable) Len() int {
	return len(r)
}

func (r responseResultSortable) Less(i, j int) bool {
	return r[i].Name > r[j].Name
}

func (r responseResultSortable) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
