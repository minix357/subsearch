package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func bufferoverrunFetcher(d string) ([]string, error) {
	var jsonResults = new(struct {
		Domains []string `json:"FDNS_A"`
	})
	r, err := http.Get(fmt.Sprintf("https://dns.bufferover.run/dns?q=.%s", d))
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
