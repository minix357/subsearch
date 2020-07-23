package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Fetcher definition
type fchr func(string) ([]string, error)

func main() {
	var o = flag.String("o", "", "Set the output file if you want to save subdomains.")
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println("syntax: subsearch [-o output.txt] [domain].")
		os.Exit(1)
	}
	var url string = flag.Arg(0)

	var subdomains = make([]string, 0)
	var wg sync.WaitGroup
	var e = make([]string, 0)
	var s = make(chan []string) // subdomains channel
	var fetchers = []fchr{
		crtshFetcher,
		certspotterFetcher,
		bufferoverrunFetcher,
		threadcrowdFetcher,
		urlscanIOFetcher,
	}
	var q = make(chan int, len(fetchers)) // when full, quit.

	for _, f := range fetchers {
		wg.Add(1)
		go func(url string, f fchr, wg *sync.WaitGroup) {
			defer wg.Done()
			subs, err := f(url)
			if err != nil {
				e = append(e, err.Error())
			}
			s <- subs
		}(url, f, &wg)
	}
	wg.Add(1)
	go func(url string, q chan int, s chan []string, wg *sync.WaitGroup) {
		var qcount = 0
		for {
			select {
			case subs := <-s:
				for i := 0; i < len(subs); i++ {
					sub := subs[i]
					sub = strings.ReplaceAll(strings.ReplaceAll(sub, "*.", ""), " ", "")
					if !strings.Contains(sub, url) {
						continue
					}
					if strings.Contains(sub, "\n") {
						subs = append(subs, strings.Split(sub, "\n")...)
						continue
					}
					exists := false
					for _, sd := range subdomains {
						if sub == sd {
							exists = true
						}
					}
					if !exists && sub != "" {
						subdomains = append(subdomains, sub)
						fmt.Println(sub)
					}
				}
				q <- 1
			case c := <-q:
				qcount += c
				if qcount == len(fetchers) {
					wg.Done()
					return
				}
			}
		}
	}(url, q, s, &wg)
	wg.Wait()
	close(q)
	close(s)

	if !(*o == "") {
		f, err := os.OpenFile(*o, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			e = append(e, err.Error())
		}
		if _, err = f.WriteString(strings.Join(subdomains, "\n") + "\n"); err != nil {
			e = append(e, err.Error())
		}
		f.Close()
	}
	/* print errors, etc... */
	if len(e) > 0 {
		fmt.Println("The following errors have been generated:")
		for _, cErr := range e {
			fmt.Println(cErr)
		}
	}
}
