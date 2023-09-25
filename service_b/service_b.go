package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func callServiceA() {
	response, err := http.Get("http://localhost:10001/bing")

	if err != nil {
		log.Fatalln(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}

func callServiceC() {
	response, err := http.Get("http://localhost:10002/ping")

	if err != nil {
		log.Fatalln(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}

func main() {
	serviceATicker := time.NewTicker(1 * time.Second)
	serviceCTicker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-serviceATicker.C:
			callServiceA()
		case <-serviceCTicker.C:
			callServiceC()
		}
	}
}