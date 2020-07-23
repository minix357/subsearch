# SubSearch

SubSearch is a fast and concurrent software written in Go which, given a domain name, will  
query online services in search of it's subdomains.  
Sometimes subdomains can contain staging software, tests created by developers and other interesting stuff; using  
SubSearch, you will quickly obtain a list of the currently known subdomains for future analysis.  
**Make sure you are authorized to conduct analysis on hosts you don't own.**

## Install
SubSearch can be installed in two ways; the first requires Go, Git and Make:
```console
user@host:˜$ git clone github.com/minix357/subsearch
user@host:˜$ cd subsearch
user@host:˜$ sudo make all # or 'su; make all' if you're a badass unix hacker.
```
The other is a more classical way and it only requires Go and Git:
```console
user@host:˜$ go get github.com/minix357/subsearch
user@host:˜$ sudo mv $GOPATH/bin/subsearch /usr/local/bin/ # optional, unless you want to cd into $GOPATH every time.
```
___
Have fun
