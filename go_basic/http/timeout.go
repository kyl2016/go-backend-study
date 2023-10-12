package main

import "net/http"

func main() {
	//http.DefaultClient use Timeout=0, A Timeout of zero means no timeout.

	http.Client{}.Transport

	var netTransport = &http.Transport {
		Dial: (),
	}
}
