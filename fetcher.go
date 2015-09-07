package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AwsResponse struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IpPrefix string `json:"ip_prefix"`
		Region   string `json:"region"`
		Service  string `json:"service"`
	} `json:"prefixes"`
}

func fetch_aws() {
	url := "https://ip-ranges.amazonaws.com/ip-ranges.json"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var i AwsResponse
	if err := json.Unmarshal(body, &i); err != nil {
		panic(err)
	}

	for _, u := range i.Prefixes {
		fmt.Println(u.IpPrefix)
	}
}
