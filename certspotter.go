package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func certspotterFetcher(d string) ([]string, error) {
	var jsonResults = new([]struct {
		Domain []string `json:"dns_names"`
	})
	r, err := http.Get(fmt.Sprintf("https://certspotter.com/api/v0/certs?domain=%s", d))
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
	for _, subs := range *jsonResults {
		results = append(results, subs.Domain...)
	}

	return results, nil
}
