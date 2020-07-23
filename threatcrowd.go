package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func threadcrowdFetcher(d string) ([]string, error) {
	var jsonResults = new(struct {
		Domains []string `json:"subdomains"`
	})
	r, err := http.Get(fmt.Sprintf("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=%s", d))
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
	for _, sub := range jsonResults.Domains {
		s := strings.Split(sub, ",")
		results = append(results, s[len(s)-1])
	}

	return results, nil
}
