package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("Sleep done!")
}

func showResponseTime(url string) {
	startTime := time.Now()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// close the response body when this function end
	defer res.Body.Close()
	elapsed := time.Since(startTime).Seconds()
	fmt.Printf("%s took %v seconds to response!\n", url, elapsed)
}

func handleApiRequest() {
	fmt.Println("****************** EXAMPLE OF HANDLING API REQUEST *******************")
	urls := make([]string, 3)
	urls[0] = "https://www.usa.gov/"
	urls[1] = "https://www.gov.uk/"
	urls[2] = "http://www.gouvernement.fr/"
	for _, url := range urls {
		// showResponseTime(url)
		// if you using goroutines here, the order of url response may differ with the ordinary array, corresponse to the reponse time of each url
		go showResponseTime(url)
	}
	time.Sleep(time.Second * 5)
}

func concurrencyDemo() {
	fmt.Println("****************** EXAMPLE OF CONCURRENCY *******************")
	// slowFunc()
	// fmt.Println("I only show after sleep done!")

	// go slowFunc()
	// fmt.Println("I show immediately!")
	// time.Sleep(time.Second * 3)
	// at this point, the function will exit, even there is a Groutine not finish yet.
}

func main() {
	// concurrencyDemo()
	// handleApiRequest()
}

/**
	We can make con run conconcurren with gorountines
	The keywork to using concurrency here is "go", so funny :)))
**/
