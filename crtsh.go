package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func crtshFetcher(d string) ([]string, error) {
	var jsonResults = new([]struct {
		Domain string `json:"name_value"`
	})
	r, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", d))
	if err != nil {
		return []string{}, err
	}
	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	err = json.Unmarshal(b, &jsonResults)
	if err != nil {
		return []string{}, err
	}

	var results []string
	for _, sub := range *jsonResults {
		results = append(results, sub.Domain)
	}

	return results, nil
}
