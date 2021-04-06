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

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	client := lib.HTTPClientWithTimeoutAndAddress(*timeout, *address, *cookies)
	models, err := lib.StreamateOnlineAPI("http://affiliate.streamate.com/SMLive/SMLResult.xml", client, nil, *verbose, nil)
	if err != nil {
		fmt.Printf("error occurred: %v", err)
		return
	}
	for _, model := range models {
		fmt.Printf("%s %s\n", model.ModelID, model.Image)
	}
}
