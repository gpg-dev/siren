package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bcmk/siren/lib"
)

var verbose = flag.Bool("v", false, "verbose output")
var timeout = flag.Int("t", 10, "timeout in seconds")
var address = flag.String("a", "", "source IP address")
var cookies = flag.Bool("c", false, "use cookies")
var clientID = flag.String("id", "", "your client id")
var secret = flag.String("secret", "", "your client secret")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	httpClient := lib.HTTPClientWithTimeoutAndAddress(*timeout, *address, *cookies)
	res, err := lib.TwitchOnlineAPI("", httpClient, nil, *verbose, map[string]string{
		"client_id":     *clientID,
		"client_secret": *secret,
	})
	if err != nil {
		fmt.Printf("error occurred: %v", err)
		return
	}
	for s := range res {
		fmt.Println(s)
	}
}
