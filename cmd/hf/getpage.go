package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)

	/*var size int
	size = go responseSize("https://example.com/")
	fmt.Println(size)
	size = go responseSize("https://golang.org/")
	fmt.Println(size)
	size = go responseSize("https://golang.org/doc")
	fmt.Println(size)
	time.Sleep(5 * time.Second)*/

	/*go repeat("y")
	go repeat("x")
	time.Sleep(time.Second)*/
}

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func responseSize(url string) int {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return len(body)
}

func repeat(s string) {
	for i := 0; i < 25; i++ {
		fmt.Print(s)
	}
}
