package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func urlscanIOFetcher(d string) ([]string, error) {
	var jsonResults = new(struct {
		Results []struct {
			Page struct {
				Domain string `json:"domain"`
			} `json:"page"`
		} `json:"results"`
	})
	r, err := http.Get(fmt.Sprintf("https://urlscan.io/api/v1/search/?q=domain:%s", d))
	if err != nil {
		return []string{}, err
	}
	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	err = json.Unmarshal(b, &jsonResults)
	if err != nil {
		return []string{}, err
	}

	var results = make([]string, 0)
	for _, p := range jsonResults.Results {
		results = append(results, p.Page.Domain)
	}

	return results, nil
}
